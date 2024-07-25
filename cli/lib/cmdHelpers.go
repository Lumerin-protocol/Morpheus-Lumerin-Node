package lib

import (
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"io/ioutil"
	"math/big"
	"net/http"
)

type SessionConfig struct {
	LocalProxyRouterUrl string `json:"localProxyRouterUrl"`
	DiamondAddress      string `json:"diamondAddress"`
	Address             string `json:"address"`
	Endpoint            string `json:"endpoint"`
	BidId               string `json:"bidId"`
	Provider            string `json:"provider"`
	Spend			   int    `json:"spend"`
}

type InitiateSessionBody struct {
	User        string `json:"user"`
	Provider    string `json:"provider"`
	Spend       int    `json:"spend"`
	BidId       string `json:"bidId"`
	ProviderUrl string `json:"providerUrl"`
}

type DataResponse struct {
	Response struct {
		Result struct {
			Approval    string `json:"approval"`
			ApprovalSig string `json:"approvalSig"`
		} `json:"result"`
	} `json:"response"`
	SessionId string `json:"sessionId"`
}

func log(level, message string) {
	// Implement your log logic here
	fmt.Printf("[%s] %s\n", level, message)
}

func initiateSession(props SessionConfig) (map[string]string, string, error) {
	fmt.Println("open-session with stake: ", props.Spend)
	log("info", "Processing...")
	signature := make(map[string]string)
	var sessionId string


	// Initiate session
	initiateSessionPath := fmt.Sprintf("%s/proxy/sessions/initiate", props.LocalProxyRouterUrl)
	initiateSessionBody := InitiateSessionBody{
		User:        props.Address,
		Provider:    props.Provider,
		Spend:       props.Spend,
		BidId:       props.BidId,
		ProviderUrl: props.Endpoint,
	}
	initiateSessionJSON, _ := json.Marshal(initiateSessionBody)

	resp, err := http.Post(initiateSessionPath, "application/json", bytes.NewBuffer(initiateSessionJSON))
	if err != nil {
		log("error", "Failed to initiate session")
		fmt.Println("Failed initiate session", err)
		return nil, "", err
	}
	defer resp.Body.Close()

	respBody, _ := io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		log("error", "Failed to initiate session")
		fmt.Println("Failed initiate session", string(respBody))
		return nil, "", fmt.Errorf("failed to initiate session: %s", string(respBody))
	}

	var dataResponse DataResponse
	err = json.Unmarshal(respBody, &dataResponse)
	if err != nil {
		log("error", "Failed to initiate session")
		fmt.Println("Failed initiate session", err)
		return nil, "", err
	}

	signature["approval"] = dataResponse.Response.Result.Approval
	signature["approvalSig"] = dataResponse.Response.Result.ApprovalSig

	// Approve blockchain transaction
	approvePath := fmt.Sprintf("%s/blockchain/approve?amount=%s&spender=%s",
		props.LocalProxyRouterUrl,
		big.NewInt(int64(props.Spend)).String(),
		props.DiamondAddress)

	resp, err = http.Post(approvePath, "application/json", nil)
	if err != nil {
		log("error", "Failed to increase allowance")
		fmt.Println("Failed to increase allowance", err)
		return nil, "", err
	}
	defer resp.Body.Close()

	respBody, _ = io.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		log("error", "Failed to increase allowance")
		fmt.Println("Failed to increase allowance", string(respBody))
		return nil, "", fmt.Errorf("failed to increase allowance: %s", string(respBody))
	}

	// Set blockchain session
	setSessionPath := fmt.Sprintf("%s/blockchain/sessions", props.LocalProxyRouterUrl)
	setSessionBody := map[string]string{
		"approval":    signature["approval"],
		"approvalSig": signature["approvalSig"],
		"stake":       big.NewInt(int64(props.Spend)).String(),
	}
	setSessionJSON, _ := json.Marshal(setSessionBody)

	resp, err = http.Post(setSessionPath, "application/json", bytes.NewBuffer(setSessionJSON))
	if err != nil {
		log("error", "Failed to set session")
		fmt.Println("Failed to set session", err)
		return nil, "", err
	}
	defer resp.Body.Close()

	respBody, _ = ioutil.ReadAll(resp.Body)
	if resp.StatusCode != http.StatusOK {
		log("error", "Failed to set session")
		fmt.Println("Failed to set session", string(respBody))
		return nil, "", fmt.Errorf("failed to set session: %s", string(respBody))
	}

	err = json.Unmarshal(respBody, &dataResponse)
	if err != nil {
		log("error", "Failed to set session")
		fmt.Println("Failed to set session", err)
		return nil, "", err
	}

	sessionId = dataResponse.SessionId
	return signature, sessionId, nil
}

func OpenSession(
	ctx context.Context,
	localProxyRouterUrl string,
	contractAddress string,
	userWalletAddress string,
	bidId string,
	provider string,
	providerEndpoint string,
	stake int,
) (map[string]string, string, error) {

	props := SessionConfig{
		LocalProxyRouterUrl: localProxyRouterUrl,
		DiamondAddress:      contractAddress,
		Address:             userWalletAddress,
		Endpoint:            providerEndpoint,
		BidId:               bidId,
		Provider:            provider,
		Spend: stake,
	}

	signature, sessionId, err := initiateSession(props)
	if err != nil {
		fmt.Println("Error:", err)
	} else {
		fmt.Println("Signature:", signature)
		fmt.Println("Session ID:", sessionId)
	}

	return signature, sessionId, err
}