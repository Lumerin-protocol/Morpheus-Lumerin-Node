package proxyapi

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"math/big"
	"net"
	"net/http"
	"net/url"
	"strconv"
	"time"

	gcs "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/chatstorage/genericchatstorage"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/interfaces"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	msgs "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/proxyapi/morrpcmessage"
	sessionrepo "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/session"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/storages"
	"github.com/ethereum/go-ethereum/common"
	"github.com/gin-gonic/gin/binding"
	"github.com/sashabaranov/go-openai"
)

var (
	ErrMissingPrKey     = fmt.Errorf("missing private key")
	ErrCreateReq        = fmt.Errorf("failed to create request")
	ErrProvider         = fmt.Errorf("provider request failed")
	ErrInvalidSig       = fmt.Errorf("received invalid signature from provider")
	ErrFailedStore      = fmt.Errorf("failed store user")
	ErrInvalidResponse  = fmt.Errorf("invalid response")
	ErrResponseErr      = fmt.Errorf("response error")
	ErrDecrFailed       = fmt.Errorf("failed to decrypt ai response chunk")
	ErrMasrshalFailed   = fmt.Errorf("failed to marshal response")
	ErrDecode           = fmt.Errorf("failed to decode response")
	ErrSessionNotFound  = fmt.Errorf("session not found")
	ErrSessionExpired   = fmt.Errorf("session expired")
	ErrProviderNotFound = fmt.Errorf("provider not found")
	ErrEmpty            = fmt.Errorf("empty result and no error")
	ErrConnectProvider  = fmt.Errorf("failed to connect to provider")
	ErrWriteProvider    = fmt.Errorf("failed to write to provider")
)

type ProxyServiceSender struct {
	chainID        *big.Int
	publicUrl      *url.URL
	privateKey     interfaces.PrKeyProvider
	logStorage     *lib.Collection[*interfaces.LogStorage]
	sessionStorage *storages.SessionStorage
	sessionRepo    *sessionrepo.SessionRepositoryCached
	morRPC         *msgs.MORRPCMessage
	sessionService SessionService
	log            lib.ILogger
}

func NewProxySender(chainID *big.Int, publicUrl *url.URL, privateKey interfaces.PrKeyProvider, logStorage *lib.Collection[*interfaces.LogStorage], sessionStorage *storages.SessionStorage, sessionRepo *sessionrepo.SessionRepositoryCached, log lib.ILogger) *ProxyServiceSender {
	return &ProxyServiceSender{
		chainID:        chainID,
		publicUrl:      publicUrl,
		privateKey:     privateKey,
		logStorage:     logStorage,
		sessionStorage: sessionStorage,
		sessionRepo:    sessionRepo,
		morRPC:         msgs.NewMorRpc(),
		log:            log,
	}
}

func (p *ProxyServiceSender) SetSessionService(service SessionService) {
	p.sessionService = service
}

func (p *ProxyServiceSender) InitiateSession(ctx context.Context, user common.Address, provider common.Address, spend *big.Int, bidID common.Hash, providerURL string) (*msgs.SessionRes, error) {
	requestID := "1"

	prKey, err := p.privateKey.GetPrivateKey()
	if err != nil {
		return nil, ErrMissingPrKey
	}

	initiateSessionRequest, err := p.morRPC.InitiateSessionRequest(user, provider, spend, bidID, prKey, requestID)
	if err != nil {
		return nil, lib.WrapError(ErrCreateReq, err)
	}

	msg, code, err := p.rpcRequest(providerURL, initiateSessionRequest)
	if err != nil {
		return nil, lib.WrapError(ErrProvider, fmt.Errorf("code: %d, msg: %v, error: %s", code, msg, err))
	}

	if msg.Error != nil {
		// TODO: verify signature
		return nil, lib.WrapError(ErrResponseErr, fmt.Errorf("error: %v, result: %v", msg.Error.Message, msg.Error.Data))
	}
	if msg.Result == nil {
		return nil, lib.WrapError(ErrInvalidResponse, ErrEmpty)
	}

	var typedMsg *msgs.SessionRes
	err = json.Unmarshal(*msg.Result, &typedMsg)
	if err != nil {
		return nil, lib.WrapError(ErrInvalidResponse, fmt.Errorf("expected InitiateSessionResponse, got %s", msg.Result))
	}

	err = binding.Validator.ValidateStruct(typedMsg)
	if err != nil {
		return nil, lib.WrapError(ErrInvalidResponse, err)
	}

	signature := typedMsg.Signature
	typedMsg.Signature = lib.HexString{}

	providerPubKey := typedMsg.PubKey
	if !p.validateMsgSignature(typedMsg, signature, typedMsg.PubKey) {
		return nil, ErrInvalidSig
	}

	err = p.sessionStorage.AddUser(&storages.User{
		Addr:   provider.Hex(),
		PubKey: providerPubKey.String(),
		Url:    providerURL,
	})
	if err != nil {
		return nil, lib.WrapError(ErrFailedStore, err)
	}

	return typedMsg, nil
}

