package lib

import (
	"math/big"
	"net/http/httptest"
	"testing"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/handlers/httphandlers"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
)


func PrepSessionForChat(t *testing.T) (server *httptest.Server, providerPubKey string, sessionId string, stake *big.Int, shouldReturn bool) {
	apiBus := InitializeApiBus(t)

	walletAddr, err := lib.PrivKeyStringToAddr(WALLET_PRIVATE_KEY)
	if err != nil {
		t.Fatalf("failed to get wallet address: %s", err)
		return nil, "", "", nil, true
	}

	handler := httphandlers.NewHTTPHandler(apiBus)

	server = httptest.NewServer(handler)
	defer server.Close()

	supply := GetTokenSupply(t, server)

	budget := GetTodaysBudget(t, server)

	bid := FindBid(t, server)

	pricePerSecond := new(big.Float).SetFloat64(bid["PricePerSecond"].(float64))
	pricePerSecondInt := new(big.Int)
	pricePerSecondInt, _ = pricePerSecondInt.SetString(pricePerSecond.Text('f', 0), 10)

	totalCost := SESSION_DURATION.Mul(pricePerSecondInt, SESSION_DURATION)
	stake = totalCost.Div(totalCost.Mul(supply, totalCost), budget)

	initiateSessionResponse := InitiateSession(t, server, walletAddr, stake)
	approval := initiateSessionResponse.Response.Result.Approval
	approvalSig := initiateSessionResponse.Response.Result.ApprovalSig
	providerPubKey = initiateSessionResponse.Response.Result.Message

	sessionId = OpenSession(t, server, approval, approvalSig, stake)
	return server, providerPubKey, sessionId, stake, shouldReturn
}
