package rpcproxy

import (
	"context"
	"encoding/hex"
	"errors"
	"fmt"
	"math/big"
	"sort"
	"strconv"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/contracts/sessionrouter"
	constants "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/internal"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/internal/interfaces"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/internal/lib"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/internal/repositories/registries"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/internal/rpcproxy/structs"
	"github.com/gin-gonic/gin"

	"github.com/ethereum/go-ethereum"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
	"github.com/ethereum/go-ethereum/core/types"
	"github.com/ethereum/go-ethereum/crypto"
	"github.com/ethereum/go-ethereum/ethclient"
)

type RpcProxy struct {
	rpcClient        *ethclient.Client
	providerRegistry *registries.ProviderRegistry
	modelRegistry    *registries.ModelRegistry
	marketplace      *registries.Marketplace
	sessionRouter    *registries.SessionRouter
	morToken         *registries.MorToken
	explorerClient   *ExplorerClient

	legacyTx   bool
	privateKey string
}

func NewRpcProxy(rpcClient *ethclient.Client, diamonContractAddr common.Address, morTokenAddr common.Address, explorerApiUrl string, privateKey string, log interfaces.ILogger, legacyTx bool) *RpcProxy {
	providerRegistry := registries.NewProviderRegistry(diamonContractAddr, rpcClient, log)
	modelRegistry := registries.NewModelRegistry(diamonContractAddr, rpcClient, log)
	marketplace := registries.NewMarketplace(diamonContractAddr, rpcClient, log)
	sessionRouter := registries.NewSessionRouter(diamonContractAddr, rpcClient, log)
	morToken := registries.NewMorToken(morTokenAddr, rpcClient, log)

	explorerClient := NewExplorerClient(explorerApiUrl, morTokenAddr.String())
	return &RpcProxy{
		rpcClient:        rpcClient,
		providerRegistry: providerRegistry,
		modelRegistry:    modelRegistry,
		marketplace:      marketplace,
		sessionRouter:    sessionRouter,
		legacyTx:         legacyTx,
		privateKey:       privateKey,
		morToken:         morToken,
		explorerClient:   explorerClient,
	}
}

func (rpcProxy *RpcProxy) GetLatestBlock(ctx context.Context) (uint64, error) {
	return rpcProxy.rpcClient.BlockNumber(ctx)
}

func (rpcProxy *RpcProxy) GetAllProviders(ctx context.Context) (int, gin.H) {
	addrs, providers, err := rpcProxy.providerRegistry.GetAllProviders(ctx)

	fmt.Println("provider registry results - addrs: ", addrs, " providers: ", providers, " err: ", err)
	if err != nil {
		return constants.HTTP_STATUS_BAD_REQUEST, gin.H{"error": err.Error()}
	}

	result := make([]*structs.Provider, len(addrs))
	for i, value := range providers {
		result[i] = &structs.Provider{
			Address:   addrs[i],
			Endpoint:  value.Endpoint,
			Stake:     value.Stake,
			IsDeleted: value.IsDeleted,
			CreatedAt: value.CreatedAt,
		}
	}

	return constants.HTTP_STATUS_OK, gin.H{"providers": result}
}

func (rpcProxy *RpcProxy) CreateNewProvider(ctx context.Context, address string, addStake uint64, endpoint string) (int, gin.H) {	
	fmt.Println("CreateNewProvider address: ", address, " addStake: ", addStake, " endpoint: ", endpoint)
	err := rpcProxy.providerRegistry.CreateNewProvider(ctx, address, addStake, endpoint)
	if err != nil {
		return constants.HTTP_STATUS_BAD_REQUEST, gin.H{"error": err.Error()}
	}
	return constants.HTTP_STATUS_OK, gin.H{"success": true}
}

