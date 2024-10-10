package ethclient

import (
	"errors"
)

var ErrEnvRPCSet = errors.New("cannot set rpc url when using env store, switch to keychain store by removing ETH_NODE_URL env var")

type RPCClientStoreEnv struct {
	rpcClient RPCClientModifiable
}

func (p *RPCClientStoreEnv) GetURLs() []string {
	return p.rpcClient.GetURLs()
}

func (p *RPCClientStoreEnv) SetURLs(urls []string) error {
	return ErrEnvRPCSet
}

func (p *RPCClientStoreEnv) RemoveURLs() error {
	return ErrEnvRPCSet
}

func (p *RPCClientStoreEnv) GetClient() RPCClient {
	return p.rpcClient
}
