package lib

import (
	"crypto/ecdsa"
	"crypto/rand"
	"encoding/hex"
	"fmt"

	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/crypto/ecies"
)

var (
	ErrInvalidPrivateKey = fmt.Errorf("invalid private key")
)

func DecryptBytes(bytes, privateKey []byte) ([]byte, error) {
	pkECDSA, err := crypto.ToECDSA(privateKey)
	if err != nil {
		return nil, WrapError(ErrInvalidPrivateKey, err)
	}

	pkECIES := ecies.ImportECDSA(pkECDSA)
	return pkECIES.Decrypt(bytes, nil, nil)
}

func DecryptString(str string, privateKey string) (string, error) {
	pkECDSA, err := crypto.ToECDSA(common.FromHex(privateKey))
	if err != nil {
		return "", WrapError(ErrInvalidPrivateKey, err)
	}

	pkECIES := ecies.ImportECDSA(pkECDSA)
	strDecodedBytes, err := hex.DecodeString(str)
	if err != nil {
		return "", err
	}

	strDecryptedBytes, err := pkECIES.Decrypt(strDecodedBytes, nil, nil)
	if err != nil {
		return "", err
	}

	return string(strDecryptedBytes), nil
}

func EncryptString(str string, publicKeyHex string) (string, error) {
	pubKeyBytes, err := hex.DecodeString(publicKeyHex)
	if err != nil {
		return "", err
	}

	urlBytes := []byte(str)

	publicKey, err := crypto.UnmarshalPubkey(pubKeyBytes)
	if err != nil {
		return "", err
	}

	pk := ecies.ImportECDSAPublic(publicKey)

	// Encrypt using ECIES
	ciphertext, err := ecies.Encrypt(rand.Reader, pk, urlBytes, nil, nil)
	if err != nil {
		return "", err
	}

	return hex.EncodeToString(ciphertext), nil
}

func PrivKeyToAddr(privateKey *ecdsa.PrivateKey) (common.Address, error) {
	publicKey := privateKey.Public()
	publicKeyECDSA, ok := publicKey.(*ecdsa.PublicKey)
	if !ok {
		return common.Address{}, fmt.Errorf("error casting public key to ECDSA")
	}

	return crypto.PubkeyToAddress(*publicKeyECDSA), nil
}

func PrivKeyStringToAddr(privateKey string) (common.Address, error) {
	return PrivKeyBytesToAddr(common.FromHex(privateKey))
}

func PrivKeyBytesToAddr(privateKey []byte) (common.Address, error) {
	privKey, err := crypto.ToECDSA(privateKey)
	if err != nil {
		return common.Address{}, WrapError(ErrInvalidPrivateKey, err)
	}

	addr, err := PrivKeyToAddr(privKey)
	if err != nil {
		return common.Address{}, WrapError(ErrInvalidPrivateKey, err)
	}
	return addr, nil
}

func PubKeyStringFromPrivate(privateKey string) (string, error) {
	publicKeyBytes, err := PubKeyFromPrivate(common.FromHex(privateKey))
	if err != nil {
		return "", err
	}
	return hex.EncodeToString(publicKeyBytes), nil
}

func PubKeyFromPrivate(privateKey HexString) (HexString, error) {
	privKey, err := crypto.ToECDSA(privateKey)
	if err != nil {
		return nil, WrapError(ErrInvalidPrivateKey, err)
	}

	pubKey := privKey.Public()
	pubKeyECDSA, ok := pubKey.(*ecdsa.PublicKey)
	if !ok {
		return nil, fmt.Errorf("error casting public key to ECDSA: %v", pubKeyECDSA)
	}
	return crypto.FromECDSAPub(pubKeyECDSA), nil
}

func MustPubKeyStringFromPrivate(privateKey string) string {
	pubKey, err := PubKeyStringFromPrivate(privateKey)
	if err != nil {
		panic(err)
	}
	return pubKey
}

func MustPrivKeyToAddr(privateKey *ecdsa.PrivateKey) common.Address {
	addr, err := PrivKeyToAddr(privateKey)
	if err != nil {
		panic(err)
	}
	return addr
}

func MustPrivKeyStringToAddr(privateKey string) common.Address {
	addr, err := PrivKeyStringToAddr(privateKey)
	if err != nil {
		panic(err)
	}
	return addr
}

// https://goethereumbook.org/signature-verify/
func VerifySignature(params []byte, signature []byte, publicKeyBytes []byte) bool {
	hash := crypto.Keccak256Hash(params)
	if len(signature) == 0 {
		return false
	}
	signatureNoRecoverID := signature[:len(signature)-1] // remove recovery ID
	return crypto.VerifySignature(publicKeyBytes, hash.Bytes(), signatureNoRecoverID)
}
