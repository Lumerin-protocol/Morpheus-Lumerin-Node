package contracts

import (
	"context"
	"errors"
	"fmt"
	"math/big"
	"testing"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/interfaces/mocks"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/ethereum/go-ethereum"
	common "github.com/ethereum/go-ethereum/common"
	types "github.com/ethereum/go-ethereum/core/types"
	"github.com/stretchr/testify/mock"
	"github.com/stretchr/testify/require"
)

var TEST_ERR = errors.New("test error")

func TestLogWatcherPolling(t *testing.T) {
	failTimes := 5

	ethClientMock := mocks.NewEthClientMock(t)
	call1 := ethClientMock.EXPECT().FilterLogs(mock.Anything, mock.Anything).Return(nil, TEST_ERR).Times(failTimes)
	_ = ethClientMock.EXPECT().FilterLogs(mock.Anything, mock.Anything).Return([]types.Log{}, nil).Times(1).NotBefore(call1)
	logWatcherPolling := NewLogWatcherPolling(ethClientMock, 0, 10, lib.NewTestLogger())

	_, err := logWatcherPolling.filterLogsRetry(context.Background(), ethereum.FilterQuery{}, make(<-chan struct{}))
	require.NoError(t, err)
	ethClientMock.AssertNumberOfCalls(t, "FilterLogs", failTimes+1)
}

func TestWatchDoesntReturnEventsTwice(t *testing.T) {
	ethClientMock := mocks.NewEthClientMock(t)
	event1 := types.Log{
		BlockNumber: 2,
		Index:       1,
		Data:        []byte{2},
	}
	event2 := types.Log{
		BlockNumber: 3,
		Index:       1,
		Data:        []byte{3},
	}

	_ = ethClientMock.EXPECT().FilterLogs(mock.Anything, matchBlockNumber(1)).Return([]types.Log{event1}, nil)
	_ = ethClientMock.EXPECT().FilterLogs(mock.Anything, matchBlockNumber(2)).Return([]types.Log{event2}, nil)
	_ = ethClientMock.EXPECT().FilterLogs(mock.Anything, mock.Anything).Return([]types.Log{}, nil)

	logWatcherPolling := NewLogWatcherPolling(ethClientMock, 0, 10, lib.NewTestLogger())
	sub, err := logWatcherPolling.Watch(context.Background(), common.Address{}, eventMapper, big.NewInt(1))
	require.NoError(t, err)
	defer sub.Unsubscribe()

	totalEvents := 0
	uniqueEvents := make(map[string]types.Log)

OUTER:
	for {
		select {
		case e := <-sub.Events():
			event := e.(types.Log)
			totalEvents++
			uniqueEvents[fmt.Sprintf("%d-%d", event.BlockNumber, event.Index)] = event
			if len(uniqueEvents) == 2 {
				break OUTER
			}
		case <-sub.Err():
			require.NoError(t, err)
		}
	}

	require.Equal(t, totalEvents, len(uniqueEvents))
}

func TestWatchShouldErrorAfterMaxReconnects(t *testing.T) {
	ethClientMock := mocks.NewEthClientMock(t)
	maxRetries := 10

	_ = ethClientMock.EXPECT().FilterLogs(mock.Anything, mock.Anything).Return([]types.Log{}, TEST_ERR)

	logWatcherPolling := NewLogWatcherPolling(ethClientMock, 0, maxRetries, lib.NewTestLogger())
	sub, err := logWatcherPolling.Watch(context.Background(), common.Address{}, eventMapper, big.NewInt(1))
	require.NoError(t, err)
	defer sub.Unsubscribe()

	for {
		select {
		case e := <-sub.Events():
			if e != nil {
				require.Fail(t, "should not receive events")
			}
		case err := <-sub.Err():
			require.ErrorIs(t, err, TEST_ERR)
			ethClientMock.AssertNumberOfCalls(t, "FilterLogs", maxRetries)
			return
		}
	}
}

func TestShouldHandleContextCancellation(t *testing.T) {
	ethClientMock := mocks.NewEthClientMock(t)
	ctx, cancel := context.WithCancel(context.Background())

	_ = ethClientMock.EXPECT().FilterLogs(mock.Anything, mock.Anything).Return([]types.Log{}, nil)

	logWatcherPolling := NewLogWatcherPolling(ethClientMock, 0, 10, lib.NewTestLogger())
	sub, err := logWatcherPolling.Watch(ctx, common.Address{}, eventMapper, big.NewInt(1))
	require.NoError(t, err)
	defer sub.Unsubscribe()

	cancel()

	for {
		select {
		case e := <-sub.Events():
			if e != nil {
				require.Fail(t, "should not receive events")
			}
		case err := <-sub.Err():
			require.ErrorIs(t, err, context.Canceled)
			return
		}
	}
}

func TestShouldUnsubscribe(t *testing.T) {
	ethClientMock := mocks.NewEthClientMock(t)
	_ = ethClientMock.EXPECT().FilterLogs(mock.Anything, mock.Anything).Return([]types.Log{}, nil)

	logWatcherPolling := NewLogWatcherPolling(ethClientMock, 0, 10, lib.NewTestLogger())
	sub, err := logWatcherPolling.Watch(context.Background(), common.Address{}, eventMapper, big.NewInt(1))
	require.NoError(t, err)

	sub.Unsubscribe()

	for {
		select {
		case e := <-sub.Events():
			if e != nil {
				require.Fail(t, "should not receive events")
			}
		case err := <-sub.Err():
			require.Nil(t, err)
			return
		}
	}
}

func matchBlockNumber(blockNumber int) any {
	return mock.MatchedBy(func(query ethereum.FilterQuery) bool {
		return int(query.FromBlock.Int64()) == blockNumber
	})
}

func eventMapper(log types.Log) (interface{}, error) {
	return log, nil
}
