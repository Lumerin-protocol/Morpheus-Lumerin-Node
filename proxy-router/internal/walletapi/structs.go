package walletapi

import "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"

type SetupWalletReqBody struct {
	PrivateKey lib.HexString `json:"privateKey" binding:"required" validate:"required,eth_addr"`
}