func (rpcProxy *RpcProxy) GetAllModels(ctx context.Context) (int, gin.H) {
	ids, models, err := rpcProxy.modelRegistry.GetAllModels(ctx)
	if err != nil {
		return constants.HTTP_STATUS_BAD_REQUEST, gin.H{"error": err.Error()}
	}

	result := make([]*structs.Model, len(ids))
	for i, value := range models {
		result[i] = &structs.Model{
			Id:        lib.BytesToString(ids[i][:]),
			IpfsCID:   lib.BytesToString(value.IpfsCID[:]),
			Fee:       value.Fee,
			Stake:     value.Stake,
			Owner:     value.Owner,
			Name:      value.Name,
			Tags:      value.Tags,
			CreatedAt: value.CreatedAt,
			IsDeleted: value.IsDeleted,
		}
	}

	return constants.HTTP_STATUS_OK, gin.H{"models": result}
}

func (rpcProxy *RpcProxy) GetBidsByProvider(ctx context.Context, providerAddr common.Address, offset *big.Int, limit uint8) (int, gin.H) {
	ids, bids, err := rpcProxy.marketplace.GetBidsByProvider(ctx, providerAddr, offset, limit)
	if err != nil {
		return constants.HTTP_STATUS_BAD_REQUEST, gin.H{"error": err.Error()}
	}

	result := make([]*structs.Bid, len(ids))
	for i, value := range bids {
		result[i] = &structs.Bid{
			Id:             lib.BytesToString(ids[i][:]),
			ModelAgentId:   lib.BytesToString(value.ModelAgentId[:]),
			Provider:       value.Provider,
			Nonce:          value.Nonce,
			CreatedAt:      value.CreatedAt,
			DeletedAt:      value.DeletedAt,
			PricePerSecond: value.PricePerSecond,
		}
	}
	return constants.HTTP_STATUS_OK, gin.H{"bids": result}
}

func (rpcProxy *RpcProxy) GetBidsByModelAgent(ctx context.Context, modelId [32]byte, offset *big.Int, limit uint8) (int, gin.H) {
	ids, bids, err := rpcProxy.marketplace.GetBidsByModelAgent(ctx, modelId, offset, limit)
	if err != nil {
		return constants.HTTP_STATUS_BAD_REQUEST, gin.H{"error": err.Error()}
	}

	result := make([]*structs.Bid, len(ids))
	for i, value := range bids {
		result[i] = &structs.Bid{
			Id:             lib.BytesToString(ids[i][:]),
			ModelAgentId:   lib.BytesToString(value.ModelAgentId[:]),
			Provider:       value.Provider,
			Nonce:          value.Nonce,
			CreatedAt:      value.CreatedAt,
			DeletedAt:      value.DeletedAt,
			PricePerSecond: value.PricePerSecond,
		}
	}

	return constants.HTTP_STATUS_OK, gin.H{"bids": result}
}

type OpenSessionRequest struct {
	Approval    string `json:"approval"`
	ApprovalSig string `json:"approvalSig"`
	Stake       string `json:"stake"`
}

type SendRequest struct {
	To     string `json:"to"`
	Amount string `json:"amount"`
}

func (rpcProxy *RpcProxy) OpenSession(ctx *gin.Context) (int, gin.H) {
	var reqPayload OpenSessionRequest
	if err := ctx.ShouldBindJSON(&reqPayload); err != nil {
		return constants.HTTP_STATUS_BAD_REQUEST, gin.H{"error": err.Error()}
	}

	approval := reqPayload.Approval
	if approval == "" {
		return constants.HTTP_STATUS_BAD_REQUEST, gin.H{"error": "approval is required"}
	}

	approvalSig := reqPayload.ApprovalSig
	if approvalSig == "" {
		return constants.HTTP_STATUS_BAD_REQUEST, gin.H{"error": "approvalSig is required"}
	}

	stakeStr := reqPayload.Stake
	if stakeStr == "" {
		return constants.HTTP_STATUS_BAD_REQUEST, gin.H{"error": "stake is required"}
	}

	stake, ok := new(big.Int).SetString(stakeStr, 10)
	if !ok {
		return constants.HTTP_STATUS_BAD_REQUEST, gin.H{"error": "stake is invalid"}
	}

	transactOpt, err := rpcProxy.getTransactOpts(ctx, rpcProxy.privateKey)
	if err != nil {
		return constants.HTTP_INTERNAL_SERVER_ERROR, gin.H{"error": err.Error()}
	}

	approvalBytes := common.FromHex(reqPayload.Approval)
	approvalSigBytes := common.FromHex(reqPayload.ApprovalSig)

	sessionId, err := rpcProxy.sessionRouter.OpenSession(transactOpt, approvalBytes, approvalSigBytes, stake, rpcProxy.privateKey)
	if err != nil {
		return constants.HTTP_STATUS_BAD_REQUEST, gin.H{"error": err.Error()}
	}

	return constants.HTTP_STATUS_OK, gin.H{"sessionId": sessionId}
}

