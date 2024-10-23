package registries

import (
	"context"
	"fmt"

	i "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/interfaces"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/contracts/bindings/providerregistry"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type ProviderRegistry struct {
	// config
	providerRegistryAddr common.Address

	// state
	nonce uint64

	// deps
	providerRegistry *providerregistry.ProviderRegistry
	client           i.ContractBackend
	log              lib.ILogger
}

func NewProviderRegistry(providerRegistryAddr common.Address, client i.ContractBackend, log lib.ILogger) *ProviderRegistry {
	pr, err := providerregistry.NewProviderRegistry(providerRegistryAddr, client)
	if err != nil {
		panic("invalid provider registry ABI")
	}
	return &ProviderRegistry{
		providerRegistry:     pr,
		providerRegistryAddr: providerRegistryAddr,
		client:               client,
		log:                  log,
	}
}

func (g *ProviderRegistry) GetAllProviders(ctx context.Context) ([]common.Address, []providerregistry.IProviderStorageProvider, error) {
	// providerAddrs, providers, err := g.providerRegistry.ProviderGetAll(&bind.CallOpts{Context: ctx})
	// if err != nil {
	// 	return nil, nil, err
	// }

	// addresses := make([]common.Address, len(providerAddrs))
	// for i, address := range providerAddrs {
	// 	addresses[i] = address
	// }

	// return addresses, providers, nil
	return nil, nil, fmt.Errorf("Not implemented")
}

func (g *ProviderRegistry) CreateNewProvider(opts *bind.TransactOpts, addStake *lib.BigInt, endpoint string) error {
	providerTx, err := g.providerRegistry.ProviderRegister(opts, &addStake.Int, endpoint)

	if err != nil {
		return lib.TryConvertGethError(err)
	}

	// Wait for the transaction receipt
	receipt, err := bind.WaitMined(opts.Context, g.client, providerTx)
	if err != nil {
		return lib.TryConvertGethError(err)
	}

	if receipt.Status != 1 {
		return fmt.Errorf("Transaction failed with status %d", receipt.Status)
	}

	return nil
}

func (g *ProviderRegistry) DeregisterProvider(opts *bind.TransactOpts) (common.Hash, error) {
	providerTx, err := g.providerRegistry.ProviderDeregister(opts)

	if err != nil {
		return common.Hash{}, lib.TryConvertGethError(err)
	}

	// Wait for the transaction receipt
	receipt, err := bind.WaitMined(opts.Context, g.client, providerTx)
	if err != nil {
		return common.Hash{}, lib.TryConvertGethError(err)
	}

	if receipt.Status != 1 {
		return receipt.TxHash, fmt.Errorf("Transaction failed with status %d", receipt.Status)
	}

	return receipt.TxHash, nil
}

func (g *ProviderRegistry) GetProviderById(ctx context.Context, id common.Address) (*providerregistry.IProviderStorageProvider, error) {
	provider, err := g.providerRegistry.GetProvider(&bind.CallOpts{Context: ctx}, id)
	if err != nil {
		return nil, err
	}

	return &provider, nil
}
