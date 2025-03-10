package lib

import (
	"fmt"

	cid "github.com/ipfs/go-cid"
	"github.com/multiformats/go-multihash"
)

func CIDToBytes32(cidStr string) ([]byte, error) {
	c, err := cid.Decode(cidStr)
	if err != nil {
		return nil, fmt.Errorf("invalid CID: %v", err)
	}
	mh := c.Hash()
	decoded, err := multihash.Decode(mh)
	if err != nil {
		return nil, fmt.Errorf("failed to decode multihash: %v", err)
	}

	// decoded.Digest is the raw 32-byte SHA-256 hash
	digest := decoded.Digest
	if len(digest) != 32 {
		return nil, fmt.Errorf("digest length is not 32: got %d", len(digest))
	}

	bytes32 := make([]byte, 32)
	copy(bytes32, digest)

	return bytes32, nil
}

// ManualBytes32ToCID reconstructs a CIDv0 from a raw SHA-256 digest
func ManualBytes32ToCID(digest []byte) (string, error) {
	if len(digest) != 32 {
		return "", fmt.Errorf("invalid digest length: expected 32 bytes, got %d", len(digest))
	}

	// Prepend the multihash prefix (0x12 = SHA-256, 0x20 = 32 bytes length)
	multihash := append([]byte{0x12, 0x20}, digest...)

	// Create CIDv0 directly
	c := cid.NewCidV0(multihash)

	return c.String(), nil
}
