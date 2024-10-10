package contracts

import (
	"context"
	"math/big"
	"time"

	i "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/interfaces"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type LogWatcherPolling struct {
	// config
	maxReconnects int
	pollInterval  time.Duration

	// deps
	client i.EthClient
	log    lib.ILogger
}

func NewLogWatcherPolling(client i.EthClient, pollInterval time.Duration, maxReconnects int, log lib.ILogger) *LogWatcherPolling {
	return &LogWatcherPolling{
		client:        client,
		pollInterval:  pollInterval,
		maxReconnects: maxReconnects,
		log:           log,
	}
}

func (w *LogWatcherPolling) Watch(ctx context.Context, contractAddr common.Address, mapper EventMapper, fromBlock *big.Int) (*lib.Subscription, error) {
	if fromBlock == nil {
		block, err := w.client.HeaderByNumber(ctx, nil)
		if err != nil {
			return nil, err
		}
		fromBlock = block.Number
	}
	lastQueriedBlock := fromBlock

	sink := make(chan interface{})
	return lib.NewSubscription(func(quit <-chan struct{}) error {
		defer close(sink)

		for {
			query := ethereum.FilterQuery{
				Addresses: []common.Address{contractAddr},
				FromBlock: lastQueriedBlock,
				ToBlock:   nil,
			}
			sub, err := w.filterLogsRetry(ctx, query)
			if err != nil {
				return err
			}

			for _, log := range sub {
				if log.Removed {
					continue
				}
				event, err := mapper(log)
				if err != nil {
					return err // mapper error, retry won't help
				}

				select {
				case <-quit:
					return nil
				case <-ctx.Done():
					return ctx.Err()
				case sink <- event:
				}
			}

			if len(sub) > 0 {
				lastQueriedBlock = new(big.Int).SetUint64(sub[len(sub)-1].BlockNumber)
			}

			select {
			case <-quit:
				return nil
			case <-ctx.Done():
				return ctx.Err()
			case <-time.After(w.pollInterval):
			}
		}
	}, sink), nil
}

func (w *LogWatcherPolling) filterLogsRetry(ctx context.Context, query ethereum.FilterQuery) ([]types.Log, error) {
	var lastErr error

	for attempts := 0; attempts < w.maxReconnects; attempts++ {
		logs, err := w.client.FilterLogs(ctx, query)
		if err != nil {
			lastErr = err
			continue
		}
		if attempts > 0 {
			w.log.Warnf("subscription reconnected due to error: %s", lastErr)
		}

		return logs, nil
	}

	return nil, lastErr
}
