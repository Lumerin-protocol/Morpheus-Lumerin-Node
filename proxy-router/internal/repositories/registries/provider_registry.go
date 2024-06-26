package registries

import (
	"context"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/contracts/providerregistry"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ProviderRegistry struct {
	// config
	providerRegistryAddr common.Address

	// state
	nonce uint64
	prABI *abi.ABI

	// deps
	providerRegistry *providerregistry.ProviderRegistry
	client           *ethclient.Client
	log              lib.ILogger
}

func NewProviderRegistry(providerRegistryAddr common.Address, client *ethclient.Client, log lib.ILogger) *ProviderRegistry {
	pr, err := providerregistry.NewProviderRegistry(providerRegistryAddr, client)
	if err != nil {
		panic("invalid provider registry ABI")
	}
	prABI, err := providerregistry.ProviderRegistryMetaData.GetAbi()
	if err != nil {
		panic("invalid provider registry ABI: " + err.Error())
	}
	return &ProviderRegistry{
		providerRegistry:     pr,
		providerRegistryAddr: providerRegistryAddr,
		client:               client,
		prABI:                prABI,
		log:                  log,
	}
}

func (g *ProviderRegistry) GetAllProviders(ctx context.Context) ([]common.Address, []providerregistry.Provider, error) {
	providerAddrs, providers, err := g.providerRegistry.ProviderGetAll(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, nil, err
	}

	addresses := make([]common.Address, len(providerAddrs))
	for i, address := range providerAddrs {
		addresses[i] = address
	}

	return addresses, providers, nil
}

func (g *ProviderRegistry) GetProviderById(ctx context.Context, id common.Address) (*providerregistry.Provider, error) {
	provider, err := g.providerRegistry.ProviderMap(&bind.CallOpts{Context: ctx}, id)
	if err != nil {
		return nil, err
	}

	return &provider, nil
}
