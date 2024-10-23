package registries

import (
	"context"
	"math/big"
	"testing"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/contracts/bindings/sessionrouter"
)

func TestCollectBids(t *testing.T) {
	// bidIDs, bids, stats, err := collectBids(context.TODO(), [32]byte{0x01}, bidsGetter1000, 100)
	// require.NoError(t, err)
	// require.Len(t, bidIDs, 1000)
	// require.Len(t, bids, 1000)
	// require.Len(t, stats, 1000)
}

// bidsGetter1000 simulates a paginated query to get all bids for a model
func bidsGetter1000(ctx context.Context, modelId [32]byte, offset *big.Int, limit uint8) ([][32]byte, []sessionrouter.IBidStorageBid, []sessionrouter.IStatsStorageProviderModelStats, error) {
	maxItems := 1000
	ids := [][32]byte{}
	bids := []sessionrouter.IBidStorageBid{}
	stats := []sessionrouter.IStatsStorageProviderModelStats{}
	for i := offset.Int64(); i < offset.Int64()+int64(limit); i++ {
		if i >= int64(maxItems) {
			break
		}
		ids = append(ids, [32]byte{byte(i)})
		bids = append(bids, sessionrouter.IBidStorageBid{
			PricePerSecond: big.NewInt(i),
			Provider:       [20]byte{byte(i)},
			ModelId:        modelId,
			Nonce:          big.NewInt(i),
			CreatedAt:      big.NewInt(i),
			DeletedAt:      big.NewInt(i),
		})
		stats = append(stats, sessionrouter.IStatsStorageProviderModelStats{
			TpsScaled1000: sessionrouter.LibSDSD{},
			TtftMs:        sessionrouter.LibSDSD{},
			TotalDuration: uint32(i),
			SuccessCount:  uint32(i),
			TotalCount:    uint32(i),
		})
	}
	return ids, bids, stats, nil
}
