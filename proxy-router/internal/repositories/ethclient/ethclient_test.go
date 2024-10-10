package ethclient

import (
	"fmt"
	"sync"
	"testing"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/contracts/modelregistry"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/ethereum/go-ethereum/common"
	"github.com/stretchr/testify/require"
)

func TestEthClient(t *testing.T) {
	t.SkipNow()
	rpcClient, err := NewRPCClientMultiple([]string{
		"https://arbitrum.blockpi.network/v1/rpc/public",
	}, &lib.LoggerMock{})
	require.NoError(t, err)
	client := NewClient(rpcClient)

	wg := sync.WaitGroup{}

	mr, err := modelregistry.NewModelRegistry(common.HexToAddress("0x0FC0c323Cf76E188654D63D62e668caBeC7a525b"), client)
	require.NoError(t, err)

	for i := 0; i < 500; i++ {
		wg.Add(1)
		go func() {
			ids, _, err := mr.ModelGetAll(nil)
			defer wg.Done()
			if err != nil {
				fmt.Printf("Error: %s\n", err)
				return
			}
			fmt.Printf("Models number: %d\n", len(ids))
		}()

	}
	wg.Wait()
}
