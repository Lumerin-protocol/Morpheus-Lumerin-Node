package chat

import (
	"net/http"

	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/sashabaranov/go-openai"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/cli/chat/common"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/cli/chat/style"
)

func New(cfg common.Config) model {
	ta := textarea.New()
	ta.Placeholder = "Send your prompt..."
	ta.Focus()

	ta.Prompt = "┃ "
	ta.CharLimit = int(cfg.MaxLength)

	ta.SetHeight(3)

	ta.FocusedStyle.CursorLine = style.Clear
	ta.FocusedStyle.Placeholder = style.Placeholder

	ta.ShowLineNumbers = false
	ta.KeyMap.InsertNewline.SetEnabled(false)

	vp := viewport.New(78, 15)
	vp.SetContent(common.ChatWelcomeMessage)
	vp.Style = style.Viewport

	// client := openai.NewClient(cfg.OpenaiAPIKey)
	client := openai.NewClientWithConfig(openai.ClientConfig{
		BaseURL:    cfg.OpenaiBaseUrl,
		APIType:    openai.APITypeOpenAI,
		HTTPClient: http.DefaultClient,
	})

	req := openai.ChatCompletionRequest{
		MaxTokens:   cfg.MaxLength,
		Model:       cfg.Model,
		Temperature: cfg.Temperature,
		TopP:        cfg.TopP,
		Stream:      true,
	}

	return model{
		config:   cfg,
		err:      nil,
		messages: []string{},
		textarea: ta,
		viewport: vp,

		openaiClient:  client,
		openaiRequest: req,
		completionChunkSub: make(chan string),
	}
}
