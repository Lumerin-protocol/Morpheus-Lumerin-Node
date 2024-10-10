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
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/contracts"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/ethclient"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/keychain"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/registries"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/transport"
	wlt "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/wallet"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/storages"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/system"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/walletapi"

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

	keychainStorage := keychain.NewKeychain()

	if cfg.App.ResetKeychain {
		appLog.Warnf("Resetting keychain...")
		wallet := wlt.NewKeychainWallet(keychainStorage)
		err = wallet.DeleteWallet()
		if err != nil {
			appLog.Warnf("Failed to delete wallet\n%s", err)
		} else {
			appLog.Info("Wallet deleted")
		}

		ethNodeStorage := ethclient.NewRPCClientStoreKeychain(keychainStorage, nil, log)
		err = ethNodeStorage.RemoveURLs()
		if err != nil {
			appLog.Warnf("Failed to remove eth node urls\n%s", err)
		} else {
			appLog.Info("Eth node urls removed")
		}
	}

	var ethNodeAddresses []string
	if cfg.Blockchain.EthNodeAddress != "" {
		ethNodeAddresses = []string{cfg.Blockchain.EthNodeAddress}
	}
	rpcClientStore, err := ethclient.ConfigureRPCClientStore(keychainStorage, ethNodeAddresses, cfg.Blockchain.ChainID, log.Named("RPC"))
	if err != nil {
		return lib.WrapError(ErrConnectToEthNode, err)
	}

	ethClient := ethclient.NewClient(rpcClientStore.GetClient())
	chainID, err := ethClient.ChainID(ctx)
	if err != nil {
		return lib.WrapError(ErrConnectToEthNode, err)
	}
	if cfg.Blockchain.ChainID != 0 && int(chainID.Uint64()) != cfg.Blockchain.ChainID {
		return lib.WrapError(ErrConnectToEthNode, fmt.Errorf("configured chainID (%d) does not match eth node chain ID (%s)", cfg.Blockchain.ChainID, chainID))
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
		wallet = wlt.NewKeychainWallet(keychainStorage)
		log.Infof("Using keychain wallet")
	}

	var logWatcher contracts.LogWatcher
	if cfg.Blockchain.UseSubscriptions {
		logWatcher = contracts.NewLogWatcherSubscription(ethClient, cfg.Blockchain.MaxReconnects, log)
		appLog.Infof("using websocket log subscription for blockchain events")
	} else {
		logWatcher = contracts.NewLogWatcherPolling(ethClient, cfg.Blockchain.PollingInterval, cfg.Blockchain.MaxReconnects, log)
		appLog.Infof("using polling for blockchain events")
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

	eventListener := blockchainapi.NewEventsListener(ethClient, sessionStorage, sessionRouter, wallet, modelConfigLoader, logWatcher, log)

	blockchainController := blockchainapi.NewBlockchainController(blockchainApi, log)

	chatStoragePath := filepath.Join(cfg.Proxy.StoragePath, "chat")
	chatStorage := proxyapi.NewChatStorage(chatStoragePath)
	proxyController := proxyapi.NewProxyController(proxyRouterApi, aiEngine, chatStorage)
	walletController := walletapi.NewWalletController(wallet)
	systemController := system.NewSystemController(&cfg, wallet, rpcClientStore, sysConfig, appStartTime, chainID, log)

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
