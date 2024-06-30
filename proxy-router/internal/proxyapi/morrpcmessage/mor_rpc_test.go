package morrpcmesssage

import (
	"encoding/hex"
	"fmt"
	"testing"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/stretchr/testify/assert"
)

func TestMorRpc_generateSignature(t *testing.T) {
	m := NewMorRpc()

	params := map[string]interface{}{
		"param1": "value1",
		"param2": "value2",
		"param3": "value3",
	}

	privateKeyHex := lib.MustStringToHexString("3ceb688d9b87c1a468a7eadde744828ec8bb2d11c9ea52a179058e47f92f25ee")

	signature, err := m.generateSignature(params, privateKeyHex)
	assert.NoError(t, err)
	assert.NotEmpty(t, signature)
}

func TestMorRpc_verifySignature(t *testing.T) {
	m := NewMorRpc()

	params := map[string]interface{}{
		"param1": "value1",
		"param2": "value2",
		"param3": "value3",
	}

	privateKeyHex := lib.MustStringToHexString("81f44a49c40f206517efbbcca783d808914841200e0ac9a769368e1b2741e227")
	publicKey := lib.MustStringToHexString("033e5e77f12aa67e52484ce64b64737d397098e78d54beba15a0bf6dcfdd5ae7e2")

	signature, err := m.generateSignature(params, privateKeyHex)
	assert.NoError(t, err)
	assert.NotEmpty(t, signature)

	isValid := m.VerifySignature(params, signature, publicKey, nil)
	assert.True(t, isValid)
}

func TestMorRpc_verifySignature_incorrect_params(t *testing.T) {
	m := NewMorRpc()

	params := map[string]interface{}{
		"param1": "value1",
		"param2": "value2",
		"param3": "value3",
	}

	privateKeyHex := lib.MustStringToHexString("81f44a49c40f206517efbbcca783d808914841200e0ac9a769368e1b2741e227")
	publicKey := lib.MustStringToHexString("033e5e77f12aa67e52484ce64b64737d397098e78d54beba15a0bf6dcfdd5ae7e2")

	signature, err := m.generateSignature(params, privateKeyHex)
	assert.NoError(t, err)
	assert.NotEmpty(t, signature)

	params["param3"] = "unknown value"
	isValid := m.VerifySignature(params, signature, publicKey, nil)
	assert.False(t, isValid)
}

func TestMorRpc_generate(t *testing.T) {
	m := NewMorRpc()

	params := map[string]interface{}{
		"user":      "2222",
		"key":       "033e5e77f12aa67e52484ce64b64737d397098e78d54beba15a0bf6dcfdd5ae7e2",
		"spend":     "10",
		"provider":  "1111",
		"timestamp": "1234567890",
	}

	privateKeyHex := lib.MustStringToHexString("81f44a49c40f206517efbbcca783d808914841200e0ac9a769368e1b2741e227")
	// publicKey := "033e5e77f12aa67e52484ce64b64737d397098e78d54beba15a0bf6dcfdd5ae7e2"

	signature, err := m.generateSignature(params, privateKeyHex)
	assert.NoError(t, err)

	hexSignature := hex.EncodeToString([]byte(signature))
	fmt.Println(hexSignature)
}
