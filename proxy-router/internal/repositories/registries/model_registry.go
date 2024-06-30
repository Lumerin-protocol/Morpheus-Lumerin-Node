package registries

import (
	"context"

	"github.com/Lumerin-protocol/Morpheus-Lumerin-Node/proxy-router/contracts/modelregistry"
	"github.com/Lumerin-protocol/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type ModelRegistry struct {
	// config
	modelRegistryAddr common.Address

	// state
	nonce uint64
	mrABI *abi.ABI

	// deps
	modelRegistry *modelregistry.ModelRegistry
	client        *ethclient.Client
	log           lib.ILogger
}

func NewModelRegistry(modelRegistryAddr common.Address, client *ethclient.Client, log lib.ILogger) *ModelRegistry {
	mr, err := modelregistry.NewModelRegistry(modelRegistryAddr, client)
	if err != nil {
		panic("invalid model registry ABI")
	}
	mrABI, err := modelregistry.ModelRegistryMetaData.GetAbi()
	if err != nil {
		panic("invalid model registry ABI: " + err.Error())
	}
	return &ModelRegistry{
		modelRegistry:     mr,
		modelRegistryAddr: modelRegistryAddr,
		client:            client,
		mrABI:             mrABI,
		log:               log,
	}
}

func (g *ModelRegistry) GetAllModels(ctx context.Context) ([][32]byte, []modelregistry.Model, error) {
	adresses, models, err := g.modelRegistry.ModelGetAll(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, nil, err
	}

	return adresses, models, nil
}
