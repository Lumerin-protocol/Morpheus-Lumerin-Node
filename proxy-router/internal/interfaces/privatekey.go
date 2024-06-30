package interfaces

import "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"

type PrKeyProvider interface {
	GetPrivateKey() (lib.HexString, error)
	PrivateKeyUpdated() <-chan struct{}
}
