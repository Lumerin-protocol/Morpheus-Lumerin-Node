package aiengine

import (
	"context"
	"errors"

	"fmt"

	gcs "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/chatstorage/genericchatstorage"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/config"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/ethereum/go-ethereum/common"
)

type AiEngine struct {
	modelsConfigLoader *config.ModelConfigLoader
	service            ProxyService
	storage            gcs.ChatStorageInterface
	log                lib.ILogger
}

var (
	ErrChatCompletion                = errors.New("chat completion error")
	ErrImageGenerationInvalidRequest = errors.New("invalid prodia image generation request")
	ErrImageGenerationRequest        = errors.New("image generation error")
	ErrJobCheckRequest               = errors.New("job status check error")
	ErrJobFailed                     = errors.New("job failed")
)

func NewAiEngine(service ProxyService, storage gcs.ChatStorageInterface, modelsConfigLoader *config.ModelConfigLoader, log lib.ILogger) *AiEngine {
	return &AiEngine{
		modelsConfigLoader: modelsConfigLoader,
		service:            service,
		storage:            storage,
		log:                log,
	}
}

func (a *AiEngine) GetAdapter(ctx context.Context, chatID, modelID, sessionID common.Hash, storeChatContext, forwardChatContext bool) (AIEngineStream, error) {
	var engine AIEngineStream
	if sessionID == (common.Hash{}) {
		// local model
		modelConfig := a.modelsConfigLoader.ModelConfigFromID(modelID.Hex())
		if modelConfig == nil {
			return nil, fmt.Errorf("model not found: %s", modelID.Hex())
		}
		var ok bool
		engine, ok = ApiAdapterFactory(modelConfig.ApiType, modelConfig.ModelName, modelConfig.ApiURL, modelConfig.ApiKey, a.log)
		if !ok {
			return nil, fmt.Errorf("api adapter not found: %s", modelConfig.ApiType)
		}
	} else {
		// remote model
		engine = &RemoteModel{sessionID: sessionID, service: a.service}
	}

	if storeChatContext {
		engine = NewHistory(engine, a.storage, chatID, modelID, forwardChatContext, a.log)
	}

	return engine, nil
}

func (a *AiEngine) GetModelsConfig() ([]string, []config.ModelConfig) {
	return a.modelsConfigLoader.GetAll()
}

func (a *AiEngine) GetLocalModels() ([]LocalModel, error) {
	models := []LocalModel{}

	IDs, modelsFromConfig := a.modelsConfigLoader.GetAll()
	for i, model := range modelsFromConfig {
		models = append(models, LocalModel{
			Id:             IDs[i],
			Name:           model.ModelName,
			Model:          model.ModelName,
			ApiType:        model.ApiType,
			ApiUrl:         model.ApiURL,
			Slots:          model.ConcurrentSlots,
			CapacityPolicy: model.CapacityPolicy,
		})
	}

	return models, nil
}
