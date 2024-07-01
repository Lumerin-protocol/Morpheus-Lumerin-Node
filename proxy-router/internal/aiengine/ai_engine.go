package aiengine

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"errors"
	"net/http"
	"strings"

	"fmt"

	c "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/gin-gonic/gin"
	"github.com/sashabaranov/go-openai"
	api "github.com/sashabaranov/go-openai"
)

type AiEngine struct {
	client  *api.Client
	baseURL string
	apiKey  string
	log     lib.ILogger
}

var (
	ErrChatCompletion = errors.New("chat completion error")
)

func NewAiEngine(apiBaseURL, apiKey string, log lib.ILogger) *AiEngine {
	return &AiEngine{
		baseURL: apiBaseURL,
		apiKey:  apiKey,
		client: api.NewClientWithConfig(api.ClientConfig{
			BaseURL:    apiBaseURL,
			APIType:    api.APITypeOpenAI,
			HTTPClient: &http.Client{},
		}),
		log: log,
	}
}

func (a *AiEngine) Prompt(ctx context.Context, request *api.ChatCompletionRequest) (*api.ChatCompletionResponse, error) {
	response, err := a.client.CreateChatCompletion(
		ctx,
		*request,
	)

	if err != nil {
		err = lib.WrapError(ErrChatCompletion, err)
		a.log.Error(err)
		return nil, err
	}
	return &response, nil
}

func (a *AiEngine) PromptStream(ctx context.Context, request *api.ChatCompletionRequest, chunkCallback CompletionCallback) (*api.ChatCompletionStreamResponse, error) {
	resp, err := a.requestChatCompletionStream(ctx, request, chunkCallback)

	if err != nil {
		err = lib.WrapError(ErrChatCompletion, err)
		a.log.Error(err)
		return nil, err
	}

	return resp, nil
}

func (a *AiEngine) PromptCb(ctx *gin.Context, body *openai.ChatCompletionRequest) {
	if body.Stream {
		response, err := a.PromptStream(ctx, body, func(response *openai.ChatCompletionStreamResponse) error {
			marshalledResponse, err := json.Marshal(response)
			if err != nil {
				return err
			}

			ctx.Writer.Header().Set(c.HEADER_CONTENT_TYPE, c.CONTENT_TYPE_EVENT_STREAM)

			_, err = ctx.Writer.Write([]byte(fmt.Sprintf("data: %s\n\n", marshalledResponse)))
			if err != nil {
				return err
			}

			ctx.Writer.Flush()
			return nil
		})

		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}

		ctx.JSON(http.StatusOK, response)
		return
	} else {
		response, err := a.Prompt(ctx, body)
		if err != nil {
			ctx.JSON(http.StatusInternalServerError, gin.H{"error": err.Error()})
			return
		}
		ctx.JSON(http.StatusOK, response)
		return
	}
}

func (a *AiEngine) requestChatCompletionStream(ctx context.Context, request *api.ChatCompletionRequest, callback CompletionCallback) (*api.ChatCompletionStreamResponse, error) {
	requestBody, err := json.Marshal(request)
	if err != nil {
		return nil, fmt.Errorf("failed to encode request: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", a.baseURL+"/chat/completions", bytes.NewReader(requestBody))
	if err != nil {
		return nil, fmt.Errorf("failed to create request: %v", err)
	}

	if a.apiKey != "" {
		req.Header.Set(c.HEADER_AUTHORIZATION, fmt.Sprintf("%s %s", c.BEARER, a.apiKey))
	}
	req.Header.Set(c.HEADER_CONTENT_TYPE, c.CONTENT_TYPE_JSON)
	req.Header.Set(c.HEADER_ACCEPT, c.CONTENT_TYPE_EVENT_STREAM)
	req.Header.Set(c.HEADER_CONNECTION, c.CONNECTION_KEEP_ALIVE)

	client := &http.Client{}
	resp, err := client.Do(req)
	if err != nil {
		return nil, fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	scanner := bufio.NewScanner(resp.Body)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, "data: ") {
			data := line[6:] // Skip the "data: " prefix
			var completion api.ChatCompletionStreamResponse
			if err := json.Unmarshal([]byte(data), &completion); err != nil {
				if strings.Index(data, "[DONE]") != -1 {
					a.log.Debugf("reached end of the response")
				} else {
					a.log.Errorf("error decoding response: %s\n%s", err, line)
				}
				continue
			}
			// Call the callback function with the unmarshalled completion
			err := callback(&completion)
			if err != nil {
				return nil, fmt.Errorf("callback failed: %v", err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("error reading stream: %v", err)
	}

	return nil, err
}
