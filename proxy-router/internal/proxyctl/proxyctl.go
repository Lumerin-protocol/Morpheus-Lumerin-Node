package proxyctl

import (
	"context"

	"github.com/Lumerin-protocol/Morpheus-Lumerin-Node/proxy-router/internal/handlers/tcphandlers"
	"github.com/Lumerin-protocol/Morpheus-Lumerin-Node/proxy-router/internal/interfaces"
	"github.com/Lumerin-protocol/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/Lumerin-protocol/Morpheus-Lumerin-Node/proxy-router/internal/repositories/transport"
	"github.com/Lumerin-protocol/Morpheus-Lumerin-Node/proxy-router/internal/storages"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/aiengine"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/blockchainapi"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/handlers/tcphandlers"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/interfaces"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/proxyapi"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/transport"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/storages"
	"github.com/go-playground/validator/v10"
	"golang.org/x/sync/errgroup"
)

type ProxyState int8

const (
	StateStopped        ProxyState = 0
	StateWaitingForPkey            = 1
	StateRunning                   = 2
	StateStopping                  = 3
)

func (s ProxyState) String() string {
	switch s {
	case StateStopped:
		return "stopped"
	case StateWaitingForPkey:
		return "waiting for private key"
	case StateRunning:
		return "running"
	case StateStopping:
		return "stopping"
	}
	return "unknown"
}

type SchedulerLogFactory = func(remoteAddr string) (lib.ILogger, error)

// Proxy is a struct that represents a proxy-router part of the system
type Proxy struct {
	eventListener       *blockchainapi.EventsListener
	wallet              interfaces.PrKeyProvider
	proxyAddr           string
	sessionStorage      *storages.SessionStorage
	log                 *lib.Logger
	connLog             *lib.Logger
	schedulerLogFactory SchedulerLogFactory
	aiEngine            *aiengine.AiEngine
	validator           *validator.Validate

	state lib.AtomicValue[ProxyState]
	tsk   *lib.Task
}

// NewProxyCtl creates a new Proxy controller instance
func NewProxyCtl(eventListerer *blockchainapi.EventsListener, wallet interfaces.PrKeyProvider, log *lib.Logger, connLog *lib.Logger, proxyAddr string, scl SchedulerLogFactory, sessionStorage *storages.SessionStorage, valid *validator.Validate, aiEngine *aiengine.AiEngine) *Proxy {
	return &Proxy{
		eventListener:       eventListerer,
		wallet:              wallet,
		log:                 log,
		connLog:             connLog,
		proxyAddr:           proxyAddr,
		schedulerLogFactory: scl,
		sessionStorage:      sessionStorage,
		aiEngine:            aiEngine,
		validator:           valid,
	}
}

func (p *Proxy) Run(ctx context.Context) error {
	var tsk *lib.Task

	for { // restart proxy loop

		prKey, err := p.wallet.GetPrivateKey()
		if err != nil {
			p.log.Errorf("cannot get private key, waiting for its update: %s", err)
			p.setState(StateWaitingForPkey)

			select {
			case <-ctx.Done():
				return ctx.Err()
			case <-p.wallet.PrivateKeyUpdated():
				continue // restart the proxy
			}
		}

		tsk = lib.NewTaskFunc(func(ctx context.Context) error {
			return p.run(ctx, prKey)
		})

		tsk.Start(ctx)
		p.setState(StateRunning)

		select {
		case <-p.wallet.PrivateKeyUpdated():
			ch := tsk.Stop()
			p.setState(StateStopping)
			<-ch
			continue // restart the proxy
		case <-tsk.Done():
			err := tsk.Err()
			if err != nil {
				p.log.Errorf("proxy stopped with error: %s", err)
				return err
			}
		}
	}
}

func (p *Proxy) run(ctx context.Context, prKey lib.HexString) error {
	tcpServer := transport.NewTCPServer(p.proxyAddr, p.connLog.Named("TCP"))
	prKey, err := p.wallet.GetPrivateKey()
	if err != nil {
		return err
	}

	walletAddr, err := lib.PrivKeyBytesToAddr(prKey)
	if err != nil {
		return err
	}
	p.log.Infof("Wallet address: %s", walletAddr.String())

	pubKey, err := lib.PubKeyFromPrivate(prKey)
	if err != nil {
		return err
	}

	proxyReceiver := proxyapi.NewProxyReceiver(prKey, pubKey, p.sessionStorage, p.aiEngine)

	morTcpHandler := proxyapi.NewMORRPCController(proxyReceiver, p.validator, p.sessionStorage)
	tcpHandler := tcphandlers.NewTCPHandler(
		p.log, p.connLog, p.schedulerLogFactory, morTcpHandler,
	)
	tcpServer.SetConnectionHandler(tcpHandler)

	g, errCtx := errgroup.WithContext(ctx)
	g.Go(func() error {
		return tcpServer.Run(errCtx)
	})

	g.Go(func() error {
		return p.eventListener.Run(errCtx)
	})

	return g.Wait()
}

func (p *Proxy) GetState() ProxyState {
	return p.state.Load()
}

func (p *Proxy) setState(s ProxyState) {
	p.state.Store(s)
	p.log.Infof("proxy state: %s", s)
}
