package lib

import (
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/contracts/bindings/lumerintoken"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/contracts/bindings/marketplace"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/contracts/bindings/modelregistry"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/contracts/bindings/morpheustoken"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/contracts/bindings/providerregistry"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/contracts/bindings/sessionrouter"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/common"
)

type Meta interface {
	GetAbi() (*abi.ABI, error)
}

var allContractsMeta = []Meta{
	marketplace.MarketplaceMetaData,
	modelregistry.ModelRegistryMetaData,
	providerregistry.ProviderRegistryMetaData,
	sessionrouter.SessionRouterMetaData,
	morpheustoken.MorpheusTokenMetaData,
	lumerintoken.LumerinTokenMetaData,
}

type EVMError struct {
	Abi   abi.Error
	Cause error
	Args  interface{}
}

func (e EVMError) Error() string {
	idBytes := e.Abi.ID.Bytes()
	if len(idBytes) > 4 {
		idBytes = idBytes[:4]
	}
	return "EVM error: " + e.Abi.Sig + " " + common.BytesToHash(idBytes).Hex()
}

// TryConvertGethError attempts to convert geth error to an EVMError, otherwise just returns original error
func TryConvertGethError(err error) error {
	evmErr, ok := ConvertGethError(err, allContractsMeta)
	if !ok {
		return err
	}
	return evmErr
}

// ConvertGethError converts a geth error to an EVMError with exposed error signature and arguments
func ConvertGethError(gethErr error, contractMeta []Meta) (*EVMError, bool) {
	errData, ok := ExtractGETHErrorData(gethErr)
	if !ok {
		return nil, false
	}

	abis := make([]*abi.ABI, len(contractMeta))
	for i, meta := range contractMeta {
		abi, err := meta.GetAbi()
		if err != nil {
			return nil, false
		}
		abis[i] = abi
	}

	abiError, args, ok := CastErrorData(errData, abis)
	if !ok {
		return nil, false
	}

	return &EVMError{
		Abi:   abiError,
		Args:  args,
		Cause: gethErr,
	}, true
}

// ExtractGETHErrorData extracts the error data from an unexproted geth error
func ExtractGETHErrorData(err error) ([]byte, bool) {
	asErr, ok := err.(interface{ ErrorData() interface{} })
	if !ok {
		return nil, false
	}
	errDataHex, ok := asErr.ErrorData().(string)
	if !ok {
		return nil, false
	}
	errDataBytes := common.FromHex(errDataHex)
	if len(errDataBytes) < 4 {
		return nil, false
	}
	return errDataBytes, true
}

// CastErrorData casts the error data to the appropriate error type
func CastErrorData(errData []byte, abis []*abi.ABI) (abi.Error, interface{}, bool) {
	for _, abiData := range abis {
		for _, abiError := range abiData.Errors {
			args, err := abiError.Unpack(errData)
			if err == nil {
				return abiError, args, true
			}
		}
	}
	return abi.Error{}, nil, false
}
