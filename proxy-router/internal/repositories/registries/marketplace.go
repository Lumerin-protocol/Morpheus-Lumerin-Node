package registries

import (
	"context"
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

func (g *Marketplace) GetBidsByProvider(ctx context.Context, provider common.Address, offset *big.Int, limit uint8) ([][32]byte, []marketplace.Bid, error) {
	return g.marketplace.GetBidsByProvider(&bind.CallOpts{Context: ctx}, provider, offset, limit)
}

func (g *Marketplace) GetBidsByModelAgent(ctx context.Context, modelAgentId common.Hash, offset *big.Int, limit uint8) ([][32]byte, []marketplace.Bid, error) {
	return g.marketplace.GetBidsByModelAgent(&bind.CallOpts{Context: ctx}, modelAgentId, offset, limit)
}
