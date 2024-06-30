package structs

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Provider struct {
	Address   common.Address
	Endpoint  string
	Stake     *big.Int
	CreatedAt *big.Int
	IsDeleted bool
}
