package blockchainapi

import (
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/blockchainapi/structs"
	"github.com/ethereum/go-ethereum/common"
)

type ExplorerClient struct {
	explorerApiUrl string
	morTokenAddr   string
}

func NewExplorerClient(explorerApiUrl string, morTokenAddr string) *ExplorerClient {
	return &ExplorerClient{
		explorerApiUrl: explorerApiUrl,
		morTokenAddr:   morTokenAddr,
	}
}

func (e *ExplorerClient) GetEthTransactions(address common.Address, page uint64, limit uint8) ([]structs.RawTransaction, error) {
	query := fmt.Sprintf("?module=account&action=txlist&address=%s&page=%d&offset=%d", address.Hex(), page, limit)
	url := e.explorerApiUrl + query

	transactions, err := e.doRequest(url)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (e *ExplorerClient) GetTokenTransactions(address common.Address, page uint64, limit uint8) ([]structs.RawTransaction, error) {
	query := fmt.Sprintf("?module=account&action=tokentx&contractaddress=%s&address=%s&page=%d&offset=%d", e.morTokenAddr, address.Hex(), page, limit)
	url := e.explorerApiUrl + query

	transactions, err := e.doRequest(url)
	if err != nil {
		return nil, err
	}
	return transactions, nil
}

func (e *ExplorerClient) doRequest(url string) ([]structs.RawTransaction, error) {
	req, err := http.NewRequest("GET", url, nil)
	if err != nil {
		return nil, err
	}

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return nil, fmt.Errorf("failed to fetch eth transactions, HTTP status code: %d", resp.StatusCode)
	}

	var response structs.RawEthTransactionResponse
	if err := json.NewDecoder(resp.Body).Decode(&response); err != nil {
		return nil, err
	}
	if response.Status == "0" {
		return make([]structs.RawTransaction, 0), nil
	}

	if response.Status != "1" {
		return nil, fmt.Errorf("failed to fetch eth transactions, response status: %s", response.Result)
	}

	return response.Result, nil
}