func (rpcProxy *RpcProxy) CloseSession(ctx *gin.Context) (int, gin.H) {
	sessionId := ctx.Param("id")

	if sessionId == "" {
		return constants.HTTP_STATUS_BAD_REQUEST, gin.H{"error": "sessionId is required"}
	}

	transactOpt, err := rpcProxy.getTransactOpts(ctx, rpcProxy.privateKey)
	if err != nil {
		return constants.HTTP_INTERNAL_SERVER_ERROR, gin.H{"error": err.Error()}
	}

	_, err = rpcProxy.sessionRouter.CloseSession(transactOpt, sessionId, rpcProxy.privateKey)
	if err != nil {
		return constants.HTTP_STATUS_BAD_REQUEST, gin.H{"error": err.Error()}
	}

	return constants.HTTP_STATUS_OK, gin.H{"success": true}
}

func (rpcProxy *RpcProxy) GetSession(ctx *gin.Context, sessionId string) (int, gin.H) {
	session, err := rpcProxy.sessionRouter.GetSession(ctx, sessionId)
	if err != nil {
		return constants.HTTP_STATUS_BAD_REQUEST, gin.H{"error": err.Error()}
	}
	return constants.HTTP_STATUS_OK, gin.H{"session": session}
}

func (rpc *RpcProxy) GetProviderClaimableBalance(ctx *gin.Context) (int, gin.H) {
	sessionId := ctx.Param("id")
	if sessionId == "" {
		return constants.HTTP_STATUS_BAD_REQUEST, gin.H{"error": "sessionId is required"}
	}

	balance, err := rpc.sessionRouter.GetProviderClaimableBalance(ctx, sessionId)
	if err != nil {
		return constants.HTTP_STATUS_BAD_REQUEST, gin.H{"error": err.Error()}
	}
	return constants.HTTP_STATUS_OK, gin.H{"balance": balance}
}

func (rpcProxy *RpcProxy) GetBalance(ctx *gin.Context) (int, gin.H) {
	transactOpt, err := rpcProxy.getTransactOpts(ctx, rpcProxy.privateKey)
	if err != nil {
		return constants.HTTP_INTERNAL_SERVER_ERROR, gin.H{"error": err.Error()}
	}

	ethBalance, err := rpcProxy.rpcClient.BalanceAt(ctx, transactOpt.From, nil)
	if err != nil {
		return constants.HTTP_INTERNAL_SERVER_ERROR, gin.H{"error": "failed to get eth balance: " + err.Error()}
	}

	balance, err := rpcProxy.morToken.GetBalance(ctx, transactOpt.From)
	if err != nil {
		return constants.HTTP_INTERNAL_SERVER_ERROR, gin.H{"error": "failed to get mor balance: " + err.Error()}
	}

	return constants.HTTP_STATUS_OK, gin.H{"eth": ethBalance.String(), "mor": balance.String()}
}

