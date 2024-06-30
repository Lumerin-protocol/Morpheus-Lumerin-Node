package lib

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

type TestStructValue struct {
	F1 HexString
}

type TestStructPointer struct {
	F1 *HexString
}

func TestHexValue(t *testing.T) {
	js := []byte(`{"F1": "0x1123"}`)
	var target TestStructValue
	err := json.Unmarshal(js, &target)
	require.NoError(t, err)
	require.Equal(t, "0x1123", target.F1.String())
}

func TestHexValueEmpty(t *testing.T) {
	js := []byte(`{"F1": ""}`)
	var target TestStructValue
	err := json.Unmarshal(js, &target)
	require.NoError(t, err)
	require.Equal(t, "0x", target.F1.String())
}

func TestHexValueMissing(t *testing.T) {
	js := []byte(`{"F2": ""}`)
	var target TestStructValue
	err := json.Unmarshal(js, &target)
	require.NoError(t, err)
	require.Equal(t, "0x", target.F1.String())
}

func TestHexPointer(t *testing.T) {
	js := []byte(`{"F1": "0x1123"}`)
	var target TestStructPointer
	err := json.Unmarshal(js, &target)
	require.NoError(t, err)
	require.Equal(t, "0x1123", target.F1.String())
}

func TestHexPointerEmpty(t *testing.T) {
	js := []byte(`{"F1": ""}`)
	var target TestStructPointer
	err := json.Unmarshal(js, &target)
	require.NoError(t, err)
	require.Equal(t, "0x", target.F1.String())
}

func TestHexPointerMissing(t *testing.T) {
	js := []byte(`{"F2": ""}`)
	var target TestStructPointer
	err := json.Unmarshal(js, &target)
	require.NoError(t, err)
	require.Nil(t, target.F1)
}
