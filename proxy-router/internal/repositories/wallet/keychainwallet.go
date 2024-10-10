package wallet

import (
	"errors"
	"sync"

	i "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/interfaces"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/keychain"
	hdwallet "github.com/miguelmota/go-ethereum-hdwallet"
)

const (
	PRIVATE_KEY_KEY     = "private-key"
	MNEMONIC_KEY        = "mnemonic"
	DERIVATION_PATH_KEY = "mnemonic-derivation-path"
)

var (
	ErrWalletNotSet = errors.New("wallet not set")
	ErrWallet       = errors.New("cannot retrieve mnemonic or private key")
)

type KeychainWallet struct {
	storage   i.KeyValueStorage
	updatedCh chan struct{}
	mutex     sync.Mutex
}

func NewKeychainWallet(keychain i.KeyValueStorage) *KeychainWallet {
	return &KeychainWallet{
		storage:   keychain,
		updatedCh: make(chan struct{}),
	}
}

// GetPrivateKey use this function to get the private key regardless of whether it was stored as a mnemonic or private key
//
// errors with ErrPkeyAndMnemonic if both mnemonic and private key are stored
func (w *KeychainWallet) GetPrivateKey() (lib.HexString, error) {
	prKey, prKeyErr := w.getStoredPrivateKey()
	mnem, derivation, mnemErr := w.getStoredMnemonic()

	if errors.Is(prKeyErr, keychain.ErrKeyNotFound) && errors.Is(mnemErr, keychain.ErrKeyNotFound) {
		return nil, ErrWalletNotSet
	}

	if prKey != nil && mnem != "" {
		return nil, errors.New("both mnemonic and private key are stored")
	}

	if prKey != nil {
		return prKey, nil
	}

	if mnem != "" && derivation != "" {
		return w.mnemonicToPrivateKey(mnem, derivation)
	}

	var err = ErrWallet

	if mnemErr != nil && !errors.Is(mnemErr, keychain.ErrKeyNotFound) {
		err = lib.WrapError(err, mnemErr)
	}
	if prKeyErr != nil && !errors.Is(mnemErr, keychain.ErrKeyNotFound) {
		err = lib.WrapError(err, prKeyErr)
	}

	return nil, err
}

// SetPrivateKey stores the private key of the wallet
func (w *KeychainWallet) SetPrivateKey(privateKey lib.HexString) error {
	err := w.storage.Upsert(PRIVATE_KEY_KEY, privateKey.Hex())
	if err != nil {
		return err
	}
	// either mnemonic or private key can be stored at a time
	err = w.storage.DeleteIfExists(MNEMONIC_KEY)
	if err == nil {
		return err
	}

	err = w.storage.DeleteIfExists(DERIVATION_PATH_KEY)
	if err != nil {
		return err
	}

	// notify the listeners that the private key has been updated
	w.notifyUpdated()

	return nil
}

// SetMnemonic stores the mnemonic of the wallet
func (w *KeychainWallet) SetMnemonic(mnemonic string, derivationPath string) error {
	err := w.storage.Upsert(MNEMONIC_KEY, mnemonic)
	if err != nil {
		return err
	}
	err = w.storage.Upsert(DERIVATION_PATH_KEY, derivationPath)
	if err != nil {
		return err
	}

	// either mnemonic or private key can be stored at a time
	err = w.storage.DeleteIfExists(PRIVATE_KEY_KEY)
	if err == nil {
		return err
	}

	w.notifyUpdated()

	return nil
}

func (w *KeychainWallet) DeleteWallet() error {
	err := w.storage.DeleteIfExists(PRIVATE_KEY_KEY)
	if err != nil {
		return err
	}

	err = w.storage.DeleteIfExists(MNEMONIC_KEY)
	if err != nil {
		return err
	}

	err = w.storage.DeleteIfExists(DERIVATION_PATH_KEY)
	if err != nil {
		return err
	}

	w.notifyUpdated()

	return nil
}

// getStoredPrivateKey retrieves the private key of the wallet
func (w *KeychainWallet) getStoredPrivateKey() (lib.HexString, error) {
	prKey, err := w.storage.Get(PRIVATE_KEY_KEY)
	if err != nil {
		return nil, err
	}
	return lib.StringToHexString(prKey)
}

// getStoredMnemonic retrieves the mnemonic of the wallet
func (w *KeychainWallet) getStoredMnemonic() (string, string, error) {
	mnemonic, err := w.storage.Get(MNEMONIC_KEY)
	if err != nil {
		return "", "", err
	}

	derivationPath, err := w.storage.Get(DERIVATION_PATH_KEY)
	if err != nil {
		return "", "", err
	}

	return mnemonic, derivationPath, nil
}

func (w *KeychainWallet) mnemonicToPrivateKey(mnemonic, derivationPath string) (lib.HexString, error) {
	wallet, err := hdwallet.NewFromMnemonic(mnemonic)
	if err != nil {
		return nil, err
	}
	path, err := hdwallet.ParseDerivationPath(derivationPath)
	if err != nil {
		return nil, err
	}
	account, err := wallet.Derive(path, true)
	if err != nil {
		return nil, err
	}
	return wallet.PrivateKeyBytes(account)
}

func (w *KeychainWallet) PrivateKeyUpdated() <-chan struct{} {
	w.mutex.Lock()
	defer w.mutex.Unlock()

	return w.updatedCh
}

func (w *KeychainWallet) notifyUpdated() {
	close(w.updatedCh)
	w.mutex.Lock()
	defer w.mutex.Unlock()
	w.updatedCh = make(chan struct{})
}
