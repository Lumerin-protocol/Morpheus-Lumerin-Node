package common

type Config struct {
	MaxLength    int     `json:"max_length" validate:"required,number"`
	Model        string  `json:"model" validate:"required"`
	OpenaiAPIKey string  `json:"openai_api_key" validate:"required,startswith=sk-"`
	WalletAddress string  `json:"wallet_address" validate:"required,startswith=0x"`
	Temperature  float32 `json:"temperature" validate:"required,number"`
	TopP         float32 `json:"top_p" validate:"required,number"`
}
