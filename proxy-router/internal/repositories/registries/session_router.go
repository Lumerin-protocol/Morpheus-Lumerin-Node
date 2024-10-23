package registries

import (
	"context"
	"fmt"
	"math/big"
	"time"

	i "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/interfaces"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/contracts/bindings/sessionrouter"
	src "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/contracts/bindings/sessionrouter"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/accounts/abi/bind"
	"github.com/ethereum/go-ethereum/common"
)

type SessionRouter struct {
	// config
	sessionRouterAddr common.Address

	// state
	nonce uint64
	srABI *abi.ABI

	// deps
	sessionRouter *src.SessionRouter
	client        i.ContractBackend
	log           lib.ILogger
}

var closeReportAbi = []lib.AbiParameter{
	{Type: "bytes32"},
	{Type: "uint128"},
	{Type: "uint32"},
}

func NewSessionRouter(sessionRouterAddr common.Address, client i.ContractBackend, log lib.ILogger) *SessionRouter {
	sr, err := sessionrouter.NewSessionRouter(sessionRouterAddr, client)
	if err != nil {
		panic("invalid marketplace ABI")
	}
	srABI, err := src.SessionRouterMetaData.GetAbi()
	if err != nil {
		panic("invalid marketplace ABI: " + err.Error())
	}
	return &SessionRouter{
		sessionRouter:     sr,
		sessionRouterAddr: sessionRouterAddr,
		client:            client,
		srABI:             srABI,
		log:               log,
	}
}

func (g *SessionRouter) OpenSession(opts *bind.TransactOpts, approval []byte, approvalSig []byte, stake *big.Int, privateKeyHex lib.HexString) (sessionID common.Hash, providerID common.Address, userID common.Address, err error) {
	sessionTx, err := g.sessionRouter.OpenSession(opts, stake, approval, approvalSig)
	if err != nil {
		return common.Hash{}, common.Address{}, common.Address{}, lib.TryConvertGethError(err)
	}

	// Wait for the transaction receipt
	receipt, err := bind.WaitMined(opts.Context, g.client, sessionTx)
	if err != nil {
		return common.Hash{}, common.Address{}, common.Address{}, lib.TryConvertGethError(err)
	}

	// Find the event log
	for _, log := range receipt.Logs {
		// Check if the log belongs to the OpenSession event
		event, err := g.sessionRouter.ParseSessionOpened(*log)
		if err == nil {
			return event.SessionId, event.ProviderId, event.User, nil
		}
	}

	return common.Hash{}, common.Address{}, common.Address{}, fmt.Errorf("OpenSession event not found in transaction logs")
}

func (g *SessionRouter) GetSession(ctx context.Context, sessionID common.Hash) (*src.ISessionStorageSession, error) {
	session, err := g.sessionRouter.GetSession(&bind.CallOpts{Context: ctx}, sessionID)
	if err != nil {
		return nil, err
	}

	return &session, nil
}

func (g *SessionRouter) GetSessionsByProvider(ctx context.Context, providerAddr common.Address, offset *big.Int, limit uint8) ([][32]byte, []src.ISessionStorageSession, error) {
	sessionIDs, err := g.sessionRouter.GetProviderSessions(&bind.CallOpts{Context: ctx}, providerAddr, offset, big.NewInt(int64(limit)))
	if err != nil {
		return nil, nil, lib.TryConvertGethError(err)
	}
	return g.getMultipleSessions(ctx, sessionIDs)
}

func (g *SessionRouter) GetSessionsByUser(ctx context.Context, userAddr common.Address, offset *big.Int, limit uint8) ([][32]byte, []src.ISessionStorageSession, error) {
	IDs, err := g.sessionRouter.GetUserSessions(&bind.CallOpts{Context: ctx}, userAddr, offset, big.NewInt(int64(limit)))
	if err != nil {
		return nil, nil, lib.TryConvertGethError(err)
	}
	return g.getMultipleSessions(ctx, IDs)
}

