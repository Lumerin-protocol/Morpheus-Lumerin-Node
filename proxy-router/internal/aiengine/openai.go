package aiengine

import (
	"bufio"
	"bytes"
	"context"
	"encoding/json"
	"fmt"
	"io"
	"net/http"
	"strings"

	c "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal"
	gcs "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/chatstorage/genericchatstorage"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/sashabaranov/go-openai"
)

const API_TYPE_OPENAI = "openai"

type OpenAI struct {
	baseURL   string
	apiKey    string
	modelName string
	client    *http.Client
	log       lib.ILogger
}

func NewOpenAIEngine(modelName, baseURL, apiKey string, log lib.ILogger) *OpenAI {
	return &OpenAI{
		baseURL:   baseURL,
		modelName: modelName,
		apiKey:    apiKey,
		client:    &http.Client{},
		log:       log,
	}
}

func (a *OpenAI) Prompt(ctx context.Context, compl *openai.ChatCompletionRequest, cb gcs.CompletionCallback) error {
	compl.Model = a.modelName
	requestBody, err := json.Marshal(compl)
	if err != nil {
		return fmt.Errorf("failed to encode request: %v", err)
	}

	req, err := http.NewRequestWithContext(ctx, "POST", a.baseURL+"/chat/completions", bytes.NewReader(requestBody))
	if err != nil {
		return fmt.Errorf("failed to create request: %v", err)
	}

	if a.apiKey != "" {
		req.Header.Set(c.HEADER_AUTHORIZATION, fmt.Sprintf("%s %s", c.BEARER, a.apiKey))
	}
	req.Header.Set(c.HEADER_CONTENT_TYPE, c.CONTENT_TYPE_JSON)
	req.Header.Set(c.HEADER_CONNECTION, c.CONNECTION_KEEP_ALIVE)
	if compl.Stream {
		req.Header.Set(c.HEADER_ACCEPT, c.CONTENT_TYPE_EVENT_STREAM)
	}
	resp, err := a.client.Do(req)
	if err != nil {
		return fmt.Errorf("failed to send request: %v", err)
	}
	defer resp.Body.Close()

	if isContentTypeStream(resp.Header) {
		return a.readStream(ctx, resp.Body, cb)
	}

	return a.readResponse(ctx, resp.Body, cb)
}

func (a *OpenAI) readResponse(ctx context.Context, body io.Reader, cb gcs.CompletionCallback) error {
	var compl openai.ChatCompletionResponse
	if err := json.NewDecoder(body).Decode(&compl); err != nil {
		return fmt.Errorf("failed to decode response: %v", err)
	}

	text := make([]string, len(compl.Choices))
	for i, choice := range compl.Choices {
		text[i] = choice.Message.Content
	}

	chunk := gcs.NewChunkText(&compl)
	err := cb(ctx, chunk)
	if err != nil {
		return fmt.Errorf("callback failed: %v", err)
	}

	return nil
}

func (a *OpenAI) readStream(ctx context.Context, body io.Reader, cb gcs.CompletionCallback) error {
	scanner := bufio.NewScanner(body)
	for scanner.Scan() {
		line := scanner.Text()

		if strings.HasPrefix(line, StreamDataPrefix) {
			data := line[len(StreamDataPrefix):] // Skip the "data: " prefix
			var compl openai.ChatCompletionStreamResponse
			if err := json.Unmarshal([]byte(data), &compl); err != nil {
				if isStreamFinished(data) {
					a.log.Debugf("reached end of the response")
					return nil
				} else {
					return fmt.Errorf("error decoding response: %s\n%s", err, line)
				}
			}
			// Call the callback function with the unmarshalled completion
			chunk := gcs.NewChunkStreaming(&compl)
			err := cb(ctx, chunk)
			if err != nil {
				return fmt.Errorf("callback failed: %v", err)
			}
		}
	}

	if err := scanner.Err(); err != nil {
		return fmt.Errorf("error reading stream: %v", err)
	}

	return nil
}

func (a *OpenAI) ApiType() string {
	return API_TYPE_OPENAI
}

func isStreamFinished(data string) bool {
	return strings.Index(data, StreamDone) != -1
}

func isContentTypeStream(header http.Header) bool {
	contentType := header.Get(c.HEADER_CONTENT_TYPE)
	cTypeParams := strings.Split(contentType, ";")
	cType := strings.TrimSpace(cTypeParams[0])
	return cType == c.CONTENT_TYPE_EVENT_STREAM
}

const (
	StreamDone       = "[DONE]"
	StreamDataPrefix = "data: "
)

var _ AIEngineStream = &OpenAI{}
