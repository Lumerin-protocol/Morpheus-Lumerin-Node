package blockchainapi

import (
	"context"

	"github.com/Lumerin-protocol/Morpheus-Lumerin-Node/proxy-router/contracts/sessionrouter"
	"github.com/Lumerin-protocol/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/Lumerin-protocol/Morpheus-Lumerin-Node/proxy-router/internal/repositories/contracts"
	"github.com/Lumerin-protocol/Morpheus-Lumerin-Node/proxy-router/internal/repositories/registries"
	"github.com/Lumerin-protocol/Morpheus-Lumerin-Node/proxy-router/internal/storages"
	"github.com/ethereum/go-ethereum/ethclient"
)

type EventsListener struct {
	sessionRouter *registries.SessionRouter
	store         *storages.SessionStorage
	tsk           *lib.Task
	log           *lib.Logger
	client        *ethclient.Client
}

func NewEventsListener(client *ethclient.Client, store *storages.SessionStorage, sessionRouter *registries.SessionRouter, log *lib.Logger) *EventsListener {
	return &EventsListener{
		store:         store,
		log:           log,
		sessionRouter: sessionRouter,
		client:        client,
	}
}

func (e *EventsListener) Run(ctx context.Context) error {
	defer func() {
		_ = e.log.Close()
	}()

	sub, err := contracts.WatchContractEvents(ctx, e.client, e.sessionRouter.GetContractAddress(), contracts.CreateEventMapper(contracts.BlockchainEventFactory, e.sessionRouter.GetABI()), e.log)
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
			// return err
		}
	}
}

func (e *EventsListener) controller(event interface{}) error {
	switch ev := event.(type) {
	case *sessionrouter.SessionRouterSessionOpened:
		return e.handleSessionOpened(ev)
	}
	return nil
}

func (e *EventsListener) handleSessionOpened(event *sessionrouter.SessionRouterSessionOpened) error {
	sessionId := lib.BytesToString(event.SessionId[:])
	e.log.Debugf("received open session router event, sessionId %s", sessionId)

	// TDO: pull session from bc
	// add
	err := e.store.AddSession(&storages.Session{
		Id:       sessionId,
		UserAddr: event.UserAddress.Hex(),
		// add more fields endsAt
	})
	if err != nil {
		return err
	}

	return nil
}
