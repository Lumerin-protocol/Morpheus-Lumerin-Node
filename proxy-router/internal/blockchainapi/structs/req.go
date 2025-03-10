package structs

import (
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/ethereum/go-ethereum/common"
)

type OpenSessionRequest struct {
	Approval      lib.HexString `json:"approval" binding:"required" validate:"hexadecimal" format:"hex" example:"0x1234"`
	ApprovalSig   lib.HexString `json:"approvalSig" binding:"required" validate:"hexadecimal" format:"hex" example:"0x1234"`
	Stake         *lib.BigInt   `json:"stake" binding:"required" validate:"number,gt=0" example:"123000000000"`
	DirectPayment bool          `json:"directPayment" binding:"omitempty"`
}

type AmountReq struct {
	Amount *lib.BigInt `json:"amount" binding:"required" validate:"number,gt=0"`
}

type SendRequest struct {
	To     common.Address `json:"to" binding:"required" validate:"eth_addr"`
	Amount *lib.BigInt    `json:"amount" binding:"required" validate:"number,gt=0"`
}

type PathHex32ID struct {
	ID lib.Hash `uri:"id" binding:"required" validate:"hex32" format:"hex" example:"0x1234"`
}

type PathEthAddrID struct {
	ID lib.Address `uri:"id" binding:"required" validate:"eth_addr"`
}

type QueryOffsetLimitOrder struct {
	Offset lib.BigInt `form:"offset,default=0" binding:"omitempty" validate:"number,gte=0" example:"0" swaggertype:"integer"`
	Limit  uint8       `form:"limit,default=10" binding:"omitempty" validate:"number,gte=1" example:"10"`
	Order  string      `form:"order,default=asc" binding:"omitempty" validate:"oneof=asc desc" example:"asc"`
}

type QueryOffsetLimitOrderNoDefault struct {
	Offset lib.BigInt `form:"offset,default=0" binding:"omitempty" validate:"number,gte=0" example:"0" swaggertype:"integer"`
	Limit  uint8       `form:"limit,default=0" binding:"omitempty" validate:"number,gte=1" example:"10"`
	Order  string      `form:"order,default=asc" binding:"omitempty" validate:"oneof=asc desc" example:"asc"`
}

type QueryPageLimit struct {
	Page  uint64 `form:"page,default=0"   binding:"omitempty" validate:"number"`
	Limit uint8  `form:"limit,default=10" binding:"omitempty" validate:"number"`
}

type QuerySpender struct {
	Spender lib.Address `form:"spender" binding:"required" validate:"eth_addr"`
}

type QueryApprove struct {
	*QuerySpender
	Amount *lib.BigInt `form:"amount" binding:"required" validate:"number,gt=0"`
}

type QueryUser struct {
	User lib.Address `form:"user" binding:"required" validate:"eth_addr"`
}

type QueryProvider struct {
	Provider lib.Address `form:"provider" binding:"required" validate:"eth_addr"`
}

type OpenSessionWithDurationRequest struct {
	SessionDuration *lib.BigInt `json:"sessionDuration"`
}

type OpenSessionWithFailover struct {
	SessionDuration *lib.BigInt `json:"sessionDuration"`
	DirectPayment   bool        `json:"directPayment" binding:"omitempty"`
	Failover        bool        `json:"failover" binding:"omitempty"`
}

type CreateBidRequest struct {
	ModelID        string      `json:"modelID" binding:"required" validate:"hex32"`
	PricePerSecond *lib.BigInt `json:"pricePerSecond" binding:"required" validate:"number,gt=0"`
}

type CreateProviderRequest struct {
	Stake    *lib.BigInt `json:"stake" binding:"required" validate:"number" example:"123000000000"`
	Endpoint string      `json:"endpoint" binding:"required" validate:"string" example:"mycoolmornode.domain.com:3989"`
}

type CreateModelRequest struct {
	ID     string      `json:"id" binding:"omitempty" validate:"hex32" example:"0x1234"`
	IpfsID string      `json:"ipfsID" binding:"required" validate:"hex32" example:"0x1234"`
	Fee    *lib.BigInt `json:"fee" binding:"required" validate:"number" example:"123000000000"`
	Stake  *lib.BigInt `json:"stake" binding:"required" validate:"number" example:"123000000000"`
	Name   string      `json:"name" binding:"required" validate:"min=1,max=64" example:"Llama 2.0"`
	Tags   []string    `json:"tags" binding:"required" validate:"min=1,max=64,dive,min=1,max=64"`
}