func (p *ProxyServiceSender) GetSessionReportFromProvider(ctx context.Context, sessionID common.Hash) (*msgs.SessionReportRes, error) {
	requestID := "1"

	prKey, err := p.privateKey.GetPrivateKey()
	if err != nil {
		return nil, ErrMissingPrKey
	}

	session, err := p.sessionRepo.GetSession(ctx, sessionID)
	if err != nil {
		return nil, ErrSessionNotFound
	}
	provider, ok := p.sessionStorage.GetUser(session.ProviderAddr().Hex())
	if !ok {
		return nil, ErrProviderNotFound
	}

	getSessionReportRequest, err := p.morRPC.SessionReportRequest(sessionID, prKey, requestID)
	if err != nil {
		return nil, lib.WrapError(ErrCreateReq, err)
	}

	msg, code, err := p.rpcRequest(provider.Url, getSessionReportRequest)
	if err != nil {
		return nil, lib.WrapError(ErrProvider, fmt.Errorf("code: %d, msg: %v, error: %s", code, msg, err))
	}

	if msg.Error != nil {
		// TODO: verify signature
		return nil, lib.WrapError(ErrResponseErr, fmt.Errorf("error: %v, result: %v", msg.Error.Message, msg.Error.Data))
	}
	if msg.Result == nil {
		return nil, lib.WrapError(ErrInvalidResponse, ErrEmpty)
	}

	var typedMsg *msgs.SessionReportRes
	err = json.Unmarshal(*msg.Result, &typedMsg)
	if err != nil {
		return nil, lib.WrapError(ErrInvalidResponse, fmt.Errorf("expected SessionReportRespose, got %s", msg.Result))
	}

	err = binding.Validator.ValidateStruct(typedMsg)
	if err != nil {
		return nil, lib.WrapError(ErrInvalidResponse, err)
	}

	signature := typedMsg.Signature
	typedMsg.Signature = lib.HexString{}

	hexPubKey, err := lib.StringToHexString(provider.PubKey)
	if err != nil {
		return nil, lib.WrapError(ErrInvalidResponse, err)
	}

	if !p.validateMsgSignature(typedMsg, signature, hexPubKey) {
		return nil, ErrInvalidSig
	}

	return typedMsg, nil
}

func (p *ProxyServiceSender) GetSessionReportFromUser(ctx context.Context, sessionID common.Hash) (lib.HexString, lib.HexString, error) {
	session, err := p.sessionRepo.GetSession(ctx, sessionID)
	if err != nil {
		return nil, nil, ErrSessionNotFound
	}

	TPSScaled1000Arr, TTFTMsArr := session.GetStats()

	tps := 0
	ttft := 0
	for _, tpsVal := range TPSScaled1000Arr {
		tps += tpsVal
	}
	for _, ttftVal := range TTFTMsArr {
		ttft += ttftVal
	}

	if len(TPSScaled1000Arr) != 0 {
		tps /= len(TPSScaled1000Arr)
	}
	if len(TTFTMsArr) != 0 {
		ttft /= len(TTFTMsArr)
	}

	prKey, err := p.privateKey.GetPrivateKey()
	if err != nil {
		return nil, nil, ErrMissingPrKey
	}

	response, err := p.morRPC.SessionReportResponse(
		uint32(tps),
		uint32(ttft),
		sessionID,
		prKey,
		"1",
		p.chainID,
	)

	if err != nil {
		return nil, nil, lib.WrapError(ErrGenerateReport, err)
	}

	var typedMsg *msgs.SessionReportRes
	err = json.Unmarshal(*response.Result, &typedMsg)
	if err != nil {
		return nil, nil, lib.WrapError(ErrInvalidResponse, fmt.Errorf("expected SessionReportRespose, got %s", response.Result))
	}

	return typedMsg.Message, typedMsg.SignedReport, nil
}

