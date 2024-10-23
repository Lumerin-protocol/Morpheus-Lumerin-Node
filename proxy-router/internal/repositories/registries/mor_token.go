package registries

import (
	"context"
	"math/big"

	i "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/interfaces"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/contracts/bindings/morpheustoken"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
)

type MorToken struct {
	// config
	morTokenAddr common.Address

	// state
	nonce  uint64
	morABI *abi.ABI

	// deps
	mor    *morpheustoken.MorpheusToken
	client i.ContractBackend
	log    lib.ILogger
}

func NewMorToken(morTokenAddr common.Address, client i.ContractBackend, log lib.ILogger) *MorToken {
	mor, err := morpheustoken.NewMorpheusToken(morTokenAddr, client)
	if err != nil {
		panic("invalid mor ABI")
	}
	return &MorToken{
		mor:          mor,
		morTokenAddr: morTokenAddr,
		client:       client,
		log:          log,
	}
}

func (g *MorToken) GetBalance(ctx context.Context, account common.Address) (*big.Int, error) {
	return g.mor.BalanceOf(&bind.CallOpts{Context: ctx}, account)
}

func (g *MorToken) GetAllowance(ctx context.Context, owner common.Address, spender common.Address) (*big.Int, error) {
	return g.mor.Allowance(&bind.CallOpts{Context: ctx}, owner, spender)
}

func (g *MorToken) Approve(ctx *bind.TransactOpts, spender common.Address, amount *big.Int) (*types.Transaction, error) {
	tx, err := g.mor.Approve(ctx, spender, amount)
	if err != nil {
		return nil, lib.TryConvertGethError(err)
	}
	// Wait for the transaction receipt
	_, err = bind.WaitMined(ctx.Context, g.client, tx)
	if err != nil {
		return nil, err
	}
	return tx, nil
}

func (g *MorToken) GetTotalSupply(ctx context.Context) (*big.Int, error) {
	supply, err := g.mor.TotalSupply(&bind.CallOpts{Context: ctx})
	if err != nil {
		return nil, lib.TryConvertGethError(err)
	}
	return supply, nil
}

func (g *MorToken) Transfer(ctx *bind.TransactOpts, to common.Address, value *big.Int) (*types.Transaction, error) {
	tx, err := g.mor.Transfer(ctx, to, value)
	if err != nil {
		return nil, lib.TryConvertGethError(err)
	}

	// Wait for the transaction receipt
	_, err = bind.WaitMined(ctx.Context, g.client, tx)
	if err != nil {
		return nil, err
	}
	return tx, nil
}
