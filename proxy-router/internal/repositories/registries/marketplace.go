package registries

import (
	"context"
	"fmt"
	"math/big"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/contracts/marketplace"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/ethclient"
)

type Marketplace struct {
	// config
	marketplaceAddr common.Address

	// deps
	marketplace *marketplace.Marketplace
	client      *ethclient.Client
	log         lib.ILogger
}

func NewMarketplace(marketplaceAddr common.Address, client *ethclient.Client, log lib.ILogger) *Marketplace {
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

func (g *Marketplace) PostModelBid(ctx *bind.TransactOpts, provider string, model [32]byte, pricePerSecond *big.Int) error {
	fmt.Println("PostModelBid", provider, model, pricePerSecond)
	tx, err := g.marketplace.PostModelBid(ctx, common.HexToAddress(provider), model, pricePerSecond)
	if err != nil {
		return lib.TryConvertGethError(err, marketplace.MarketplaceMetaData)
	}

	// Wait for the transaction receipt
	receipt, err := bind.WaitMined(context.Background(), g.client, tx)

	if err != nil {
		return lib.TryConvertGethError(err, marketplace.MarketplaceMetaData)
	}

	// Find the event log
	for _, log := range receipt.Logs {
		// Check if the log belongs to the OpenSession event
		_, err := g.marketplace.ParseBidPosted(*log)

		if err != nil {
			continue // not our event, skip it
		}

		return nil
	}

	return fmt.Errorf("PostModelBid event not found in transaction logs")
}

func (g *Marketplace) GetBidById(ctx context.Context, bidId [32]byte) (*marketplace.Bid, error) {
	bid, err := g.marketplace.BidMap(&bind.CallOpts{Context: ctx}, bidId)
	if err != nil {
		return nil, err
	}
	return &bid, nil
}

func (g *Marketplace) GetBestBidByModelId(ctx context.Context, modelId common.Hash) (common.Hash, *marketplace.Bid, error) {
	limit := uint8(100)
	offset := big.NewInt(0)

	ids, bids, err := g.marketplace.GetBidsByModelAgent(&bind.CallOpts{Context: ctx}, modelId, offset, limit)
	if err != nil {
		return common.Hash{}, nil, err
	}

	// TODO: replace with a rating system
	cheapestBid := bids[0]
	idIndex := 0
	for i, bid := range bids {
		if bid.PricePerSecond.Cmp(cheapestBid.PricePerSecond) < 0 {
			cheapestBid = bid
			idIndex = i
		}
	}

	return ids[idIndex], &cheapestBid, nil
}

func (g *Marketplace) GetAllBidsWithRating(ctx context.Context, modelAgentID [32]byte) ([][32]byte, []marketplace.Bid, []marketplace.ProviderModelStats, error) {
	batchSize := uint8(255)
	return collectBids(ctx, modelAgentID, g.GetBidsWithRating, batchSize)
}

func (g *Marketplace) GetBidsWithRating(ctx context.Context, modelAgentID [32]byte, offset *big.Int, limit uint8) ([][32]byte, []marketplace.Bid, []marketplace.ProviderModelStats, error) {
	return g.marketplace.GetActiveBidsRatingByModelAgent(&bind.CallOpts{Context: ctx}, modelAgentID, offset, limit)
}

func (g *Marketplace) GetModelStats(ctx context.Context, modelID [32]byte) (marketplace.ModelStats, error) {
	return g.marketplace.GetModelStats(&bind.CallOpts{Context: ctx}, modelID)
}

func (g *Marketplace) GetBidsByProvider(ctx context.Context, provider common.Address, offset *big.Int, limit uint8) ([][32]byte, []marketplace.Bid, error) {
	return g.marketplace.GetBidsByProvider(&bind.CallOpts{Context: ctx}, provider, offset, limit)
}

func (g *Marketplace) GetBidsByModelAgent(ctx context.Context, modelAgentId common.Hash, offset *big.Int, limit uint8) ([][32]byte, []marketplace.Bid, error) {
	return g.marketplace.GetBidsByModelAgent(&bind.CallOpts{Context: ctx}, modelAgentId, offset, limit)
}

func (g *Marketplace) GetActiveBidsByProvider(ctx context.Context, provider common.Address) ([][32]byte, []marketplace.Bid, error) {
	return g.marketplace.GetActiveBidsByProvider(&bind.CallOpts{Context: ctx}, provider)
}

func (g *Marketplace) GetActiveBidsByModel(ctx context.Context, modelAgentId common.Hash) ([][32]byte, []marketplace.Bid, error) {
	return g.marketplace.GetActiveBidsByModelAgent(&bind.CallOpts{Context: ctx}, modelAgentId)
}

type BidsGetter = func(ctx context.Context, modelAgentID [32]byte, offset *big.Int, limit uint8) ([][32]byte, []marketplace.Bid, []marketplace.ProviderModelStats, error)

func collectBids(ctx context.Context, modelAgentID [32]byte, bidsGetter BidsGetter, batchSize uint8) ([][32]byte, []marketplace.Bid, []marketplace.ProviderModelStats, error) {
	offset := big.NewInt(0)
	bids := make([]marketplace.Bid, 0)
	ids := make([][32]byte, 0)
	providerModelStats := make([]marketplace.ProviderModelStats, 0)

	for {
		if ctx.Err() != nil {
			return nil, nil, nil, ctx.Err()
		}

		idsBatch, bidsBatch, providerModelStatsBatch, err := bidsGetter(ctx, modelAgentID, offset, batchSize)
		if err != nil {
			return nil, nil, nil, err
		}

		ids = append(ids, idsBatch...)
		bids = append(bids, bidsBatch...)
		providerModelStats = append(providerModelStats, providerModelStatsBatch...)

		if len(bidsBatch) < int(batchSize) {
			break
		}

		offset.Add(offset, big.NewInt(int64(batchSize)))
	}

	return ids, bids, providerModelStats, nil
}
