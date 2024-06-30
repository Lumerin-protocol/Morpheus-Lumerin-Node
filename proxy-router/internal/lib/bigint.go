package lib

import (
	"fmt"
	"math/big"
)

type BigInt struct {
	big.Int
}

// Enables the use of big.Int as a JSON field
func (b *BigInt) UnmarshalJSON(data []byte) error {
	// Remove quotes from the JSON string (if any)
	str := string(data)
	if len(str) > 0 && str[0] == '"' {
		str = str[1 : len(str)-1]
	}

	// Parse the string as a big.Int
	_, ok := b.Int.SetString(str, 10)
	if !ok {
		return fmt.Errorf("invalid big.Int string: %s", str)
	}
	return nil
}

func (b BigInt) MarshalJSON() ([]byte, error) {
	return []byte(fmt.Sprintf("\"%s\"", b.String())), nil
}

func (b *BigInt) Unpack() *big.Int {
	return &b.Int
}
