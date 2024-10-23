package blockchainapi

import (
	"math"
	"math/big"

	m "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/contracts/bindings/marketplace"
	s "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/contracts/bindings/sessionrouter"
	"github.com/ethereum/go-ethereum/common"
)

type ModelStats struct {
	TpsScaled1000 s.LibSDSD
	TtftMs        s.LibSDSD
	TotalDuration s.LibSDSD
	Count         int
}

func sampleDataTPS() ([][32]byte, []m.IBidStorageBid, []s.IStatsStorageProviderModelStats, ModelStats) {
	modelID := common.HexToHash("0x01")
	bidIds := [][32]byte{
		{0x01},
		{0x02},
		{0x03},
	}

	bids := []m.IBidStorageBid{
		{
			PricePerSecond: ToDecimal(10, DecimalsMOR),
			Provider:       common.HexToAddress("0x01"),
			ModelId:        modelID,
			Nonce:          common.Big0,
			CreatedAt:      common.Big1,
			DeletedAt:      common.Big0,
		},
		{
			PricePerSecond: ToDecimal(10, DecimalsMOR),
			Provider:       common.HexToAddress("0x02"),
			ModelId:        modelID,
			Nonce:          common.Big0,
			CreatedAt:      common.Big1,
			DeletedAt:      common.Big0,
		},
		{
			PricePerSecond: ToDecimal(10, DecimalsMOR),
			Provider:       common.HexToAddress("0x03"),
			ModelId:        modelID,
			Nonce:          common.Big0,
			CreatedAt:      common.Big1,
			DeletedAt:      common.Big0,
		},
	}
	pmStats := []s.IStatsStorageProviderModelStats{
		{
			TpsScaled1000: s.LibSDSD{Mean: 10, SqSum: 100},
			TtftMs:        s.LibSDSD{Mean: 20, SqSum: 200},
			TotalDuration: 30,
			SuccessCount:  6,
			TotalCount:    10,
		},
		{
			TpsScaled1000: s.LibSDSD{Mean: 20, SqSum: 100},
			TtftMs:        s.LibSDSD{Mean: 20, SqSum: 200},
			TotalDuration: 30,
			SuccessCount:  6,
			TotalCount:    10,
		},
		{
			TpsScaled1000: s.LibSDSD{Mean: 30, SqSum: 100},
			TtftMs:        s.LibSDSD{Mean: 20, SqSum: 200},
			TotalDuration: 30,
			SuccessCount:  6,
			TotalCount:    10,
		},
	}
	mStats := ModelStats{
		TpsScaled1000: s.LibSDSD{Mean: 20, SqSum: 100},
		TtftMs:        s.LibSDSD{Mean: 20, SqSum: 200},
		TotalDuration: s.LibSDSD{Mean: 30, SqSum: 300},
		Count:         3,
	}

	return bidIds, bids, pmStats, mStats
}

func FromDecimal(value *big.Int, decimals int) float64 {
	return float64(value.Int64()) / float64(math.Pow10(decimals))
}

func ToDecimal(value float64, decimals int) *big.Int {
	a, _ := big.NewFloat(0).Mul(big.NewFloat(math.Pow10(decimals)), big.NewFloat(value)).Int(nil)
	return a
}

const DecimalsMOR = 18
