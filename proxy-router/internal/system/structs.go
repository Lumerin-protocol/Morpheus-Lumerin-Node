package system

type FD struct {
	ID   string
	Path string
}

type SetEthNodeURLReq struct {
	URLs []string `json:"urls" binding:"required" validate:"required,url"`
}

type ConfigResponse struct {
	Version       string
	Commit        string
	DerivedConfig interface{}
	Config        interface{}
}

type HealthCheckResponse struct {
	Status  string
	Version string
	Uptime  string
}
