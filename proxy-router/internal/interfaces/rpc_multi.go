package interfaces

type RPCEndpoints interface {
	GetURLs() []string
	SetURLs(urls []string) error
	RemoveURLs() error
}