func (rpcProxy *RpcProxy) SendEth(ctx *gin.Context) (int, gin.H) {
	to, amount, err := rpcProxy.getSendParams(ctx)
	if err != nil {
		return constants.HTTP_STATUS_BAD_REQUEST, gin.H{"error": err.Error()}
	}

	transactOpt, err := rpcProxy.getTransactOpts(ctx, rpcProxy.privateKey)
	if err != nil {
		return constants.HTTP_INTERNAL_SERVER_ERROR, gin.H{"error": err.Error()}
	}

	nonce, err := rpcProxy.rpcClient.PendingNonceAt(context.Background(), transactOpt.From)
	if err != nil {
		return constants.HTTP_INTERNAL_SERVER_ERROR, gin.H{"error": "failed to get nonce: " + err.Error()}
	}

	toAddr := common.HexToAddress(to)
	estimatedGas, err := rpcProxy.rpcClient.EstimateGas(context.Background(), ethereum.CallMsg{
		From:  transactOpt.From,
		To:    &toAddr,
		Value: amount,
	})
	if err != nil {
		return constants.HTTP_INTERNAL_SERVER_ERROR, gin.H{"error": "failed to estimate gas: " + err.Error()}
	}

	gas := float64(estimatedGas) * 1.5
	tx := types.NewTransaction(nonce, toAddr, amount, uint64(gas), transactOpt.GasPrice, nil)
	signedTx, err := rpcProxy.signTx(ctx, tx, rpcProxy.privateKey)
	if err != nil {
		return constants.HTTP_INTERNAL_SERVER_ERROR, gin.H{"error": "failed to sign eth: " + err.Error()}
	}

	err = rpcProxy.rpcClient.SendTransaction(context.Background(), signedTx)
	if err != nil {
		return constants.HTTP_INTERNAL_SERVER_ERROR, gin.H{"error": "failed to send eth: " + err.Error()}
	}

	// Wait for the transaction receipt
	_, err = bind.WaitMined(context.Background(), rpcProxy.rpcClient, signedTx)
	if err != nil {
		return constants.HTTP_INTERNAL_SERVER_ERROR, gin.H{"error": "failed to send eth: " + err.Error()}
	}

	return constants.HTTP_STATUS_OK, gin.H{"txHash": signedTx.Hash().String()}
}

func (rpcProxy *RpcProxy) SendMor(ctx *gin.Context) (int, gin.H) {
	to, amount, err := rpcProxy.getSendParams(ctx)
	if err != nil {
		return constants.HTTP_STATUS_BAD_REQUEST, gin.H{"error": err.Error()}
	}

	transactOpt, err := rpcProxy.getTransactOpts(ctx, rpcProxy.privateKey)
	if err != nil {
		return constants.HTTP_INTERNAL_SERVER_ERROR, gin.H{"error": err.Error()}
	}

	tx, err := rpcProxy.morToken.Transfer(transactOpt, common.HexToAddress(to), amount)
	if err != nil {
		return constants.HTTP_INTERNAL_SERVER_ERROR, gin.H{"error": "failed to transfer mor: " + err.Error()}
	}

	return constants.HTTP_STATUS_OK, gin.H{"txHash": tx.Hash().String()}
}

func (rpcProxy *RpcProxy) GetAllowance(ctx *gin.Context) (int, gin.H) {
	spender := ctx.Query("spender")

	if spender == "" {
		return constants.HTTP_STATUS_BAD_REQUEST, gin.H{"error": "spender is required"}
	}

	spenderAddr := common.HexToAddress(spender)

	transactOpt, err := rpcProxy.getTransactOpts(ctx, rpcProxy.privateKey)
	if err != nil {
		return constants.HTTP_INTERNAL_SERVER_ERROR, gin.H{"error": "failed to get transactOpts: " + err.Error()}
	}

	allowance, err := rpcProxy.morToken.GetAllowance(ctx, transactOpt.From, spenderAddr)
	if err != nil {
		return constants.HTTP_INTERNAL_SERVER_ERROR, gin.H{"error": "failed to get allowance: " + err.Error()}
	}

	return constants.HTTP_STATUS_OK, gin.H{"allowance": allowance.String()}
}

func (rpcProxy *RpcProxy) Approve(ctx *gin.Context) (int, gin.H) {
	spender := ctx.Query("spender")

	if spender == "" {
		return constants.HTTP_STATUS_BAD_REQUEST, gin.H{"error": "spender is required"}
	}

	spenderAddr := common.HexToAddress(spender)

	amount := ctx.Query("amount")
	if amount == "" {
		return constants.HTTP_STATUS_BAD_REQUEST, gin.H{"error": "amount is required"}
	}

	amountInt, ok := new(big.Int).SetString(amount, 10)
	if !ok {
		return constants.HTTP_STATUS_BAD_REQUEST, gin.H{"error": "invalid amount"}
	}

	transactOpt, err := rpcProxy.getTransactOpts(ctx, rpcProxy.privateKey)
	if err != nil {
		return constants.HTTP_INTERNAL_SERVER_ERROR, gin.H{"error": "failed to get transactOpts: " + err.Error()}
	}

	tx, err := rpcProxy.morToken.Approve(transactOpt, spenderAddr, amountInt)
	if err != nil {
		return constants.HTTP_INTERNAL_SERVER_ERROR, gin.H{"error": "failed to approve: " + err.Error()}
	}

	return constants.HTTP_STATUS_OK, gin.H{"tx": tx}
}

