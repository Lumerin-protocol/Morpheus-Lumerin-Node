package config

import (
	"fmt"

	"github.com/charmbracelet/bubbles/textinput"

	"github.com/Lumerin-protocol/Morpheus-Lumerin-Node/cli/chat/common"
	"github.com/Lumerin-protocol/Morpheus-Lumerin-Node/cli/chat/style"
)

func New(cfgs ...common.Config) model {
	var (
		cfg common.Config
		t   textinput.Model

		isEdit bool
	)

	if len(cfgs) > 0 {
		isEdit = true
		cfg = cfgs[0]
	}

	m := model{inputs: make([]textinput.Model, 5)}
	m.configs = make([]configInput, len(m.inputs))

	for i := range m.configs {
		switch i {
		// case 0:
		// 	m.configs[i].label = "OpenAI API key"
		// 	m.configs[i].name = "openai_api_key"
		case 0:
			m.configs[i].label = "Wallet Address"
			m.configs[i].name = "wallet_address"
		case 1:
			m.configs[i].label = "Model"
			m.configs[i].name = "model"
			m.configs[i].defaultValue = "lamma2"
		case 2:
			m.configs[i].label = "Temperature"
			m.configs[i].name = "temperature"
			m.configs[i].defaultValue = "0.7"
		case 3:
			m.configs[i].label = "Maximum length"
			m.configs[i].name = "max_length"
			m.configs[i].defaultValue = "256"
		case 4:
			m.configs[i].label = "Top P"
			m.configs[i].name = "top_p"
			m.configs[i].defaultValue = "1"
		case 5:
			m.configs[i].label = "Session Length"
			m.configs[i].name = "session_length"
			m.configs[i].defaultValue = "5"
		}
	}

	for i := range m.inputs {
		t = textinput.New()
		t.CursorStyle = style.Cursor
		t.CharLimit = 64

		switch i {
		case 0:
			t.Focus()
			t.Placeholder = m.getPlaceholder(i)
			t.PromptStyle = style.Focused
			t.TextStyle = style.Focused
			t.EchoMode = textinput.EchoPassword
			t.EchoCharacter = '•'

			if isEdit {
				t.SetValue(cfg.OpenaiAPIKey)
			}
		default:
			t.Placeholder = m.getPlaceholder(i)

			if isEdit {
				switch m.configs[i].name {
				case "model":
					t.SetValue(cfg.Model)
				case "temperature":
					t.SetValue(fmt.Sprintf("%f", cfg.Temperature))
				case "max_length":
					t.SetValue(fmt.Sprintf("%d", cfg.MaxLength))
				case "top_p":
					t.SetValue(fmt.Sprintf("%f", cfg.TopP))
				}
			}
		}

		m.inputs[i] = t
	}

	return m
}
