package structs

import (
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/ethereum/go-ethereum/common"
)

type OpenSessionRequest struct {
	Approval    lib.HexString `json:"approval" binding:"required" validate:"hexadecimal"`
	ApprovalSig lib.HexString `json:"approvalSig" binding:"required" validate:"hexadecimal"`
	Stake       *lib.BigInt   `json:"stake" binding:"required" validate:"number,gt=0"`
}

type SendRequest struct {
	To     common.Address `json:"to" binding:"required" validate:"eth_addr"`
	Amount *lib.BigInt    `json:"amount" binding:"required" validate:"number,gt=0"`
}

type PathHex32ID struct {
	ID lib.Hash `uri:"id" binding:"required" validate:"hex32"`
}

type PathEthAddrID struct {
	ID common.Address `uri:"id" binding:"required" validate:"eth_addr"`
}

type QueryOffsetLimit struct {
	Offset lib.BigInt `form:"offset,default=0" binding:"omitempty" validate:"number"`
	Limit  uint8      `form:"limit,default=10" binding:"omitempty" validate:"number"`
}

type QueryPageLimit struct {
	Page  uint64 `form:"page,default=0"   binding:"omitempty" validate:"number"`
	Limit uint8  `form:"limit,default=10" binding:"omitempty" validate:"number"`
}

type QuerySpender struct {
	Spender common.Address `form:"spender" binding:"required" validate:"eth_addr"`
}

type QueryApprove struct {
	*QuerySpender
	Amount *lib.BigInt `form:"amount" binding:"required" validate:"number,gt=0"`
}

type QueryUserOrProvider struct {
	User     lib.Address `form:"user" binding:"omitempty" validate:"eth_addr"`
	Provider lib.Address `form:"provider" binding:"omitempty" validate:"eth_addr"`
}

type OpenSessionWithDurationRequest struct {
	SessionDuration *lib.BigInt `json:"sessionDuration"`
}