func (p *ProxyServiceSender) rpcRequest(url string, rpcMessage *msgs.RPCMessage) (*msgs.RpcResponse, int, error) {
	TIMEOUT_TO_ESTABLISH_CONNECTION := time.Second * 3
	dialer := net.Dialer{Timeout: TIMEOUT_TO_ESTABLISH_CONNECTION}

	conn, err := dialer.Dial("tcp", url)
	if err != nil {
		err = lib.WrapError(ErrConnectProvider, err)
		p.log.Errorf("%s", err)
		return nil, http.StatusInternalServerError, err
	}
	defer conn.Close()

	msgJSON, err := json.Marshal(rpcMessage)
	if err != nil {
		err = lib.WrapError(ErrMasrshalFailed, err)
		p.log.Errorf("%s", err)
		return nil, http.StatusInternalServerError, err
	}
	_, err = conn.Write(msgJSON)
	if err != nil {
		err = lib.WrapError(ErrWriteProvider, err)
		p.log.Errorf("%s", err)
		return nil, http.StatusInternalServerError, err
	}

	// read response
	reader := bufio.NewReader(conn)
	d := json.NewDecoder(reader)

	var msg *msgs.RpcResponse
	err = d.Decode(&msg)
	if err != nil {
		err = lib.WrapError(ErrDecode, err)
		p.log.Errorf("%s", err)
		return nil, http.StatusBadRequest, err
	}
	return msg, 0, nil
}

func (p *ProxyServiceSender) validateMsgSignature(result any, signature lib.HexString, providerPubicKey lib.HexString) bool {
	return p.morRPC.VerifySignature(result, signature, providerPubicKey, p.log)
}

