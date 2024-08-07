package registries

import (
	"context"
	"math/big"
	"testing"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/contracts/marketplace"
	"github.com/stretchr/testify/require"
)

func TestCollectBids(t *testing.T) {
	bidIDs, bids, stats, err := collectBids(context.TODO(), [32]byte{0x01}, bidsGetter1000, 100)
	require.NoError(t, err)
	require.Len(t, bidIDs, 1000)
	require.Len(t, bids, 1000)
	require.Len(t, stats, 1000)
}

// bidsGetter1000 simulates a paginated query to get all bids for a model
func bidsGetter1000(ctx context.Context, modelId [32]byte, offset *big.Int, limit uint8) ([][32]byte, []marketplace.Bid, []marketplace.ProviderModelStats, error) {
	maxItems := 1000
	ids := [][32]byte{}
	bids := []marketplace.Bid{}
	stats := []marketplace.ProviderModelStats{}
	for i := offset.Int64(); i < offset.Int64()+int64(limit); i++ {
		if i >= int64(maxItems) {
			break
		}
		ids = append(ids, [32]byte{byte(i)})
		bids = append(bids, marketplace.Bid{
			PricePerSecond: big.NewInt(i),
			Provider:       [20]byte{byte(i)},
			ModelAgentId:   modelId,
			Nonce:          big.NewInt(i),
			CreatedAt:      big.NewInt(i),
			DeletedAt:      big.NewInt(i),
		})
		stats = append(stats, marketplace.ProviderModelStats{
			TpsScaled1000: marketplace.LibSDSD{},
			TtftMs:        marketplace.LibSDSD{},
			TotalDuration: uint32(i),
			SuccessCount:  uint32(i),
			TotalCount:    uint32(i),
		})
	}
	return ids, bids, stats, nil
}
