package wallet

import (
	"errors"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
)

var ErrEnvWalletSet = errors.New("cannot set private key for env wallet, switch to keychain wallet by removing WALLET_PRIVATE_KEY env var")

type EnvWallet struct {
	privKey   lib.HexString
	updatedCh chan struct{}
}

func NewEnvWallet(privKey lib.HexString) *EnvWallet {
	return &EnvWallet{
		updatedCh: make(chan struct{}),
		privKey:   privKey,
	}
}

func (w *EnvWallet) GetPrivateKey() (lib.HexString, error) {
	return w.privKey, nil
}

func (w *EnvWallet) SetPrivateKey(privateKeyOxHex lib.HexString) error {
	return ErrEnvWalletSet
}

func (w *EnvWallet) SetMnemonic(mnemonic string, derivationPath string) error {
	return ErrEnvWalletSet
}

func (w *EnvWallet) PrivateKeyUpdated() <-chan struct{} {
	return w.updatedCh
}
