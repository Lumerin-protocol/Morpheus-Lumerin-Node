package contracts

import (
	"time"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/repositories/contracts/bindings/sessionrouter"
	"github.com/ethereum/go-ethereum/core/types"
)

const RECONNECT_TIMEOUT = 2 * time.Second

type EventMapper func(types.Log) (interface{}, error)

func SessionRouterEventFactory(name string) interface{} {
	switch name {
	case "SessionOpened":
		return new(sessionrouter.SessionRouterSessionOpened)
	case "SessionClosed":
		return new(sessionrouter.SessionRouterSessionClosed)
	default:
		return nil
	}
}
