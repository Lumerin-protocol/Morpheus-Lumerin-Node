package lib

import (
	"fmt"

	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/crypto"
)

type AbiParameter struct {
	Type string
}

func EncodeAbiParameters(abiParams []AbiParameter, params []interface{}) ([]byte, error) {
	// Construct the ABI arguments
	arguments := make(abi.Arguments, len(abiParams))
	for i, param := range abiParams {
		argType, err := abi.NewType(param.Type, "", nil)
		if err != nil {
			return nil, err
		}
		arguments[i] = abi.Argument{Type: argType}
	}

	// Pack the parameters into a byte array
	return arguments.Pack(params...)
}

func SignEthMessage(msg []byte, privateKeyHex string) ([]byte, error) {
	privateKey, err := crypto.HexToECDSA(privateKeyHex)
	if err != nil {
		return nil, err
	}
	hash := crypto.Keccak256Hash(msg)

	prefixStr := fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(hash.Bytes()))
	message := append([]byte(prefixStr), hash.Bytes()...)
	resultHash := crypto.Keccak256Hash(message)

	signature, err := crypto.Sign(resultHash.Bytes(), privateKey)
	if err != nil {
		return nil, err
	}

	// https://github.com/ethereum/go-ethereum/blob/44a50c9f96386f44a8682d51cf7500044f6cbaea/ethapi/api.go#L580
	signature[64] += 27 // Transform V from 0/1 to 27/28
	return signature, nil
}

func SignEthMessageV2(msg []byte, privateKeyBytes []byte) ([]byte, error) {
	privateKey, err := crypto.ToECDSA(privateKeyBytes)
	if err != nil {
		return nil, err
	}
	hash := crypto.Keccak256Hash(msg)

	prefixStr := fmt.Sprintf("\x19Ethereum Signed Message:\n%d", len(hash.Bytes()))
	message := append([]byte(prefixStr), hash.Bytes()...)
	resultHash := crypto.Keccak256Hash(message)

	signature, err := crypto.Sign(resultHash.Bytes(), privateKey)
	if err != nil {
		return nil, err
	}

	// https://github.com/ethereum/go-ethereum/blob/44a50c9f96386f44a8682d51cf7500044f6cbaea/ethapi/api.go#L580
	signature[64] += 27 // Transform V from 0/1 to 27/28
	return signature, nil
}
