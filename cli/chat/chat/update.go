package chat

import (
	"fmt"
	"os"
	"strings"

	tea "github.com/charmbracelet/bubbletea"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/cli/chat/common"
	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/cli/chat/style"
	"github.com/sashabaranov/go-openai"
)

func (m model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	var (
		tiCmd tea.Cmd
		vpCmd tea.Cmd
	)

	m.textarea, tiCmd = m.textarea.Update(msg)
	m.viewport, vpCmd = m.viewport.Update(msg)

	switch msg := msg.(type) {
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "esc":
			return m, tea.Quit
		case "enter":
			val := m.textarea.Value()
			if val == "" {
				return m, nil
			}

			m.textarea.Reset()

			switch val {
			case "/c", "/clear":
				m.viewport.SetContent(common.ChatWelcomeMessage)
				m.messages = []string{}
			default:
				m.messages = append(m.messages, userMessage(val), botMessage(""))

				m.viewport.SetContent(strings.Join(m.messages, "\n\n") + "\n\n")

				m.viewport.GotoBottom()

				return m, tea.Batch(tiCmd, vpCmd, streamCompletions(m, val), waitForCompletion(m.completionChunkSub))
			}
		}
	case completionMsg:

		newMessage := botMessage(string(msg))

		trimmedMessages := m.messages[:len(m.messages)-1]

		file, _ := os.OpenFile("./test.txt", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)

		m.messages = append(trimmedMessages, wordWrap(newMessage, 78, file))

		messageContent := strings.Join(m.messages, "\n\n")

		m.viewport.SetContent(messageContent + "\n\n")
		
		file.WriteString(fmt.Sprintf("at bottom? %v", m.viewport.AtBottom()))
		file.WriteString(fmt.Sprintf("past bottom? %v", m.viewport.PastBottom()))

		return m, tea.Batch(tiCmd, vpCmd, waitForCompletion(m.completionChunkSub))
	}

	return m, tea.Batch(tiCmd, vpCmd)
}

func userMessage(val string) string {
	return fmt.Sprintf("%s %s", style.Sender.Render("👤:"), val)
}

func botMessage(val string) string {
	return fmt.Sprintf("%s %s", style.Response.Render("🤖:"), val)
}

// WordWrap wraps the given string to fit within a specified width.
func wordWrap(text string, width int, file *os.File) string {
	words := strings.Fields(text)
	if len(words) == 0 {
		return ""
	}

	var wrappedText strings.Builder
	currentLineLength := 0

	for _, word := range words {
		wordLength := len(word)
		if currentLineLength+wordLength+5 > width { // +1 for the space
			if currentLineLength > 0 {
				wrappedText.WriteString("\n")
				currentLineLength = 0
			}
		}

		if currentLineLength > 0 {
			wrappedText.WriteString(" ")
			currentLineLength++
		}
		wrappedText.WriteString(word)
		currentLineLength += wordLength
	}

	file.WriteString("wrapped text: \n")
	file.WriteString(wrappedText.String() + "\n\n\n")
	file.WriteString("text: \n")
	file.WriteString(text + "\n\n\n")

	return wrappedText.String()
}

type completionMsg string

// Simulate a process that sends events at an irregular interval in real time.
// In this case, we'll send events on the channel at a random interval between
// 100 to 1000 milliseconds. As a command, Bubble Tea will run this
// asynchronously.
func streamCompletions(m model, val string) tea.Cmd {
	return func() tea.Msg {

		newChunk := ""

		m.err = m.sendChat(val, func(completion *openai.ChatCompletionStreamResponse) error {

			if completion.Choices[0].Delta.Content != "[DONE]" {
				newChunk += completion.Choices[0].Delta.Content
				m.completionChunkSub <- newChunk
			}

			return nil
		})

		return nil
	}
}

// A command that waits for the next chunk on a channel.
func waitForCompletion(sub chan string) tea.Cmd {
	return func() tea.Msg {
		chunk := <-sub

		return completionMsg(chunk)
	}
}
