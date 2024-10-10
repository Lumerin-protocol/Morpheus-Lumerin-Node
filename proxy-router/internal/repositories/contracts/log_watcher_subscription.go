package contracts

import (
	"context"
	"math/big"

	i "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/interfaces"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type LogWatcherSubscription struct {
	// config
	maxReconnects int

	// deps
	client i.EthClient
	log    lib.ILogger
}

// NewLogWatcherSubscription creates a new log subscription using websocket
// TODO: if it is going to be primary implementation we should rewrite it so it doesn't skip events in case of temporary downtime
func NewLogWatcherSubscription(client i.EthClient, maxReconnects int, log lib.ILogger) *LogWatcherSubscription {
	return &LogWatcherSubscription{
		maxReconnects: maxReconnects,
		client:        client,
		log:           log,
	}
}

func (w *LogWatcherSubscription) Watch(ctx context.Context, contractAddr common.Address, mapper EventMapper, fromBlock *big.Int) (*lib.Subscription, error) {
	sink := make(chan interface{})

	return lib.NewSubscription(func(quit <-chan struct{}) error {
		defer close(sink)

		query := ethereum.FilterQuery{
			Addresses: []common.Address{contractAddr},
		}
		in := make(chan types.Log)
		defer close(in)

		for {

			sub, err := w.subscribeFilterLogsRetry(ctx, query, in)
			if err != nil {
				w.log.Errorf("failed to subscribe to logs: %s", err)
				return err
			}

			defer sub.Unsubscribe()

		EVENTS_LOOP:
			for {
				select {
				case log := <-in:
					event, err := mapper(log)
					if err != nil {
						w.log.Errorf("failed to map event: %s", err)
						// mapper error, retry won't help
						return err
					}

					select {
					case sink <- event:
					case err := <-sub.Err():
						w.log.Debugf("subscription error: %s", err)
						break EVENTS_LOOP
					case <-quit:
						return nil
					case <-ctx.Done():
						return ctx.Err()
					}
				case err := <-sub.Err():
					w.log.Debugf("subscription error: %s", err)
					break EVENTS_LOOP
				case <-quit:
					return nil
				case <-ctx.Done():
					return ctx.Err()
				}
			}
		}
	}, sink), nil
}

func (w *LogWatcherSubscription) subscribeFilterLogsRetry(ctx context.Context, query ethereum.FilterQuery, ch chan<- types.Log) (ethereum.Subscription, error) {
	var lastErr error

	for attempts := 0; attempts < w.maxReconnects; attempts++ {
		sub, err := w.client.SubscribeFilterLogs(ctx, query, ch)
		if err != nil {
			lastErr = err
			continue
		}
		if attempts > 0 {
			w.log.Warnf("subscription successfully reconnected after error: %s", lastErr)
		}

		return sub, nil
	}

	return nil, lastErr
}
