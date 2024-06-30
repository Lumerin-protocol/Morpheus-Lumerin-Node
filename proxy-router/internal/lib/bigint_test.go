package lib

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBigIntUnmarshallPointer(t *testing.T) {
	a := struct {
		F1 *BigInt
	}{}

	dataJSON := []byte(`{"F1":"123"}`)
	err := json.Unmarshal(dataJSON, &a)
	require.NoError(t, err)
	require.Equal(t, "123", a.F1.String())
}

func TestBigIntUnmarshallPointerNotSet(t *testing.T) {
	a := struct {
		F1 *BigInt
	}{}

	dataJSON := []byte(`{}`)
	err := json.Unmarshal(dataJSON, &a)
	require.NoError(t, err)
	require.Nil(t, a.F1)
}

func TestBigIntUnmarshallPointerNumber(t *testing.T) {
	a := struct {
		F1 *BigInt
	}{}

	dataJSON := []byte(`{"F1":123}`)
	err := json.Unmarshal(dataJSON, &a)
	require.NoError(t, err)
	require.Equal(t, "123", a.F1.String())
}

func TestBigIntUnmarshallValue(t *testing.T) {
	a := struct {
		F1 BigInt
	}{}

	dataJSON := []byte(`{"F1":"123"}`)
	err := json.Unmarshal(dataJSON, &a)
	require.NoError(t, err)
	require.Equal(t, "123", a.F1.String())
}

func TestBigIntUnmarshallValueNotSet(t *testing.T) {
	a := struct {
		F1 BigInt
	}{}

	dataJSON := []byte(`{}`)
	err := json.Unmarshal(dataJSON, &a)
	require.NoError(t, err)
	require.Equal(t, "0", a.F1.String())
}

func TestBigIntUnmarshallValueNumber(t *testing.T) {
	a := struct {
		F1 BigInt
	}{}

	dataJSON := []byte(`{"F1":123}`)
	err := json.Unmarshal(dataJSON, &a)
	require.NoError(t, err)
	require.Equal(t, "123", a.F1.String())
}

func TestBigIntMarshallPointer(t *testing.T) {
	a := struct {
		F1 *BigInt
	}{F1: &BigInt{}}
	a.F1.SetString("123", 10)

	dataJSON, err := json.Marshal(a)
	require.NoError(t, err)
	require.Equal(t, `{"F1":"123"}`, string(dataJSON))
}

func TestBigIntMarshallPointerNotSet(t *testing.T) {
	a := struct {
		F1 *BigInt
	}{}

	dataJSON, err := json.Marshal(a)
	require.NoError(t, err)
	require.Equal(t, `{"F1":null}`, string(dataJSON))
}

func TestBigIntMarshallValue(t *testing.T) {
	a := struct {
		F1 BigInt
	}{}
	a.F1.SetString("123", 10)

	dataJSON, err := json.Marshal(a)
	require.NoError(t, err)
	require.Equal(t, `{"F1":"123"}`, string(dataJSON))
}
