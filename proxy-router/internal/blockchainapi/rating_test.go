package blockchainapi

import (
	"math"
	"testing"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	s "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/contracts/bindings/sessionrouter"
	"github.com/stretchr/testify/require"
)

func TestRating(t *testing.T) {
	bidIds, bids, pmStats, mStats := sampleDataTPS()

	scoredBids := rateBids(bidIds, bids, pmStats, mStats, lib.NewTestLogger())

	for i := 1; i < len(scoredBids); i++ {
		require.GreaterOrEqual(t, scoredBids[i-1].Score, scoredBids[i].Score, "scoredBids not sorted")
	}
}

func TestGetScoreZeroObservations(t *testing.T) {
	_, bids, _, _ := sampleDataTPS()
	score := getScore(bids[0], s.IStatsStorageProviderModelStats{}, ModelStats{})
	if math.IsNaN(score) {
		require.Fail(t, "score is NaN")
	}
}

func TestGetScoreSingleObservation(t *testing.T) {
	_, bids, _, _ := sampleDataTPS()
	pmStats := s.IStatsStorageProviderModelStats{
		TpsScaled1000: s.LibSDSD{Mean: 1000, SqSum: 0},
		TtftMs:        s.LibSDSD{Mean: 1000, SqSum: 0},
		TotalDuration: 1000,
		SuccessCount:  1,
		TotalCount:    1,
	}
	mStats := ModelStats{
		TpsScaled1000: s.LibSDSD{Mean: 1000, SqSum: 0},
		TtftMs:        s.LibSDSD{Mean: 1000, SqSum: 0},
		TotalDuration: s.LibSDSD{Mean: 1000, SqSum: 0},
		Count:         1,
	}
	score := getScore(bids[0], pmStats, mStats)
	if math.IsNaN(score) {
		require.Fail(t, "score is NaN")
	}
}
