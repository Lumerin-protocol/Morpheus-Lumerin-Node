package contracts

import (
	"context"
	"errors"
	"time"

	"github.com/Lumerin-protocol/Morpheus-Lumerin-Node/proxy-router/contracts/sessionrouter"
	"github.com/Lumerin-protocol/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/contracts/sessionrouter"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

const RECONNECT_TIMEOUT = 2 * time.Second

type EventMapper func(types.Log) (interface{}, error)

func BlockchainEventFactory(name string) interface{} {
	switch name {
	case "SessionOpened":
		return new(sessionrouter.SessionRouterSessionOpened)
	default:
		return nil
	}
}

// WatchContractEvents watches for all events from the contract and converts them to the concrete type, using mapper
func WatchContractEvents(ctx context.Context, client EthereumClient, contractAddr common.Address, mapper EventMapper, log lib.ILogger) (*lib.Subscription, error) {
	sink := make(chan interface{})

	return lib.NewSubscription(func(quit <-chan struct{}) error {
		defer close(sink)

		query := ethereum.FilterQuery{
			Addresses: []common.Address{contractAddr},
		}
		in := make(chan types.Log)
		defer close(in)

		var lastErr error

		for attempts := 0; true; attempts++ {
			sub, err := client.SubscribeFilterLogs(ctx, query, in)
			if err != nil {
				lastErr = err

				log.Warnf("subscription error, reconnect in %s: %s", RECONNECT_TIMEOUT, lastErr)

				select {
				case <-quit:
					return nil
				case <-ctx.Done():
					return ctx.Err()
				case <-time.After(RECONNECT_TIMEOUT):
				}

				continue
			}
			if attempts > 0 {
				log.Warnf("subscription reconnected due to error: %s", lastErr)
			}
			attempts = 0

			defer sub.Unsubscribe()

		EVENTS_LOOP:
			for {
				select {
				case logEntry := <-in:
					event, err := mapper(logEntry)
					if err != nil {

						if errors.Is(err, ErrUnknownEvent) {
							log.Warnf("unknown event: %s", err)
							continue
						}
						// mapper error, retry won't help
						// return err
						continue
					}

					select {
					case sink <- event:
					case err := <-sub.Err():
						lastErr = err
						break EVENTS_LOOP
					case <-quit:
						return nil
					case <-ctx.Done():
						return ctx.Err()
					}
				case err := <-sub.Err():
					lastErr = err
					break EVENTS_LOOP
				case <-quit:
					return nil
				case <-ctx.Done():
					return ctx.Err()
				}
			}

			log.Warnf("subscription error, reconnect in %s: %s", RECONNECT_TIMEOUT, lastErr)

			select {
			case <-quit:
				return nil
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(RECONNECT_TIMEOUT):
			}
		}

		return lastErr
	}, sink), nil
}
