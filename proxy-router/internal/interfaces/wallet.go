package interfaces

import "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"

type Wallet interface {
	GetPrivateKey() (lib.HexString, error)
	SetPrivateKey(privateKey lib.HexString) error
	SetMnemonic(mnemonic string, derivationPath string) error
	PrivateKeyUpdated() <-chan struct{}
}