func (p *ProxyServiceSender) SendPromptV2(ctx context.Context, sessionID common.Hash, prompt *openai.ChatCompletionRequest, cb gcs.CompletionCallback) (interface{}, error) {
	session, err := p.sessionRepo.GetSession(ctx, sessionID)
	if err != nil {
		return nil, ErrSessionNotFound
	}

	isExpired := session.EndsAt().Int64()-time.Now().Unix() < 0
	if isExpired {
		return nil, ErrSessionExpired
	}

	provider, ok := p.sessionStorage.GetUser(session.ProviderAddr().Hex())
	if !ok {
		return nil, ErrProviderNotFound
	}

	prKey, err := p.privateKey.GetPrivateKey()
	if err != nil {
		return nil, ErrMissingPrKey
	}

	requestID := "1"
	pubKey, err := lib.StringToHexString(provider.PubKey)
	if err != nil {
		return nil, lib.WrapError(ErrCreateReq, err)
	}
	promptRequest, err := p.morRPC.SessionPromptRequest(sessionID, prompt, pubKey, prKey, requestID)
	if err != nil {
		return nil, lib.WrapError(ErrCreateReq, err)
	}

	now := time.Now().Unix()
	result, ttftMs, totalTokens, err := p.rpcRequestStreamV2(ctx, cb, provider.Url, promptRequest, pubKey)
	if err != nil {
		if !session.FailoverEnabled() {
			return nil, lib.WrapError(ErrProvider, err)
		}

		_, err := p.sessionService.CloseSession(ctx, sessionID)
		if err != nil {
			return nil, err
		}

		err = cb(ctx, gcs.NewChunkControl("provider failed, failover enabled"))
		if err != nil {
			return nil, err
		}

		duration := session.EndsAt().Int64() - time.Now().Unix()

		newSessionID, err := p.sessionService.OpenSessionByModelId(
			ctx,
			session.ModelID(),
			big.NewInt(duration),
			session.DirectPayment(),
			session.FailoverEnabled(),
			session.ProviderAddr(),
		)
		if err != nil {
			return nil, err
		}

		msg := fmt.Sprintf("new session opened: %s", newSessionID.Hex())
		err = cb(ctx, gcs.NewChunkControl(msg))
		if err != nil {
			return nil, err
		}

		return p.SendPromptV2(ctx, newSessionID, prompt, cb)
	}

	requestDuration := int(time.Now().Unix() - now)
	if requestDuration == 0 {
		requestDuration = 1
	}
	session.AddStats(totalTokens*1000/requestDuration, ttftMs)

	err = p.sessionRepo.SaveSession(ctx, session)
	if err != nil {
		p.log.Error(`failed to update session report stats`, err)
	}

	return result, nil
}
func (p *ProxyServiceSender) rpcRequestStreamV2(
	ctx context.Context,
	cb gcs.CompletionCallback,
	url string,
	rpcMessage *msgs.RPCMessage,
	providerPublicKey lib.HexString,
) (interface{}, int, int, error) {
	const (
		TIMEOUT_TO_ESTABLISH_CONNECTION   = time.Second * 3
		TIMEOUT_TO_RECEIVE_FIRST_RESPONSE = time.Second * 5
		MAX_RETRIES                       = 5
	)

	dialer := net.Dialer{Timeout: TIMEOUT_TO_ESTABLISH_CONNECTION}

	prKey, err := p.privateKey.GetPrivateKey()
	if err != nil {
		return nil, 0, 0, ErrMissingPrKey
	}

	conn, err := dialer.Dial("tcp", url)
	if err != nil {
		err = lib.WrapError(ErrConnectProvider, err)
		p.log.Errorf("%s", err)
		return nil, 0, 0, err
	}
	defer conn.Close()

	// Set initial read deadline
	_ = conn.SetReadDeadline(time.Now().Add(TIMEOUT_TO_RECEIVE_FIRST_RESPONSE))

	msgJSON, err := json.Marshal(rpcMessage)
	if err != nil {
		return nil, 0, 0, lib.WrapError(ErrMasrshalFailed, err)
	}

	ttftMs := 0
	totalTokens := 0
	now := time.Now().UnixMilli()

	_, err = conn.Write(msgJSON)
	if err != nil {
		return nil, ttftMs, totalTokens, err
	}

	reader := bufio.NewReader(conn)
	// We need to recreate the decoder if it becomes invalid
	var d *json.Decoder

	responses := make([]interface{}, 0)

	retryCount := 0

	for {
		if ctx.Err() != nil {
			return nil, ttftMs, totalTokens, ctx.Err()
		}

		// Initialize or reset the decoder
		if d == nil {
			d = json.NewDecoder(reader)
		}

		var msg *msgs.RpcResponse
		err = d.Decode(&msg)
		if err != nil {
			if netErr, ok := err.(net.Error); ok && netErr.Timeout() {
				p.log.Warnf("Read operation timed out: %v", err)
				if retryCount < MAX_RETRIES {
					alive, availErr := checkProviderAvailability(url)
					if availErr != nil {
						p.log.Warnf("Provider availability check failed: %v", availErr)
						return nil, ttftMs, totalTokens, fmt.Errorf("provider availability check failed: %w", availErr)
					}
					if alive {
						retryCount++
						p.log.Infof("Provider is alive, retrying (%d/%d)...", retryCount, MAX_RETRIES)
						// Reset the read deadline
						conn.SetReadDeadline(time.Now().Add(TIMEOUT_TO_RECEIVE_FIRST_RESPONSE))
						// Clear the error state by reading any remaining data
						reader.Discard(reader.Buffered())
						// Reset the decoder
						d = nil
						continue
					} else {
						return nil, ttftMs, totalTokens, fmt.Errorf("provider is not available")
					}
				} else {
					return nil, ttftMs, totalTokens, fmt.Errorf("read timed out after %d retries: %w", retryCount, err)
				}
			} else if err == io.EOF {
				p.log.Warnf("Connection closed by provider")
				return nil, ttftMs, totalTokens, fmt.Errorf("connection closed by provider")
			} else {
				p.log.Warnf("Failed to decode response: %v", err)
				return nil, ttftMs, totalTokens, lib.WrapError(ErrInvalidResponse, err)
			}
		}

		if msg.Error != nil {
			return nil, ttftMs, totalTokens, lib.WrapError(ErrResponseErr, fmt.Errorf("error: %v, data: %v", msg.Error.Message, msg.Error.Data))
		}

		if msg.Result == nil {
			return nil, ttftMs, totalTokens, lib.WrapError(ErrInvalidResponse, ErrEmpty)
		}

		if ttftMs == 0 {
			ttftMs = int(time.Now().UnixMilli() - now)
			_ = conn.SetReadDeadline(time.Time{}) // Clear read deadline
		}

		var inferenceRes InferenceRes
		err = json.Unmarshal(*msg.Result, &inferenceRes)
		if err != nil {
			return nil, ttftMs, totalTokens, lib.WrapError(ErrInvalidResponse, err)
		}
		sig := inferenceRes.Signature
		inferenceRes.Signature = []byte{}

		if !p.validateMsgSignature(inferenceRes, sig, providerPublicKey) {
			return nil, ttftMs, totalTokens, ErrInvalidSig
		}

		var message lib.HexString
		err = json.Unmarshal(inferenceRes.Message, &message)
		if err != nil {
			return nil, ttftMs, totalTokens, lib.WrapError(ErrInvalidResponse, err)
		}

		aiResponse, err := lib.DecryptBytes(message, prKey)
		if err != nil {
			return nil, ttftMs, totalTokens, lib.WrapError(ErrDecrFailed, err)
		}

		var payload openai.ChatCompletionStreamResponse
		err = json.Unmarshal(aiResponse, &payload)
		var stop = true
		var chunk gcs.Chunk
		if err == nil && len(payload.Choices) > 0 {
			stop = false
			choices := payload.Choices
			for _, choice := range choices {
				if choice.FinishReason == openai.FinishReasonStop {
					stop = true
				}
			}
			totalTokens += len(choices)
			responses = append(responses, payload)
			chunk = gcs.NewChunkStreaming(&payload)
		} else {
			var imageGenerationResult gcs.ImageGenerationResult
			err = json.Unmarshal(aiResponse, &imageGenerationResult)
			if err != nil {
				return nil, ttftMs, totalTokens, lib.WrapError(ErrInvalidResponse, err)
			}
			totalTokens += 1
			responses = append(responses, imageGenerationResult)
			chunk = gcs.NewChunkImage(&imageGenerationResult)
		}

		if ctx.Err() != nil {
			return nil, ttftMs, totalTokens, ctx.Err()
		}
		err = cb(ctx, chunk)
		if err != nil {
			return nil, ttftMs, totalTokens, lib.WrapError(ErrResponseErr, err)
		}
		if stop {
			break
		}
	}

	return responses, ttftMs, totalTokens, nil
}

// checkProviderAvailability checks if the provider is alive using portchecker.io API
func checkProviderAvailability(url string) (bool, error) {
	host, port, err := net.SplitHostPort(url)
	if err != nil {
		return false, err
	}

	portInt, err := strconv.Atoi(port)
	if err != nil {
		return false, err
	}

	requestBody, err := json.Marshal(map[string]interface{}{
		"host":  host,
		"ports": []int{portInt},
	})
	if err != nil {
		return false, err
	}

	req, err := http.NewRequest("POST", "https://portchecker.io/api/v1/query", bytes.NewBuffer(requestBody))
	if err != nil {
		return false, err
	}
	req.Header.Set("Content-Type", "application/json")

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return false, err
	}
	defer resp.Body.Close()

	body, err := io.ReadAll(resp.Body)
	if err != nil {
		return false, err
	}

	var response struct {
		Check []struct {
			Status bool `json:"status"`
			Port   int  `json:"port"`
		} `json:"check"`
	}
	err = json.Unmarshal(body, &response)
	if err != nil {
		return false, err
	}

	for _, check := range response.Check {
		if check.Port == portInt {
			return check.Status, nil
		}
	}

	return false, fmt.Errorf("port status not found in response")
}
