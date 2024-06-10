package registries

import (
	"context"
	"fmt"
	"math/big"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/contracts/providerregistry"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/internal/interfaces"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/internal/lib"
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
	mutex lib.Mutex
	prABI *abi.ABI

	// deps
	providerRegistry *providerregistry.ProviderRegistry
	client           *ethclient.Client
	log              interfaces.ILogger
}

func NewProviderRegistry(providerRegistryAddr common.Address, client *ethclient.Client, log interfaces.ILogger) *ProviderRegistry {
	pr, err := providerregistry.NewProviderRegistry(providerRegistryAddr, client)
	if err != nil {
		panic("invalid provider registry ABI")
	}
	prABI, err := providerregistry.ProviderRegistryMetaData.GetAbi()
	if err != nil {
		panic("invalid provider registry ABI: " + err.Error())
	}
	fmt.Println(providerRegistryAddr)
	return &ProviderRegistry{
		providerRegistry:     pr,
		providerRegistryAddr: providerRegistryAddr,
		client:               client,
		prABI:                prABI,
		mutex:                lib.NewMutex(),
		log:                  log,
	}
}

func (g *ProviderRegistry) GetAllProviders(ctx context.Context) ([]string, []providerregistry.Provider, error) {
	providerAddrs, providers, err := g.providerRegistry.ProviderGetAll(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, nil, err
	}

	addresses := make([]string, len(providerAddrs))
	for i, address := range providerAddrs {
		addresses[i] = address.Hex()
	}

	return addresses, providers, nil
}

func (g *ProviderRegistry) CreateNewProvider(ctx context.Context, address string, addStake uint64, endpoint string) ( error) {

	bigAddStake := big.NewInt(int64(addStake))

	providerTx, err := g.providerRegistry.ProviderRegister(&bind.TransactOpts{Context: ctx}, common.HexToAddress(address), bigAddStake, endpoint)
	
	if err != nil {
		return  err
	}

	// Wait for the transaction receipt
	receipt, err := bind.WaitMined(context.Background(), g.client, providerTx)
	if err != nil {
		return err
	}

	// Find the event log
	for _, log := range receipt.Logs {
		// Check if the log belongs to the OpenSession event
		_, err := g.providerRegistry.ParseProviderRegisteredUpdated(*log)
		
		if err != nil {
			continue // not our event, skip it
		}
		
		return  nil
	}

	return fmt.Errorf("OpenSession event not found in transaction logs")
}
