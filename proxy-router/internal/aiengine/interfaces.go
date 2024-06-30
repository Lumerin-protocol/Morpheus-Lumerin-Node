package aiengine

import (
	"net/http"

	api "github.com/sashabaranov/go-openai"
)

type ResponderFlusher interface {
	http.ResponseWriter
	http.Flusher
}

type CompletionCallback func(completion *api.ChatCompletionStreamResponse) error
