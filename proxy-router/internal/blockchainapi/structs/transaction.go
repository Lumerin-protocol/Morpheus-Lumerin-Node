package structs

type RawEthTransactionResponse struct {
	Status  string           `json:"status"`
	Message string           `json:"message"`
	Result  []RawTransaction `json:"result"`
}

type RawTransaction struct {
	BlockHash         string `json:"blockHash"`
	BlockNumber       string `json:"blockNumber"`
	Confirmations     string `json:"confirmations"`
	ContractAddress   string `json:"contractAddress"`
	CumulativeGasUsed string `json:"cumulativeGasUsed"`

	From             string `json:"from"`
	FunctionName     string `json:"functionName"`
	Gas              string `json:"gas"`
	GasPrice         string `json:"gasPrice"`
	GasPriceBid      string `json:"gasPriceBid"`
	GasUsed          string `json:"gasUsed"`
	Hash             string `json:"hash"`
	Input            string `json:"input"`
	IsError          string `json:"isError"`
	MethodId         string `json:"methodId"`
	Nonce            string `json:"nonce"`
	TimeStamp        string `json:"timeStamp"`
	To               string `json:"to"`
	TransactionIndex string `json:"transactionIndex"`
	TxReceiptStatus  string `json:"txreceipt_status"`
	Value            string `json:"value"`
}
