package registries

import (
	"context"
	"fmt"
	"math/big"

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

func (g *ProviderRegistry) CreateNewProvider(ctx *bind.TransactOpts, address string, addStake uint64, endpoint string) error {
	bigAddStake := big.NewInt(int64(addStake))

	providerTx, err := g.providerRegistry.ProviderRegister(ctx, common.HexToAddress(address), bigAddStake, endpoint)

	if err != nil {
		return lib.TryConvertGethError(err, providerregistry.ProviderRegistryMetaData)
	}

	// Wait for the transaction receipt
	receipt, err := bind.WaitMined(context.Background(), g.client, providerTx)
	if err != nil {
		return lib.TryConvertGethError(err, providerregistry.ProviderRegistryMetaData)
	}

	// Find the event log
	for _, log := range receipt.Logs {
		// Check if the log belongs to the OpenSession event
		_, err := g.providerRegistry.ParseProviderRegisteredUpdated(*log)

		if err != nil {
			continue // not our event, skip it
		}

		return nil
	}

	return fmt.Errorf("OpenSession event not found in transaction logs")
}

func (g *ProviderRegistry) GetProviderById(ctx context.Context, id common.Address) (*providerregistry.Provider, error) {
	provider, err := g.providerRegistry.ProviderMap(&bind.CallOpts{Context: ctx}, id)
	if err != nil {
		return nil, err
	}

	return &provider, nil
}
