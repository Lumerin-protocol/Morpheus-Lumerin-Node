package structs

import (
	"math/big"

	"github.com/ethereum/go-ethereum/common"
)

type Session struct {
	Id                      string
	User                    common.Address
	Provider                common.Address
	ModelAgentId            string
	BidID                   string
	Stake                   *big.Int
	PricePerSecond          *big.Int
	CloseoutReceipt         string
	CloseoutType            *big.Int
	ProviderWithdrawnAmount *big.Int
	OpenedAt                *big.Int
	EndsAt                  *big.Int
	ClosedAt                *big.Int
}
