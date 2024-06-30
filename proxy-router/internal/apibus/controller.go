package apibus

import "github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/interfaces"

type APIBus struct {
	controllers []Registrable
}

type Registrable interface {
	RegisterRoutes(r interfaces.Router)
}

func NewApiBus(controllers ...Registrable) *APIBus {
	return &APIBus{
		controllers: controllers,
	}
}

func (a *APIBus) RegisterRoutes(r interfaces.Router) {
	for _, c := range a.controllers {
		c.RegisterRoutes(r)
	}
}
