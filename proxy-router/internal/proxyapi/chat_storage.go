package proxyapi

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sync"
	"time"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/aiengine"
	"github.com/sashabaranov/go-openai"
)

type ChatStorageInterface interface {
	LoadChatFromFile(chatID string) (*ChatHistory, error)
	StorePromptResponseToFile(chatID string, sessionID string, modelID string, prompt openai.ChatCompletionRequest, responses []interface{}, promptAt time.Time, responseAt time.Time) error
	GetChats() []Chat
	DeleteChat(chatID string) error
	UpdateChatTitle(chatID string, title string) error
}

type ChatHistory struct {
	Title    string        `json:"title"`
	ModelId  string        `json:"modelId"`
	Messages []ChatMessage `json:"messages"`
}

type ChatMessage struct {
	Prompt         OpenAiCompletitionRequest `json:"prompt"`
	Response       string                    `json:"response"`
	PromptAt       int64                     `json:"promptAt"`
	ResponseAt     int64                     `json:"responseAt"`
	IsImageContent bool                      `json:"isImageContent"`
}
type Chat struct {
	ChatID    string `json:"chatId"`
	ModelID   string `json:"modelId"`
	Title     string `json:"title"`
	CreatedAt int64  `json:"createdAt"`
}

// ChatStorage handles storing conversations to files.
type ChatStorage struct {
	dirPath     string                 // Directory path to store the files
	fileMutexes map[string]*sync.Mutex // Map to store mutexes for each file
}

// NewChatStorage creates a new instance of ChatStorage.
func NewChatStorage(dirPath string) *ChatStorage {
	return &ChatStorage{
		dirPath:     dirPath,
		fileMutexes: make(map[string]*sync.Mutex),
	}
}

// StorePromptResponseToFile stores the prompt and response to a file.
func (cs *ChatStorage) StorePromptResponseToFile(identifier string, sessionId string, modelId string, prompt openai.ChatCompletionRequest, responses []interface{}, promptAt, responseAt time.Time) error {
	if err := os.MkdirAll(cs.dirPath, os.ModePerm); err != nil {
		return err
	}

	filePath := filepath.Join(cs.dirPath, identifier+".json")
	cs.initFileMutex(filePath)

	// Lock the file mutex
	cs.fileMutexes[filePath].Lock()
	defer cs.fileMutexes[filePath].Unlock()

	var data ChatHistory
	if _, err := os.Stat(filePath); err == nil {
		fileContent, err := os.ReadFile(filePath)
		if err != nil {
			return err
		}
		if err := json.Unmarshal(fileContent, &data); err != nil {
			return err
		}
	}

	response := ""
	var isImageContent bool = false
	for _, r := range responses {
		switch v := r.(type) {
		case ChatCompletionResponse:
			response += fmt.Sprintf("%v", v.Choices[0].Delta.Content)
		case *openai.ChatCompletionStreamResponse:
			response += fmt.Sprintf("%v", v.Choices[0].Delta.Content)
		case aiengine.ProdiaGenerationResult:
			response += fmt.Sprintf("%v", v.ImageUrl)
			isImageContent = true
		case *aiengine.ProdiaGenerationResult:
			response += fmt.Sprintf("%v", v.ImageUrl)
			isImageContent = true
		default:
			return fmt.Errorf("unknown response type")
		}
	}

	messages := make([]ChatCompletionMessage, 0)
	for _, r := range prompt.Messages {
		messages = append(messages, ChatCompletionMessage{
			Content: r.Content,
			Role:    r.Role,
		})
	}

	p := OpenAiCompletitionRequest{
		Messages:         messages,
		Model:            prompt.Model,
		MaxTokens:        prompt.MaxTokens,
		Temperature:      prompt.Temperature,
		TopP:             prompt.TopP,
		FrequencyPenalty: prompt.FrequencyPenalty,
		PresencePenalty:  prompt.PresencePenalty,
		Stop:             prompt.Stop,
	}

	newEntry := ChatMessage{
		Prompt:         p,
		Response:       response,
		PromptAt:       promptAt.Unix(),
		ResponseAt:     responseAt.Unix(),
		IsImageContent: isImageContent,
	}

	if data.Messages == nil && len(data.Messages) == 0 {
		data.ModelId = modelId
		data.Title = prompt.Messages[0].Content
	}

	newMessages := append(data.Messages, newEntry)
	data.Messages = newMessages

	updatedContent, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(filePath, updatedContent, 0644); err != nil {
		return err
	}

	return nil
}

