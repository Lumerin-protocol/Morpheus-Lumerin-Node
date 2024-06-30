package structs

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Bid struct {
	Id             common.Hash
	Provider       common.Address
	ModelAgentId   common.Hash
	PricePerSecond *big.Int
	Nonce          *big.Int
	CreatedAt      *big.Int
	DeletedAt      *big.Int
}
