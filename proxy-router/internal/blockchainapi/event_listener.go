package blockchainapi

import (
	"context"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/config"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/interfaces"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/contracts"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/contracts/bindings/sessionrouter"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/registries"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/storages"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
)

type EventsListener struct {
	sessionRouter     *registries.SessionRouter
	marketplace       *registries.Marketplace
	store             *storages.SessionStorage
	tsk               *lib.Task
	log               *lib.Logger
	client            bind.ContractFilterer
	wallet            interfaces.Wallet
	modelConfigLoader *config.ModelConfigLoader
	logWatcher        contracts.LogWatcher
}

func NewEventsListener(client bind.ContractFilterer, store *storages.SessionStorage, sessionRouter *registries.SessionRouter, marketplace *registries.Marketplace, wallet interfaces.Wallet, modelConfigLoader *config.ModelConfigLoader, logWatcher contracts.LogWatcher, log *lib.Logger) *EventsListener {
	return &EventsListener{
		store:             store,
		log:               log,
		sessionRouter:     sessionRouter,
		marketplace:       marketplace,
		client:            client,
		wallet:            wallet,
		modelConfigLoader: modelConfigLoader,
		logWatcher:        logWatcher,
	}
}

func (e *EventsListener) Run(ctx context.Context) error {
	defer func() {
		_ = e.log.Close()
	}()

	sub, err := e.logWatcher.Watch(ctx, e.sessionRouter.GetContractAddress(), contracts.CreateEventMapper(contracts.SessionRouterEventFactory, e.sessionRouter.GetABI()), nil)
	if err != nil {
		return err
	}
	defer sub.Unsubscribe()
	e.log.Infof("started watching events, address %s", e.sessionRouter.GetContractAddress())

	for {
		select {
		case <-ctx.Done():
			return ctx.Err()
		case event := <-sub.Events():
			err := e.controller(event)
			if err != nil {
				e.log.Errorf("error loading data: %s", err)
			}
		case err := <-sub.Err():
			e.log.Errorf("error in event listener: %s", err)
			return err
		}
	}
}

func (e *EventsListener) controller(event interface{}) error {
	switch ev := event.(type) {
	case *sessionrouter.SessionRouterSessionOpened:
		return e.handleSessionOpened(ev)
	case *sessionrouter.SessionRouterSessionClosed:
		return e.handleSessionClosed(ev)
	}
	return nil
}

func (e *EventsListener) handleSessionOpened(event *sessionrouter.SessionRouterSessionOpened) error {
	ctx := context.Background()

	sessionId := lib.BytesToString(event.SessionId[:])
	e.log.Debugf("received open session router event, sessionId %s", sessionId)

	session, err := e.sessionRouter.GetSession(ctx, event.SessionId)
	if err != nil {
		e.log.Errorf("failed to get session from blockchain: %s, sessionId %s", err, sessionId)
		return err
	}

	bid, err := e.marketplace.GetBidById(ctx, session.BidId)
	if err != nil {
		e.log.Errorf("failed to get bid from blockchain: %s, sessionId %s", err, sessionId)
		return err
	}

	privateKey, err := e.wallet.GetPrivateKey()
	if err != nil {
		e.log.Errorf("failed to get private key: %s", err)
		return err
	}

	address, err := lib.PrivKeyBytesToAddr(privateKey)
	if err != nil {
		e.log.Errorf("failed to get address from private key: %s", err)
		return err
	}

	if bid.Provider.Hex() != address.Hex() && event.User.Hex() != address.Hex() {
		e.log.Debugf("session provider/user is not me, skipping, sessionId %s", sessionId)
		return nil
	}

	modelID := lib.BytesToString(bid.ModelId[:])

	var modelConfig *config.ModelConfig
	if bid.Provider.Hex() == address.Hex() {
		modelConfig = e.modelConfigLoader.ModelConfigFromID(modelID)
		err = e.store.AddSessionToModel(modelID, sessionId)
		if err != nil {
			return err
		}
	} else {
		modelConfig = &config.ModelConfig{}
	}

	err = e.store.AddSession(&storages.Session{
		Id:           sessionId,
		UserAddr:     event.User.Hex(),
		ProviderAddr: bid.Provider.Hex(),
		EndsAt:       session.EndsAt,
		ModelID:      modelID,
		ModelName:    modelConfig.ModelName,
		ModelApiType: modelConfig.ApiType,
	})
	if err != nil {
		return err
	}

	return nil
}

func (e *EventsListener) handleSessionClosed(event *sessionrouter.SessionRouterSessionClosed) error {
	sessionId := lib.BytesToString(event.SessionId[:])
	e.log.Debugf("received close session router event, sessionId %s", sessionId)

	_, ok := e.store.GetSession(sessionId)
	if !ok {
		e.log.Debugf("session not found in storage, sessionId %s", sessionId)
		return nil
	}

	err := e.store.RemoveSession(sessionId)
	if err != nil {
		e.log.Errorf("failed to remove session from storage: %s, sessionId %s", err, sessionId)
		return err
	}

	return nil
}
