package registries

import (
	"context"
	"fmt"

	i "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/interfaces"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/contracts/bindings/modelregistry"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type ModelRegistry struct {
	// config
	modelRegistryAddr common.Address

	// state
	nonce uint64

	// deps
	modelRegistry *modelregistry.ModelRegistry
	client        i.ContractBackend
	log           lib.ILogger
}

func NewModelRegistry(modelRegistryAddr common.Address, client i.ContractBackend, log lib.ILogger) *ModelRegistry {
	mr, err := modelregistry.NewModelRegistry(modelRegistryAddr, client)
	if err != nil {
		panic("invalid model registry ABI")
	}
	return &ModelRegistry{
		modelRegistry:     mr,
		modelRegistryAddr: modelRegistryAddr,
		client:            client,
		log:               log,
	}
}

func (g *ModelRegistry) GetAllModels(ctx context.Context) ([][32]byte, []modelregistry.IModelStorageModel, error) {
	// adresses, models, err := g.modelRegistry.Models(&bind.CallOpts{Context: ctx},)
	// if err != nil {
	// 	return nil, nil, err
	// }

	return nil, nil, fmt.Errorf("Not implemented")
}

func (g *ModelRegistry) CreateNewModel(opts *bind.TransactOpts, modelId common.Hash, ipfsID common.Hash, fee *lib.BigInt, stake *lib.BigInt, name string, tags []string) error {
	tx, err := g.modelRegistry.ModelRegister(opts, modelId, ipfsID, &fee.Int, &stake.Int, name, tags)
	if err != nil {
		return lib.TryConvertGethError(err)
	}

	// Wait for the transaction receipt
	receipt, err := bind.WaitMined(opts.Context, g.client, tx)
	if err != nil {
		return lib.TryConvertGethError(err)
	}

	// Find the event log
	for _, log := range receipt.Logs {
		_, err := g.modelRegistry.ParseModelRegisteredUpdated(*log)

		if err != nil {
			continue // not our event, skip it
		}

		return nil
	}

	return fmt.Errorf("ModelRegistered event not found in transaction logs")
}

func (g *ModelRegistry) DeregisterModel(opts *bind.TransactOpts, modelId common.Hash) (common.Hash, error) {
	tx, err := g.modelRegistry.ModelDeregister(opts, modelId)

	if err != nil {
		return common.Hash{}, lib.TryConvertGethError(err)
	}

	// Wait for the transaction receipt
	receipt, err := bind.WaitMined(opts.Context, g.client, tx)
	if err != nil {
		return common.Hash{}, lib.TryConvertGethError(err)
	}

	// Find the event log
	for _, log := range receipt.Logs {
		_, err := g.modelRegistry.ParseModelDeregistered(*log)

		if err != nil {
			continue // not our event, skip it
		}

		return tx.Hash(), nil
	}

	return common.Hash{}, fmt.Errorf("ModelDeregistered event not found in transaction logs")
}

func (g *ModelRegistry) GetModelById(ctx context.Context, modelId common.Hash) (*modelregistry.IModelStorageModel, error) {
	model, err := g.modelRegistry.GetModel(&bind.CallOpts{Context: ctx}, modelId)
	if err != nil {
		return nil, err
	}
	return &model, nil
}
