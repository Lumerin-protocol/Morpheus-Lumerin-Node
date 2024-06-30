package lib

import (
	"encoding/hex"
	"errors"

	"github.com/ethereum/go-ethereum/common"
)

var (
	ErrInvalidHashLen = errors.New("invalid hash length")
)

type Hash struct {
	common.Hash
}

func StringToHash(s string) (Hash, error) {
	hs, err := HexToHash(s)
	if err != nil {
		return Hash{}, err
	}
	return Hash{hs}, nil
}

func MustStringToHash(s string) Hash {
	hs, err := StringToHash(s)
	if err != nil {
		panic(err)
	}
	return hs
}

func (h *Hash) UnmarshalParam(param string) error {
	if param == "" {
		return nil
	}
	hs, err := HexToHash(param)
	if err != nil {
		return err
	}

	h.Hash = hs
	return nil
}

func HexToHash(s string) (common.Hash, error) {
	if has0xPrefix(s) {
		s = s[2:]
	}
	if len(s)%2 == 1 {
		s = "0" + s
	}
	bytes, err := hex.DecodeString(s)
	if err != nil {
		return common.Hash{}, err
	}
	if len(bytes) != common.HashLength {
		return common.Hash{}, ErrInvalidHashLen
	}
	return common.BytesToHash(bytes), nil
}

func has0xPrefix(str string) bool {
	return len(str) >= 2 && str[0] == '0' && (str[1] == 'x' || str[1] == 'X')
}
