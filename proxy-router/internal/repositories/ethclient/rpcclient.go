package ethclient

import (
	"context"

	"github.com/ethereum/go-ethereum/rpc"
)

type RPCClient interface {
	Close()
	BatchCallContext(context.Context, []rpc.BatchElem) error
	CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error
	EthSubscribe(ctx context.Context, channel interface{}, args ...interface{}) (*rpc.ClientSubscription, error)
}

type RPCClientModifiable interface {
	RPCClient
	GetURLs() []string
	SetURLs(urls []string) error
}
