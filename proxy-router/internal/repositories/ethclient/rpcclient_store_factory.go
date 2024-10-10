package ethclient

import (
	"errors"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/interfaces"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/keychain"
)

type RPCEndpointsPersister interface {
	GetURLs() []string
	SetURLs(urls []string) error
	RemoveURLs() error
	GetClient() RPCClient
}

func ConfigureRPCClientStore(storage interfaces.KeyValueStorage, envURLs []string, chainID int, log lib.ILogger) (RPCEndpointsPersister, error) {
	// if env set, use env store
	if len(envURLs) > 0 {
		rpcClient, err := NewRPCClientMultiple(envURLs, log)
		if err != nil {
			return nil, err
		}

		log.Info("using eth node address configured in env")

		p := &RPCClientStoreEnv{
			rpcClient: rpcClient,
		}
		return p, nil
	}

	// if no env set, try use keychain store
	rpcClient := &RPCClientMultiple{
		log: log,
	}
	p := &RPCClientStoreKeychain{
		storage:   storage,
		log:       log,
		rpcClient: rpcClient,
	}

	urls, err := p.loadURLsFromStorage()
	if err == nil {
		log.Info("using eth node address configured in keychain")

		err = rpcClient.SetURLs(urls)
		if err != nil {
			return nil, err
		}

		return p, nil
	}

	// if error during loading keychain, use fallback URLs
	if !errors.Is(err, keychain.ErrKeyNotFound) {
		p.log.Warn("Error during loading keychain eth client URLs, using fallback URLs", err)
	}

	if chainID == 0 {
		return nil, errors.New("Missing chainID. You need to set CHAIN_ID env variable to use public eth node addresses")
	}

	publicURLs, err := GetPublicRPCURLs(chainID)
	if err != nil {
		return nil, err
	}

	log.Info("using public eth node addresses")

	rpcClient, err = NewRPCClientMultiple(publicURLs, log)
	if err != nil {
		return nil, err
	}

	rpc := &RPCClientStoreKeychain{
		rpcClient: rpcClient,
		storage:   storage,
	}

	return rpc, nil
}
