package common

type Config struct {
	MaxLength     int     `json:"max_length" validate:"required,number"`
	Model         string  `json:"model" validate:"required"`
	WalletAddress string  `json:"wallet_address" validate:"required,startswith=0x"`
	WalletKey     string  `json:"wallet_key" validate:"required"`
	Temperature   float32 `json:"temperature" validate:"required,number"`
	TopP          float32 `json:"top_p" validate:"required,number"`
	OpenaiAPIKey  string  `json:"openai_api_key"`
	OpenaiBaseUrl string  `json:"openai_base_url" validate:"required,uri"`
	SessionId     string  `json:"session_id"`
}