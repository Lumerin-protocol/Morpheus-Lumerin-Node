package main

import (
	"context"
	"fmt"
	"net/url"
	"os"
	"os/signal"
	"path/filepath"
	"syscall"
	"time"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/aiengine"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/apibus"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/blockchainapi"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/config"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/handlers/httphandlers"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/interfaces"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/proxyapi"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/proxyctl"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/registries"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/transport"
	wlt "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/wallet"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/storages"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/system"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/walletapi"
	"github.com/ethereum/go-ethereum/ethclient"

	docs "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/docs"
)

const (
	IDLE_READ_CLOSE_TIMEOUT  = 10 * time.Minute
	IDLE_WRITE_CLOSE_TIMEOUT = 10 * time.Minute
)

var (
	ErrConnectToEthNode = fmt.Errorf("cannot connect to ethereum node")
)

func main() {
	err := start()
	if err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
	os.Exit(0)
}

func start() error {
	valid, err := config.NewValidator()
	if err != nil {
		return err
	}

	var cfg config.Config
	err = config.LoadConfig(&cfg, &os.Args, valid)
	if err != nil {
		return err
	}

	fmt.Printf("Loaded config: %+v\n", cfg.GetSanitized())

	mainLogFilePath := ""
	logFolderPath := ""
	appStartTime := time.Now()

	if cfg.Log.FolderPath != "" {
		folderName := lib.SanitizeFilename(appStartTime.Format("2006-01-02T15-04-05Z07:00"))
		logFolderPath = filepath.Join(cfg.Log.FolderPath, folderName)
		err = os.MkdirAll(logFolderPath, os.ModePerm)
		if err != nil {
			return err
		}

		mainLogFilePath = filepath.Join(logFolderPath, "main.log")
	}

	if cfg.Web.PublicUrl != "" {
		hostWithoutProtocol := cfg.Web.PublicUrl
		if u, err := url.Parse(cfg.Web.PublicUrl); err == nil {
			hostWithoutProtocol = u.Host
		}
		docs.SwaggerInfo.Host = hostWithoutProtocol
	} else if cfg.Web.Address != "" {
		docs.SwaggerInfo.Host = cfg.Web.Address
	} else {
		docs.SwaggerInfo.Host = "localhost:8082"
	}

	log, err := lib.NewLogger(cfg.Log.LevelApp, cfg.Log.Color, cfg.Log.IsProd, cfg.Log.JSON, mainLogFilePath)
	if err != nil {
		return err
	}

	appLog := log.Named("APP")

	proxyLog, err := lib.NewLogger(cfg.Log.LevelProxy, cfg.Log.Color, cfg.Log.IsProd, cfg.Log.JSON, mainLogFilePath)
	if err != nil {
		return err
	}

	connLog, err := lib.NewLogger(cfg.Log.LevelConnection, cfg.Log.Color, cfg.Log.IsProd, cfg.Log.JSON, mainLogFilePath)
	if err != nil {
		return err
	}

	schedulerLogFactory := func(remoteAddr string) (lib.ILogger, error) {
		fp := ""
		if logFolderPath != "" {
			fp = filepath.Join(logFolderPath, fmt.Sprintf("scheduler-%s.log", lib.SanitizeFilename(remoteAddr)))
		}
		return lib.NewLogger(cfg.Log.LevelScheduler, cfg.Log.Color, cfg.Log.IsProd, cfg.Log.JSON, fp)
	}

	contractLogStorage := lib.NewCollection[*interfaces.LogStorage]()

	// contractLogFactory := func(contractID string) (lib.ILogger, error) {
	// 	logStorage := interfaces.NewLogStorage(contractID)
	// 	contractLogStorage.Store(logStorage)
	// 	fp := ""
	// 	if logFolderPath != "" {
	// 		fp = filepath.Join(logFolderPath, fmt.Sprintf("contract-%s.log", lib.SanitizeFilename(lib.StrShort(contractID))))
	// 	}
	// 	return lib.NewLoggerMemory(cfg.Log.LevelContract, cfg.Log.Color, cfg.Log.IsProd, cfg.Log.JSON, fp, logStorage.Buffer)
	// }

	defer func() {
		_ = connLog.Close()
		_ = proxyLog.Close()
		_ = log.Close()
	}()

	appLog.Infof("proxy-router %s", config.BuildVersion)

	sysConfig := system.CreateConfigurator(log)

	if cfg.System.Enable {
		if err != nil {
			return err
		}

		err = sysConfig.ApplyConfig(&system.Config{
			LocalPortRange:   cfg.System.LocalPortRange,
			TcpMaxSynBacklog: cfg.System.TcpMaxSynBacklog,
			Somaxconn:        cfg.System.Somaxconn,
			NetdevMaxBacklog: cfg.System.NetdevMaxBacklog,
			RlimitSoft:       cfg.System.RlimitSoft,
			RlimitHard:       cfg.System.RlimitHard,
		})
		if err != nil {
			appLog.Warnf("failed to apply system config, try using sudo or set SYS_ENABLE to false to disable\n%s", err)
			return err
		}

		defer func() {
			err = sysConfig.RestoreConfig()
			if err != nil {
				appLog.Warnf("failed to restore system config\n%s", err)
				return
			}
		}()
	}

	// graceful shutdown
	ctx, cancel := context.WithCancel(context.Background())
	defer cancel()

	shutdownChan := make(chan os.Signal, 1)
	signal.Notify(shutdownChan, syscall.SIGINT, syscall.SIGTERM)
	go func() {
		s := <-shutdownChan
		appLog.Warnf("Received signal: %s", s)
		cancel()

		s = <-shutdownChan

		appLog.Warnf("Received signal: %s. \n", s)

		appLog.Warnf("Forcing exit...")
		os.Exit(1)
	}()

	ethClient, err := ethclient.DialContext(ctx, cfg.Blockchain.EthNodeAddress)
	if err != nil {
		return lib.WrapError(ErrConnectToEthNode, err)
	}
	chainID, err := ethClient.ChainID(ctx)
	if err != nil {
		return lib.WrapError(ErrConnectToEthNode, err)
	}
	appLog.Infof("connected to ethereum node: %s, chainID: %d", cfg.Blockchain.EthNodeAddress, chainID)

	publicUrl, err := url.Parse(cfg.Web.PublicUrl)
	if err != nil {
		return err
	}

	storage := storages.NewStorage(log, cfg.Proxy.StoragePath)
	sessionStorage := storages.NewSessionStorage(storage)

	var wallet interfaces.Wallet
	if len(*cfg.Marketplace.WalletPrivateKey) > 0 {
		wallet = wlt.NewEnvWallet(*cfg.Marketplace.WalletPrivateKey)
		log.Warnf("Using env wallet. Private key persistance unavailable")
	} else {
		wallet = wlt.NewKeychainWallet()
		log.Infof("Using keychain wallet")
	}

	modelConfigLoader := config.NewModelConfigLoader(cfg.Proxy.ModelsConfigPath, log)
	err = modelConfigLoader.Init()
	if err != nil {
		log.Warnf("failed to load model config: %s, run with empty", err)
	}

	proxyRouterApi := proxyapi.NewProxySender(publicUrl, wallet, contractLogStorage, sessionStorage, log)
	blockchainApi := blockchainapi.NewBlockchainService(ethClient, *cfg.Marketplace.DiamondContractAddress, *cfg.Marketplace.MorTokenAddress, cfg.Blockchain.ExplorerApiUrl, wallet, sessionStorage, proxyRouterApi, proxyLog, cfg.Blockchain.EthLegacyTx)
	aiEngine := aiengine.NewAiEngine(cfg.AIEngine.OpenAIBaseURL, cfg.AIEngine.OpenAIKey, modelConfigLoader, log)

	sessionRouter := registries.NewSessionRouter(*cfg.Marketplace.DiamondContractAddress, ethClient, log)

	eventListener := blockchainapi.NewEventsListener(ethClient, sessionStorage, sessionRouter, wallet, modelConfigLoader, log)

	blockchainController := blockchainapi.NewBlockchainController(blockchainApi, log)

	var chatStorage proxyapi.ChatStorageInterface
	if cfg.Proxy.StoreChatContext {
		chatStoragePath := filepath.Join(cfg.Proxy.StoragePath, "chats")
		chatStorage = proxyapi.NewChatStorage(chatStoragePath)
	} else {
		log.Warnf("chat context storage is disabled")
		chatStorage = proxyapi.NewNoOpChatStorage()
	}

	proxyController := proxyapi.NewProxyController(proxyRouterApi, aiEngine, chatStorage)
	walletController := walletapi.NewWalletController(wallet)
	systemController := system.NewSystemController(&cfg, wallet, sysConfig, appStartTime, chainID, log)

	apiBus := apibus.NewApiBus(blockchainController, proxyController, walletController, systemController)
	httpHandler := httphandlers.CreateHTTPServer(log, apiBus)
	httpServer := transport.NewServer(cfg.Web.Address, httpHandler, log.Named("HTTP"))

	// http server should shut down latest to keep pprof running
	serverErrCh := make(chan error, 1)
	serverCtx, cancelServer := context.WithCancel(context.Background())
	go func() {
		serverErrCh <- httpServer.Run(serverCtx)
		cancel()
	}()

	proxy := proxyctl.NewProxyCtl(eventListener, wallet, chainID, log, connLog, cfg.Proxy.Address, schedulerLogFactory, sessionStorage, modelConfigLoader, valid, aiEngine)
	err = proxy.Run(ctx)

	cancelServer()
	<-serverErrCh

	appLog.Warnf("App exited due to %s", err)
	return err
}
