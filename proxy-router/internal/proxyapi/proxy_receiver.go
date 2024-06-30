package proxyapi

import (
	"context"
	"encoding/json"
	"fmt"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/aiengine"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	m "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/proxyapi/morrpcmessage"
	msg "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/proxyapi/morrpcmessage"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/storages"
	"github.com/sashabaranov/go-openai"
)

type ProxyReceiver struct {
	privateKeyHex  lib.HexString
	publicKeyHex   lib.HexString
	morRpc         *m.MORRPCMessage
	sessionStorage *storages.SessionStorage
	aiEngine       *aiengine.AiEngine
}

func NewProxyReceiver(privateKeyHex, publicKeyHex lib.HexString, sessionStorage *storages.SessionStorage, aiEngine *aiengine.AiEngine) *ProxyReceiver {
	return &ProxyReceiver{
		privateKeyHex:  privateKeyHex,
		publicKeyHex:   publicKeyHex,
		morRpc:         m.NewMorRpc(),
		sessionStorage: sessionStorage,
		aiEngine:       aiEngine,
	}
}

func (s *ProxyReceiver) SessionPrompt(ctx context.Context, requestID string, userPubKey string, rq *m.SessionPromptReq, sendResponse SendResponse, sourceLog lib.ILogger) error {
	// TODO: check if session is still active

	var req *openai.ChatCompletionRequest

	err := json.Unmarshal([]byte(rq.Message), &req)
	if err != nil {
		err := lib.WrapError(fmt.Errorf("failed to unmarshal prompt"), err)
		sourceLog.Error(err)
		return err
	}

	_, err = s.aiEngine.PromptStream(ctx, req, func(response *openai.ChatCompletionStreamResponse) error {
		marshalledResponse, err := json.Marshal(response)
		if err != nil {
			return err
		}

		encryptedResponse, err := lib.EncryptString(string(marshalledResponse), userPubKey)
		if err != nil {
			return err
		}

		// Send response
		r, err := s.morRpc.SessionPromptResponse(
			encryptedResponse,
			s.privateKeyHex,
			requestID,
		)
		if err != nil {
			err := lib.WrapError(fmt.Errorf("failed to create response"), err)
			sourceLog.Error(err)
			return err
		}
		return sendResponse(r)
	})

	if err != nil {
		err := lib.WrapError(fmt.Errorf("failed to prompt"), err)
		sourceLog.Error(err)
		return err
	}
	return nil
}

func (s *ProxyReceiver) SessionRequest(ctx context.Context, msgID string, reqID string, req *m.SessionReq, sourceLog lib.ILogger) (*msg.RpcResponse, error) {
	sourceLog.Debugf("Received session request from %s, timestamp: %s", req.User, req.Timestamp)

	hasCapacity := true // check if there is capacity
	if !hasCapacity {
		err := fmt.Errorf("no capacity")
		sourceLog.Error(err)
		return nil, err
	}

	// Send response
	response, err := s.morRpc.InitiateSessionResponse(
		s.publicKeyHex,
		req.User,
		req.BidID,
		s.privateKeyHex,
		reqID,
	)
	if err != nil {
		err := lib.WrapError(fmt.Errorf("failed to create response"), err)
		sourceLog.Error(err)
		return nil, err
	}

	user := storages.User{
		Addr:   req.User.Hex(),
		PubKey: req.Key.Hex(),
	}

	err = s.sessionStorage.AddUser(&user)
	if err != nil {
		err := lib.WrapError(fmt.Errorf("failed store user"), err)
		sourceLog.Error(err)
		return nil, err
	}
	return response, nil
}
