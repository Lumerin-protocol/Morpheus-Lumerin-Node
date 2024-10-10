package ethclient

import (
	"context"
	"sync"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/ethereum/go-ethereum/rpc"
)

// Wrapper around multiple RPC clients, used to retry calls on multiple endpoints
type RPCClientMultiple struct {
	lock    sync.RWMutex
	clients []*rpcClient
	log     lib.ILogger
}

func NewRPCClientMultiple(urls []string, log lib.ILogger) (*RPCClientMultiple, error) {
	clients := make([]*rpcClient, len(urls))

	for i, url := range urls {
		client, err := rpc.DialOptions(context.Background(), url)
		if err != nil {
			return nil, err
		}
		clients[i] = &rpcClient{
			url:    url,
			client: client,
		}
	}

	return &RPCClientMultiple{clients: clients, log: log}, nil
}

func (c *RPCClientMultiple) CallContext(ctx context.Context, result interface{}, method string, args ...interface{}) error {
	return c.retriableCall(ctx, func(client *rpcClient) error {
		return client.client.CallContext(ctx, result, method, args...)
	})
}

func (c *RPCClientMultiple) BatchCallContext(ctx context.Context, b []rpc.BatchElem) error {
	return c.retriableCall(ctx, func(client *rpcClient) error {
		return client.client.BatchCallContext(ctx, b)
	})
}

func (c *RPCClientMultiple) Close() {
	for _, rpcClient := range c.getClients() {
		rpcClient.client.Close()
	}
}

func (c *RPCClientMultiple) EthSubscribe(ctx context.Context, channel interface{}, args ...interface{}) (*rpc.ClientSubscription, error) {
	client := c.getClients()[0]
	return client.client.EthSubscribe(ctx, channel, args...)
}

func (c *RPCClientMultiple) GetURLs() []string {
	clients := c.getClients()
	urls := make([]string, len(clients))
	for i, rpcClient := range clients {
		urls[i] = rpcClient.url
	}
	return urls
}

func (c *RPCClientMultiple) SetURLs(urls []string) error {
	clients := make([]*rpcClient, len(urls))

	for i, url := range urls {
		client, err := rpc.DialOptions(context.Background(), url)
		if err != nil {
			return err
		}
		clients[i] = &rpcClient{
			url:    url,
			client: client,
		}
	}

	c.lock.Lock()
	defer c.lock.Unlock()
	for _, rpcClient := range c.clients {
		rpcClient.client.Close()
	}

	c.clients = clients

	return nil
}

func (c *RPCClientMultiple) getClients() []*rpcClient {
	c.lock.RLock()
	defer c.lock.RUnlock()
	return c.clients
}

// retriableCall is a helper function that retries the call on different endpoints
func (c *RPCClientMultiple) retriableCall(ctx context.Context, fn func(client *rpcClient) error) error {
	var lastErr error

	for _, rpcClient := range c.getClients() {
		if ctx.Err() != nil {
			return ctx.Err()
		}

		c.log.Debugf("calling eth endpoint %s", rpcClient.url)
		err := fn(rpcClient)
		if err == nil {
			return nil
		}

		retryable := c.shouldBeRetried(err)
		c.log.Debugf("error (retryable: %t) calling eth endpoint %s: %s", retryable, rpcClient.url, err)
		if !retryable {
			return err
		}

		lastErr = err
	}

	c.log.Debugf("all endpoints failed")
	return lastErr
}

func (c *RPCClientMultiple) shouldBeRetried(err error) bool {
	switch err.(type) {
	case rpc.HTTPError:
		// if err.(rpc.HTTPError).StatusCode == 429 {
		// 	return true
		// }
		return true
	case JSONError:
		return false
	}
	return false
}

type JSONError interface {
	Error() string
	ErrorCode() int
	ErrorData() interface{}
}

type rpcClient struct {
	url    string
	client *rpc.Client
}
