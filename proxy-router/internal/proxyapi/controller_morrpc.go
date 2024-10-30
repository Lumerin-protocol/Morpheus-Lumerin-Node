package proxyapi

import (
	"context"
	"encoding/json"
	"fmt"
	"time"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	m "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/proxyapi/morrpcmessage"
	msg "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/proxyapi/morrpcmessage"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/storages"
	"github.com/go-playground/validator/v10"
)

type MORRPCController struct {
	service        *ProxyReceiver
	validator      *validator.Validate
	sessionStorage *storages.SessionStorage
	morRpc         *m.MORRPCMessage
}

type SendResponse func(*msg.RpcResponse) error

var (
	ErrUnknownMethod = fmt.Errorf("unknown method")
)

func NewMORRPCController(service *ProxyReceiver, validator *validator.Validate, sessionStorage *storages.SessionStorage) *MORRPCController {
	c := &MORRPCController{
		service:        service,
		validator:      validator,
		sessionStorage: sessionStorage,
		morRpc:         m.NewMorRpc(),
	}

	return c
}

func (s *MORRPCController) Handle(ctx context.Context, msg m.RPCMessage, sourceLog lib.ILogger, sendResponse SendResponse) error {
	switch msg.Method {
	case "session.request":
		return s.sessionRequest(ctx, msg, sendResponse, sourceLog)
	case "session.prompt":
		return s.sessionPrompt(ctx, msg, sendResponse, sourceLog)
	case "session.report":
		return s.sessionReport(ctx, msg, sendResponse, sourceLog)
	default:
		return lib.WrapError(ErrUnknownMethod, fmt.Errorf("unknown method: %s", msg.Method))
	}
}

var (
	ErrValidation     = fmt.Errorf("request validation failed")
	ErrUnmarshal      = fmt.Errorf("failed to unmarshal request")
	ErrGenerateReport = fmt.Errorf("failed to generate report")
)

func (s *MORRPCController) sessionRequest(ctx context.Context, msg m.RPCMessage, sendResponse SendResponse, sourceLog lib.ILogger) error {
	var req m.SessionReq
	err := json.Unmarshal(msg.Params, &req)
	if err != nil {
		return lib.WrapError(ErrUnmarshal, err)
	}

	if err := s.validator.Struct(req); err != nil {
		return lib.WrapError(ErrValidation, err)
	}
	sig := req.Signature
	req.Signature = lib.HexString{}
	isValid := s.morRpc.VerifySignature(req, sig, req.Key, sourceLog)
	if !isValid {
		err := ErrInvalidSig
		sourceLog.Error(err)
		return err
	}

	res, err := s.service.SessionRequest(ctx, msg.ID, msg.ID, &req, sourceLog)
	if err != nil {
		sourceLog.Error(err)
		return err
	}

	return sendResponse(res)
}

func (s *MORRPCController) sessionPrompt(ctx context.Context, msg m.RPCMessage, sendResponse SendResponse, sourceLog lib.ILogger) error {
	var req m.SessionPromptReq
	err := json.Unmarshal(msg.Params, &req)
	if err != nil {
		err := lib.WrapError(ErrUnmarshal, err)
		sourceLog.Error(err)
		return err
	}

	if err := s.validator.Struct(req); err != nil {
		err := lib.WrapError(ErrValidation, err)
		sourceLog.Error(err)
		return err
	}

	sourceLog.Debugf("Received prompt from session %s, timestamp: %s", req.SessionID, req.Timestamp)
	session, ok := s.sessionStorage.GetSession(req.SessionID.Hex())
	if !ok {
		err := fmt.Errorf("session not found")
		sourceLog.Error(err)
		return err
	}

	isSessionExpired := session.EndsAt.Uint64()*1000 < req.Timestamp
	if isSessionExpired {
		err := fmt.Errorf("session expired")
		sourceLog.Error(err)
		return err
	}

	user, ok := s.sessionStorage.GetUser(session.UserAddr)
	if !ok {
		err := fmt.Errorf("user not found")
		sourceLog.Error(err)
		return err
	}
	pubKeyHex, err := lib.StringToHexString(user.PubKey)
	if err != nil {
		sourceLog.Error(err)
		return err
	}

	sig := req.Signature
	req.Signature = lib.HexString{}

	isValid := s.morRpc.VerifySignature(req, sig, pubKeyHex, sourceLog)
	if !isValid {
		err := fmt.Errorf("invalid signature")
		sourceLog.Error(err)
		return err
	}

	now := time.Now().Unix()
	ttftMs, totalTokens, err := s.service.SessionPrompt(ctx, msg.ID, user.PubKey, &req, sendResponse, sourceLog)
	if err != nil {
		sourceLog.Error(err)
		return err
	}

	requestDuration := int(time.Now().Unix() - now)
	if requestDuration == 0 {
		requestDuration = 1
	}
	session.TTFTMsArr = append(session.TTFTMsArr, ttftMs)
	session.TPSScaled1000Arr = append(session.TPSScaled1000Arr, totalTokens*1000/requestDuration)
	err = s.sessionStorage.AddSession(session)
	if err != nil {
		sourceLog.Error(err)
		return err
	}
	return err
}

func (s *MORRPCController) sessionReport(ctx context.Context, msg m.RPCMessage, sendResponse SendResponse, sourceLog lib.ILogger) error {
	var req m.SessionReportReq
	err := json.Unmarshal(msg.Params, &req)
	if err != nil {
		err := lib.WrapError(ErrUnmarshal, err)
		sourceLog.Error(err)
		return err
	}

	if err := s.validator.Struct(req); err != nil {
		err := lib.WrapError(ErrValidation, err)
		sourceLog.Error(err)
		return err
	}

	sessionID := req.Message
	sourceLog.Debugf("Requested report from session %s, timestamp: %s", sessionID, req.Timestamp)
	session, ok := s.sessionStorage.GetSession(sessionID)
	if !ok {
		err := fmt.Errorf("session not found")
		sourceLog.Error(err)
		return err
	}

	user, ok := s.sessionStorage.GetUser(session.UserAddr)
	if !ok {
		err := fmt.Errorf("user not found")
		sourceLog.Error(err)
		return err
	}
	pubKeyHex, err := lib.StringToHexString(user.PubKey)
	if err != nil {
		sourceLog.Error(err)
		return err
	}

	sig := req.Signature
	req.Signature = lib.HexString{}

	isValid := s.morRpc.VerifySignature(req, sig, pubKeyHex, sourceLog)
	if !isValid {
		err := fmt.Errorf("invalid signature")
		sourceLog.Error(err)
		return err
	}

	res, err := s.service.SessionReport(ctx, msg.ID, msg.ID, session, sourceLog)
	if err != nil {
		sourceLog.Error(err)
		return ErrGenerateReport
	}

	return sendResponse(res)
}
