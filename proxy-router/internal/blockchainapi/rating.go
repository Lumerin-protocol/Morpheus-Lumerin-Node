package blockchainapi

import (
	"math"
	"sort"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/blockchainapi/structs"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	m "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/contracts/bindings/marketplace"
	s "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/contracts/bindings/sessionrouter"
)

func rateBids(bidIds [][32]byte, bids []m.IBidStorageBid, pmStats []s.IStatsStorageProviderModelStats, mStats ModelStats, log lib.ILogger) []structs.ScoredBid {
	scoredBids := make([]structs.ScoredBid, len(bids))

	for i := range bids {
		score := getScore(bids[i], pmStats[i], mStats)
		if math.IsNaN(score) || math.IsInf(score, 0) {
			log.Errorf("provider score is not valid %d for bid %v, pmStats %v, mStats %v", score, bidIds[i], pmStats[i], mStats)
			score = 0
		}
		scoredBid := structs.ScoredBid{
			Bid: structs.Bid{
				Id:             bidIds[i],
				Provider:       bids[i].Provider,
				ModelAgentId:   bids[i].ModelId,
				PricePerSecond: &lib.BigInt{Int: *(bids[i].PricePerSecond)},
				Nonce:          &lib.BigInt{Int: *(bids[i].Nonce)},
				CreatedAt:      &lib.BigInt{Int: *(bids[i].CreatedAt)},
				DeletedAt:      &lib.BigInt{Int: *(bids[i].DeletedAt)},
			},
			Score: score,
		}
		scoredBids[i] = scoredBid
	}

	sort.Slice(scoredBids, func(i, j int) bool {
		return scoredBids[i].Score > scoredBids[j].Score
	})

	return scoredBids
}

func getScore(bid m.IBidStorageBid, pmStats s.IStatsStorageProviderModelStats, mStats ModelStats) float64 {
	tpsWeight, ttftWeight, durationWeight, successWeight := 0.27, 0.11, 0.27, 0.35
	count := int64(mStats.Count)

	tpsScore := tpsWeight * normRange(normZIndex(pmStats.TpsScaled1000.Mean, mStats.TpsScaled1000, count), 3.0)
	// ttft impact is negative
	ttftScore := ttftWeight * normRange(-1*normZIndex(pmStats.TtftMs.Mean, mStats.TtftMs, count), 3.0)
	durationScore := durationWeight * normRange(normZIndex(int64(pmStats.TotalDuration), mStats.TotalDuration, count), 3.0)
	successScore := successWeight * math.Pow(ratioScore(pmStats.SuccessCount, pmStats.TotalCount), 2)

	priceFloatDecimal, _ := bid.PricePerSecond.Float64()
	priceFloat := priceFloatDecimal / math.Pow10(18)

	result := (tpsScore + ttftScore + durationScore + successScore) / priceFloat

	return result
}

func ratioScore(num, denom uint32) float64 {
	if denom == 0 {
		return 0
	}
	return float64(num) / float64(denom)
}

// normZIndex normalizes the value using z-index
func normZIndex(pmMean int64, mSD s.LibSDSD, obsNum int64) float64 {
	sd := getSD(mSD, obsNum)
	if sd == 0 {
		return 0
	}
	// TODO: consider variance(SD) of provider model stats
	return float64(pmMean-mSD.Mean) / getSD(mSD, obsNum)
}

// normRange normalizes the incoming data within the range [-normRange, normRange]
// to the range [0, 1] cutting off the values outside the range
func normRange(input float64, normRange float64) float64 {
	return cutRange01((input + normRange) / (2.0 * normRange))
}

func getSD(sd s.LibSDSD, obsNum int64) float64 {
	return math.Sqrt(getVariance(sd, obsNum))
}

func getVariance(sd s.LibSDSD, obsNum int64) float64 {
	if obsNum <= 1 {
		return 0
	}
	return float64(sd.SqSum) / float64(obsNum-1)
}

func cutRange01(val float64) float64 {
	if val > 1 {
		return 1
	}
	if val < 0 {
		return 0
	}
	return val
}