func (rpcProxy *RpcProxy) GetTodaysBudget(ctx *gin.Context) (int, gin.H) {
	budget, err := rpcProxy.sessionRouter.GetTodaysBudget(ctx)
	if err != nil {
		return constants.HTTP_INTERNAL_SERVER_ERROR, gin.H{"error": "failed to get budget: " + err.Error()}
	}

	return constants.HTTP_STATUS_OK, gin.H{"budget": budget.String()}
}

func (rpcProxy *RpcProxy) ClaimProviderBalance(ctx *gin.Context) (int, gin.H) {
	sessionId := ctx.Param("id")
	if sessionId == "" {
		return constants.HTTP_STATUS_BAD_REQUEST, gin.H{"error": "sessionId is required"}
	}

	to, amount, err := rpcProxy.getSendParams(ctx)
	if err != nil {
		return constants.HTTP_STATUS_BAD_REQUEST, gin.H{"error": err.Error()}
	}

	transactOpt, err := rpcProxy.getTransactOpts(ctx, rpcProxy.privateKey)
	if err != nil {
		return constants.HTTP_INTERNAL_SERVER_ERROR, gin.H{"error": err.Error()}
	}

	txHash, err := rpcProxy.sessionRouter.ClaimProviderBalance(transactOpt, sessionId, amount, common.HexToAddress(to))
	if err != nil {
		return constants.HTTP_STATUS_BAD_REQUEST, gin.H{"error": err.Error()}
	}

	return constants.HTTP_STATUS_OK, gin.H{"txHash": txHash}
}

func (rpcProxy *RpcProxy) GetTokenSupply(ctx *gin.Context) (int, gin.H) {
	supply, err := rpcProxy.morToken.GetTotalSupply(ctx)
	if err != nil {
		return constants.HTTP_INTERNAL_SERVER_ERROR, gin.H{"error": "failed to get supply: " + err.Error()}
	}
	return constants.HTTP_STATUS_OK, gin.H{"supply": supply.String()}
}

func (rpcProxy *RpcProxy) GetSessions(ctx *gin.Context, offset *big.Int, limit uint8) (int, gin.H) {
	if ctx.Query("user") != "" {
		sessions, err := rpcProxy.sessionRouter.GetSessionsByUser(ctx, common.HexToAddress(ctx.Query("user")), offset, limit)
		if err != nil {
			return constants.HTTP_INTERNAL_SERVER_ERROR, gin.H{"error": err.Error()}
		}
		return constants.HTTP_STATUS_OK, gin.H{"sessions": rpcProxy.mapSessions(sessions)}
	} else if ctx.Query("provider") != "" {
		sessions, err := rpcProxy.sessionRouter.GetSessionsByProvider(ctx, common.HexToAddress(ctx.Query("provider")), offset, limit)
		if err != nil {
			return constants.HTTP_INTERNAL_SERVER_ERROR, gin.H{"error": err.Error()}
		}
		return constants.HTTP_STATUS_OK, gin.H{"sessions": rpcProxy.mapSessions(sessions)}
	} else {
		return constants.HTTP_STATUS_BAD_REQUEST, gin.H{"error": "user or provider is required"}
	}
}

