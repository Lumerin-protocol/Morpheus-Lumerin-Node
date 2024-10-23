package registries

import (
	"context"
	"fmt"
	"math/big"

	i "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/interfaces"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/contracts/bindings/marketplace"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type Marketplace struct {
	// config
	marketplaceAddr common.Address

	// deps
	marketplace *marketplace.Marketplace
	client      i.ContractBackend
	log         lib.ILogger
}

func NewMarketplace(marketplaceAddr common.Address, client i.ContractBackend, log lib.ILogger) *Marketplace {
	mp, err := marketplace.NewMarketplace(marketplaceAddr, client)
	if err != nil {
		panic("invalid marketplace ABI")
	}
	return &Marketplace{
		marketplace:     mp,
		marketplaceAddr: marketplaceAddr,
		client:          client,
		log:             log,
	}
}

func (g *Marketplace) PostModelBid(opts *bind.TransactOpts, model common.Hash, pricePerSecond *big.Int) (common.Hash, error) {
	tx, err := g.marketplace.PostModelBid(opts, model, pricePerSecond)
	if err != nil {
		return common.Hash{}, lib.TryConvertGethError(err)
	}

	// Wait for the transaction receipt
	receipt, err := bind.WaitMined(opts.Context, g.client, tx)
	if err != nil {
		return common.Hash{}, lib.TryConvertGethError(err)
	}

	if receipt.Status != 1 {
		return receipt.TxHash, fmt.Errorf("Transaction failed")
	}

	return receipt.TxHash, nil
}

func (g *Marketplace) DeleteBid(opts *bind.TransactOpts, bidID common.Hash) (common.Hash, error) {
	tx, err := g.marketplace.DeleteModelBid(opts, bidID)
	if err != nil {
		return common.Hash{}, lib.TryConvertGethError(err)
	}

	// Wait for the transaction receipt
	receipt, err := bind.WaitMined(opts.Context, g.client, tx)
	if err != nil {
		return common.Hash{}, lib.TryConvertGethError(err)
	}

	if receipt.Status != 1 {
		return receipt.TxHash, fmt.Errorf("Transaction failed")
	}

	return receipt.TxHash, nil
}

func (g *Marketplace) GetBidById(ctx context.Context, bidID common.Hash) (*marketplace.IBidStorageBid, error) {
	bid, err := g.marketplace.GetBid(&bind.CallOpts{Context: ctx}, bidID)
	if err != nil {
		return nil, err
	}
	return &bid, nil
}

func (g *Marketplace) GetBestBidByModelId(ctx context.Context, modelID common.Hash) (common.Hash, *marketplace.IBidStorageBid, error) {
	limit := big.NewInt(100)
	offset := big.NewInt(0)

	bidIDs, err := g.marketplace.GetModelActiveBids(&bind.CallOpts{Context: ctx}, modelID, offset, limit)
	if err != nil {
		return common.Hash{}, nil, err
	}

	var cheapestBidID common.Hash
	var cheapestBid *marketplace.IBidStorageBid

	for _, bidID := range bidIDs {
		bid, err := g.marketplace.GetBid(&bind.CallOpts{Context: ctx}, bidID)
		if err != nil {
			return common.Hash{}, nil, err
		}
		if cheapestBid == nil {
			cheapestBid = &bid
			cheapestBidID = bidID

		} else if bid.PricePerSecond.Cmp(cheapestBid.PricePerSecond) < 0 {
			cheapestBid = &bid
			cheapestBidID = bidID
		}
	}

	return cheapestBidID, cheapestBid, nil
}

func (g *Marketplace) GetBidsByProvider(ctx context.Context, provider common.Address, offset *big.Int, limit uint8) ([][32]byte, []marketplace.IBidStorageBid, error) {
	bidIDs, err := g.marketplace.GetProviderBids(&bind.CallOpts{Context: ctx}, provider, offset, big.NewInt(int64(limit)))
	if err != nil {
		return nil, nil, err
	}

	return g.GetMultipleBids(ctx, bidIDs)
}

func (g *Marketplace) GetBidsByModelAgent(ctx context.Context, modelAgentId common.Hash, offset *big.Int, limit uint8) ([][32]byte, []marketplace.IBidStorageBid, error) {
	bidIDs, err := g.marketplace.GetModelBids(&bind.CallOpts{Context: ctx}, modelAgentId, offset, big.NewInt(int64(limit)))
	if err != nil {
		return nil, nil, err
	}
	return g.GetMultipleBids(ctx, bidIDs)
}

func (g *Marketplace) GetActiveBidsByProvider(ctx context.Context, provider common.Address, offset *big.Int, limit uint8) ([][32]byte, []marketplace.IBidStorageBid, error) {
	bidIDs, err := g.marketplace.GetProviderActiveBids(&bind.CallOpts{Context: ctx}, provider, offset, big.NewInt(int64(limit)))
	if err != nil {
		return nil, nil, err
	}
	return g.GetMultipleBids(ctx, bidIDs)
}

func (g *Marketplace) GetActiveBidsByModel(ctx context.Context, modelAgentId common.Hash, offset *big.Int, limit uint8) ([][32]byte, []marketplace.IBidStorageBid, error) {
	bidIDs, err := g.marketplace.GetModelActiveBids(&bind.CallOpts{Context: ctx}, modelAgentId, offset, big.NewInt(int64(limit)))
	if err != nil {
		return nil, nil, err
	}
	return g.GetMultipleBids(ctx, bidIDs)
}

func (g *Marketplace) GetMultipleBids(ctx context.Context, IDs [][32]byte) ([][32]byte, []marketplace.IBidStorageBid, error) {
	// todo: replace with multicall
	bids := make([]marketplace.IBidStorageBid, len(IDs))
	for i, ID := range IDs {
		bid, err := g.marketplace.GetBid(&bind.CallOpts{Context: ctx}, ID)
		if err != nil {
			return nil, nil, err
		}
		bids[i] = bid
	}
	return IDs, bids, nil
}
