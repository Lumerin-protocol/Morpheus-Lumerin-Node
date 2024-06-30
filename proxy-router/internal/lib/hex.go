package lib

import (
	"encoding"
	"fmt"
	"regexp"

	"github.com/ethereum/go-ethereum/common"
)

var (
	// ethAddressRegexString = `^0x[0-9a-fA-F]{40}$`
	// ethAddressRegex       = regexp.MustCompile(ethAddressRegexString)

	hexadecimalRegexString = "^(0[xX])?[0-9a-fA-F]+$"
	hexadecimalRegex       = regexp.MustCompile(hexadecimalRegexString)

	ErrInvalidHexString = fmt.Errorf("invalid hex string")
)

// Arbitrary length hex string represented internally as a byte slice
type HexString []byte

func StringToHexString(s string) (HexString, error) {
	if !hexadecimalRegex.MatchString(s) {
		return nil, WrapError(ErrInvalidHexString, fmt.Errorf(`"%s"`, s))
	}
	return common.FromHex(s), nil
}

func MustStringToHexString(s string) HexString {
	h, err := StringToHexString(s)
	if err != nil {
		panic(err)
	}
	return h
}

func (b *HexString) UnmarshalJSON(data []byte) error {
	str := string(data)
	if len(str) > 0 && str[0] == '"' {
		str = str[1 : len(str)-1]
	}
	if len(str) == 0 {
		*b = common.FromHex("")
		return nil
	}
	if !hexadecimalRegex.MatchString(str) {
		return WrapError(ErrInvalidHexString, fmt.Errorf(`"%s"`, str))
	}
	*b = common.FromHex(str)
	return nil
}

func (b *HexString) UnmarshalText(data []byte) error {
	dataStr := string(data)
	if !hexadecimalRegex.MatchString(dataStr) {
		return WrapError(ErrInvalidHexString, fmt.Errorf(`"%s"`, dataStr))
	}
	d := common.FromHex(string(data))
	*b = d
	return nil
}

func (b HexString) String() string {
	return b.Hex()
}

func (b HexString) Hex() string {
	return BytesToString(b)
}

func (b HexString) MarshalJSON() ([]byte, error) {
	return []byte(`"` + b.String() + `"`), nil
}

var _ encoding.TextUnmarshaler = &HexString{}
