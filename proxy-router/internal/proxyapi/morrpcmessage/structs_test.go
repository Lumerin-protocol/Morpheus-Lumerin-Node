package morrpcmesssage

import (
	"encoding/json"
	"testing"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/config"
	"github.com/go-playground/validator/v10"
	"github.com/stretchr/testify/require"
)

func TestStruct(t *testing.T) {
	// Test for SessionReq struct
	var sessionReq SessionReq
	msg := []byte(`{"bidid":"0xb0119306909b29feb8a2a30f343c32f7ceb0927325a4a6e4936d11d1236556c3","key":"0x049d9031e97dd78ff8c15aa86939de9b1e791066a0224e331bc962a2099a7b1f0464b8bbafe1535f2301c72c2cb3535b172da30b02686ab0393d348614f157fbdb","provider":"0x70997970c51812dc3a010c7d01b50e0d17dc79c8","signature":"0xa2bcdbd6ab6bf137c9372488139b9c4ed42503bb145f6ee9f72efd3eae6133934f0f24d8f09a8f05d057e80f8ef320d74abef9fdc915e47cdc06c55f51d78a4e01","spend":"270233162839357259772","timestamp":1719253477054,"user":"0x3c44cdddb6a900fa2b585dd299e03d12fa4293bc"}`)
	err := json.Unmarshal(msg, &sessionReq)
	require.NoError(t, err)

	val := validator.New()
	err = config.RegisterHex32(val)
	require.NoError(t, err)
	err = config.RegisterDuration(val)
	require.NoError(t, err)
	err = config.RegisterEthAddr(val)
	require.NoError(t, err)
	err = config.RegisterHexadecimal(val)
	require.NoError(t, err)

	err = val.Struct(sessionReq)
	require.NoError(t, err)
}