func (cs *ChatStorage) GetChats() []Chat {
	var chats []Chat
	files, err := os.ReadDir(cs.dirPath)
	if err != nil {
		return chats
	}

	for _, file := range files {
		if file.IsDir() {
			continue
		}

		chatID := file.Name()
		chatID = chatID[:len(chatID)-5]

		fileContent, err := cs.LoadChatFromFile(chatID)
		if err != nil {
			continue
		}
		chats = append(chats, Chat{
			ChatID:    chatID,
			Title:     fileContent.Title,
			CreatedAt: fileContent.Messages[0].PromptAt,
			ModelID:   fileContent.ModelId,
		})
	}

	return chats
}

func (cs *ChatStorage) DeleteChat(identifier string) error {
	filePath := filepath.Join(cs.dirPath, identifier+".json")
	cs.initFileMutex(filePath)

	cs.fileMutexes[filePath].Lock()
	defer cs.fileMutexes[filePath].Unlock()

	if err := os.Remove(filePath); err != nil {
		return err
	}
	return nil
}

func (cs *ChatStorage) UpdateChatTitle(identifier string, title string) error {
	chat, err := cs.LoadChatFromFile(identifier)
	if err != nil {
		return err
	}
	chat.Title = title

	filePath := filepath.Join(cs.dirPath, identifier+".json")
	cs.initFileMutex(filePath)

	cs.fileMutexes[filePath].Lock()
	defer cs.fileMutexes[filePath].Unlock()

	updatedContent, err := json.MarshalIndent(chat, "", "  ")
	if err != nil {
		return err
	}

	if err := os.WriteFile(filePath, updatedContent, 0644); err != nil {
		return err
	}

	return nil
}

func (cs *ChatStorage) LoadChatFromFile(identifier string) (*ChatHistory, error) {
	filePath := filepath.Join(cs.dirPath, identifier+".json")
	cs.initFileMutex(filePath)

	cs.fileMutexes[filePath].Lock()
	defer cs.fileMutexes[filePath].Unlock()

	var data ChatHistory
	fileContent, err := os.ReadFile(filePath)
	if err != nil {
		return &data, err
	}

	if err := json.Unmarshal(fileContent, &data); err != nil {
		fmt.Println("Error unmarshalling file content:", err)
		return nil, err
	}

	return &data, nil
}

// initFileMutex initializes a mutex for the file if not already present.
func (cs *ChatStorage) initFileMutex(filePath string) {
	if _, exists := cs.fileMutexes[filePath]; !exists {
		cs.fileMutexes[filePath] = &sync.Mutex{}
	}
}

type NoOpChatStorage struct{}

func NewNoOpChatStorage() *NoOpChatStorage {
	return &NoOpChatStorage{}
}

func (cs *NoOpChatStorage) LoadChatFromFile(chatID string) (*ChatHistory, error) {
	return nil, nil
}

func (cs *NoOpChatStorage) StorePromptResponseToFile(chatID string, sessionID string, modelID string, prompt openai.ChatCompletionRequest, responses []interface{}, promptAt time.Time, responseAt time.Time) error {
	return nil
}

func (cs *NoOpChatStorage) GetChats() []Chat {
	return []Chat{}
}

func (cs *NoOpChatStorage) DeleteChat(chatID string) error {
	return nil
}

func (cs *NoOpChatStorage) UpdateChatTitle(chatID string, title string) error {
	return nil
}
