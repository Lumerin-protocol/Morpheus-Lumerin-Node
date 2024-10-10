package ethclient

import (
	"encoding/json"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/interfaces"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
)

const (
	ETH_NODE_URL_KEY = "eth-node-url"
)

type RPCClientStoreKeychain struct {
	storage   interfaces.KeyValueStorage
	rpcClient RPCClientModifiable
	log       lib.ILogger
}

func NewRPCClientStoreKeychain(storage interfaces.KeyValueStorage, rpcClient RPCClientModifiable, log lib.ILogger) *RPCClientStoreKeychain {
	return &RPCClientStoreKeychain{
		storage:   storage,
		rpcClient: rpcClient,
		log:       log,
	}
}

func (p *RPCClientStoreKeychain) GetURLs() []string {
	return p.rpcClient.GetURLs()
}

func (p *RPCClientStoreKeychain) SetURLs(urls []string) error {
	err := p.storeURLsInStorage(urls)
	if err != nil {
		return err
	}

	return p.rpcClient.SetURLs(urls)
}

func (p *RPCClientStoreKeychain) RemoveURLs() error {
	return p.deleteURLsInStorage()
}

func (p *RPCClientStoreKeychain) GetClient() RPCClient {
	return p.rpcClient
}

func (p *RPCClientStoreKeychain) loadURLsFromStorage() ([]string, error) {
	// return []string{"https://arb-sepolia.g.alchemy.com/v2/3-pxwBaJ7vilkz1jl-fMmCvZThGxpmo2"}, nil
	str, err := p.storage.Get(ETH_NODE_URL_KEY)
	if err != nil {
		return nil, err
	}

	var urls []string
	err = json.Unmarshal([]byte(str), &urls)
	if err != nil {
		return nil, err
	}

	return urls, nil
}

func (p *RPCClientStoreKeychain) storeURLsInStorage(urls []string) error {
	str, err := json.Marshal(urls)
	if err != nil {
		return err
	}

	err = p.storage.Upsert(ETH_NODE_URL_KEY, string(str))
	if err != nil {
		return err
	}

	return nil
}

func (p *RPCClientStoreKeychain) deleteURLsInStorage() error {
	return p.storage.DeleteIfExists(ETH_NODE_URL_KEY)
}