func (g *SessionRouter) CloseSession(opts *bind.TransactOpts, sessionID common.Hash, report []byte, signedReport []byte, privateKeyHex lib.HexString) (common.Hash, error) {
	sessionTx, err := g.sessionRouter.CloseSession(opts, report, signedReport)
	if err != nil {
		return common.Hash{}, lib.TryConvertGethError(err)
	}

	// Wait for the transaction receipt
	receipt, err := bind.WaitMined(opts.Context, g.client, sessionTx)
	if err != nil {
		return common.Hash{}, err
	}

	if receipt.Status != 1 {
		return receipt.TxHash, fmt.Errorf("Transaction failed with status %d", receipt.Status)
	}

	return sessionTx.Hash(), nil
}

func (g *SessionRouter) GetProviderClaimableBalance(ctx context.Context, sessionId [32]byte) (*big.Int, error) {
	session, err := g.sessionRouter.GetSession(&bind.CallOpts{Context: ctx}, sessionId)
	if err != nil {
		return nil, lib.TryConvertGethError(err)
	}

	bid, err := g.sessionRouter.GetBid(&bind.CallOpts{Context: ctx}, session.BidId)
	if err != nil {
		return nil, lib.TryConvertGethError(err)
	}

	var sessionEnd *big.Int
	if session.ClosedAt.Cmp(big.NewInt(0)) == 0 {
		sessionEnd = session.EndsAt
	} else {
		sessionEnd = session.ClosedAt
	}

	if sessionEnd.Cmp(big.NewInt(time.Now().Unix())) > 0 {
		return nil, fmt.Errorf("session not ended or does not exist")
	}

	duration := new(big.Int).Sub(sessionEnd, session.OpenedAt)
	amount := new(big.Int).Mul(duration, bid.PricePerSecond)
	amount.Sub(amount, session.ProviderWithdrawnAmount)

	return amount, nil
}

func (g *SessionRouter) ClaimProviderBalance(opts *bind.TransactOpts, sessionId [32]byte) (common.Hash, error) {
	tx, err := g.sessionRouter.ClaimForProvider(opts, sessionId)
	if err != nil {
		return common.Hash{}, lib.TryConvertGethError(err)
	}

	// Wait for the transaction receipt
	_, err = bind.WaitMined(opts.Context, g.client, tx)
	if err != nil {
		return common.Hash{}, err
	}

	return tx.Hash(), nil
}

func (g *SessionRouter) GetTodaysBudget(ctx context.Context) (*big.Int, error) {
	timestamp := big.NewInt(time.Now().Unix())
	budget, err := g.sessionRouter.GetTodaysBudget(&bind.CallOpts{Context: ctx}, timestamp)
	if err != nil {
		return nil, lib.TryConvertGethError(err)
	}
	return budget, nil
}

func (g *SessionRouter) GetModelStats(ctx context.Context, modelID [32]byte) (interface{}, error) {
	res, err := g.sessionRouter.GetModelStats(&bind.CallOpts{Context: ctx}, modelID)
	if err != nil {
		return nil, lib.TryConvertGethError(err)
	}
	return res, nil
}

func (g *SessionRouter) GetProviderModelStats(ctx context.Context, modelID [32]byte, provider common.Address) (*src.IStatsStorageProviderModelStats, error) {
	res, err := g.sessionRouter.GetProviderModelStats(&bind.CallOpts{Context: ctx}, modelID, provider)
	if err != nil {
		return nil, lib.TryConvertGethError(err)
	}
	return &res, nil
}

func (g *SessionRouter) getMultipleSessions(ctx context.Context, IDs [][32]byte) ([][32]byte, []src.ISessionStorageSession, error) {
	// todo: replace with multicall
	sessions := make([]src.ISessionStorageSession, len(IDs))
	for i, id := range IDs {
		session, err := g.sessionRouter.GetSession(&bind.CallOpts{Context: ctx}, id)
		if err != nil {
			return nil, nil, err
		}
		sessions[i] = session
	}
	return IDs, sessions, nil
}

func (g *SessionRouter) GetContractAddress() common.Address {
	return g.sessionRouterAddr
}

func (g *SessionRouter) GetABI() *abi.ABI {
	return g.srABI
}
