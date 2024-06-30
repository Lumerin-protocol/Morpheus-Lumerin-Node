package lib

import (
	"errors"
	"math/big"
	"math/rand"

	"github.com/ethereum/go-ethereum/common"
)

var (
	ErrInvalidAddress = errors.New("invalid address")
)

func GetRandomAddr() common.Address {
	return common.BigToAddress(big.NewInt(rand.Int63()))
}

// AddrShort returns a short representation of an address in "0x123..567" format
func AddrShort(addr string) string {
	return StrShortConf(addr, 5, 3)
}

func RemoveHexPrefix(s string) string {
	if len(s) >= 2 && s[0:2] == "0x" {
		return s[2:]
	}
	return s
}

type Address struct {
	common.Address
}

func StringToAddress(s string) (Address, error) {
	if !common.IsHexAddress(s) {
		return Address{}, ErrInvalidAddress
	}
	return Address{common.HexToAddress(s)}, nil
}

func MustStringToAddress(s string) Address {
	a, err := StringToAddress(s)
	if err != nil {
		panic(err)
	}
	return a
}

func (a *Address) UnmarshalParam(param string) error {
	if param == "" {
		return nil
	}
	addr, err := StringToAddress(param)
	if err != nil {
		return err
	}
	*a = addr
	return nil
}
