package registries

import (
	"context"

	"github.com/Lumerin-protocol/Morpheus-Lumerin-Node/proxy-router/contracts/agentregistry"
	"github.com/Lumerin-protocol/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type AgentRegistry struct {
	// config
	agentRegistryAddr common.Address

	// state
	nonce uint64
	arABI *abi.ABI

	// deps
	agentRegistry *agentregistry.AgentRegistry
	client        *ethclient.Client
	log           lib.ILogger
}

func NewAgentRegistry(agentRegistryAddr common.Address, client *ethclient.Client, log lib.ILogger) *AgentRegistry {
	ar, err := agentregistry.NewAgentRegistry(agentRegistryAddr, client)
	if err != nil {
		panic("invalid agent registry ABI")
	}
	arABI, err := agentregistry.AgentRegistryMetaData.GetAbi()
	if err != nil {
		panic("invalid agent registry ABI: " + err.Error())
	}
	return &AgentRegistry{
		agentRegistry:     ar,
		agentRegistryAddr: agentRegistryAddr,
		client:            client,
		arABI:             arABI,
		log:               log,
	}
}

func (ar *AgentRegistry) GetAllAgents(ctx context.Context) ([]agentregistry.AgentRegistryAgent, error) {
	return ar.agentRegistry.GetAll(&bind.CallOpts{Context: ctx})
}