func (rpcProxy *RpcProxy) mapSessions(sessions []sessionrouter.Session) []*structs.Session {
	result := make([]*structs.Session, len(sessions))
	for i, value := range sessions {
		result[i] = &structs.Session{
			Id:                      lib.BytesToString(value.Id[:]),
			Provider:                value.Provider,
			User:                    value.User,
			ModelAgentId:            lib.BytesToString(value.ModelAgentId[:]),
			BidID:                   lib.BytesToString(value.BidID[:]),
			Stake:                   value.Stake,
			PricePerSecond:          value.PricePerSecond,
			CloseoutReceipt:         hex.EncodeToString(value.CloseoutReceipt),
			CloseoutType:            value.CloseoutType,
			ProviderWithdrawnAmount: value.ProviderWithdrawnAmount,
			OpenedAt:                value.OpenedAt,
			EndsAt:                  value.EndsAt,
			ClosedAt:                value.ClosedAt,
		}
	}
	return result
}

func (rpcProxy *RpcProxy) GetTransactions(ctx *gin.Context) (int, gin.H) {
	page := ctx.Query("page")
	limit := ctx.Query("limit")
	if page == "" {
		page = "1"
	}

	if limit == "" {
		limit = "10"
	}

	transactOpt, err := rpcProxy.getTransactOpts(ctx, rpcProxy.privateKey)
	if err != nil {
		return constants.HTTP_INTERNAL_SERVER_ERROR, gin.H{"error": err.Error()}
	}
	address := transactOpt.From

	ethTrxs, err := rpcProxy.explorerClient.GetEthTransactions(address.String(), page, limit)
	if err != nil {
		return constants.HTTP_INTERNAL_SERVER_ERROR, gin.H{"error": err.Error()}
	}
	morTrxs, err := rpcProxy.explorerClient.GetTokenTransactions(address.String(), page, limit)
	if err != nil {
		return constants.HTTP_INTERNAL_SERVER_ERROR, gin.H{"error": err.Error()}
	}

	allTrxs := append(ethTrxs, morTrxs...)
	sort.Slice(allTrxs, func(i, j int) bool {
		blockNumber1, err := strconv.ParseInt(allTrxs[i].BlockNumber, 10, 0)
		if err != nil {
			return false
		}
		blockNumber2, err := strconv.ParseInt(allTrxs[j].BlockNumber, 10, 0)
		if err != nil {
			return false
		}

		return blockNumber1 > blockNumber2
	})

	return constants.HTTP_STATUS_OK, gin.H{"transactions": allTrxs}
}

func (rpcProxy *RpcProxy) getTransactOpts(ctx context.Context, privKey string) (*bind.TransactOpts, error) {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return nil, err
	}

	chainId, err := rpcProxy.rpcClient.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	transactOpts, err := bind.NewKeyedTransactorWithChainID(privateKey, chainId)
	if err != nil {
		return nil, err
	}

	// TODO: deal with likely gasPrice issue so our transaction processes before another pending nonce.
	if rpcProxy.legacyTx {
		gasPrice, err := rpcProxy.rpcClient.SuggestGasPrice(ctx)
		if err != nil {
			return nil, err
		}
		transactOpts.GasPrice = gasPrice
	}

	transactOpts.Value = big.NewInt(0)
	transactOpts.Context = ctx

	return transactOpts, nil
}

func (rpcProxy *RpcProxy) signTx(ctx context.Context, tx *types.Transaction, privKey string) (*types.Transaction, error) {
	privateKey, err := crypto.HexToECDSA(privKey)
	if err != nil {
		return nil, err
	}

	chainId, err := rpcProxy.rpcClient.ChainID(ctx)
	if err != nil {
		return nil, err
	}

	return types.SignTx(tx, types.NewEIP155Signer(chainId), privateKey)
}

func (rpcProxy *RpcProxy) getSendParams(ctx *gin.Context) (string, *big.Int, error) {
	var reqPayload SendRequest
	if err := ctx.ShouldBindJSON(&reqPayload); err != nil {
		return "", &big.Int{}, err
	}

	to := reqPayload.To
	amountStr := reqPayload.Amount

	if to == "0" {
		return "", &big.Int{}, errors.New("to is required")
	}

	if amountStr == "" {
		return "", &big.Int{}, errors.New("amount is required")
	}

	amount, ok := new(big.Int).SetString(amountStr, 10)
	if !ok {
		return "", &big.Int{}, errors.New("invalid amount" + amountStr)
	}

	return to, amount, nil
}
