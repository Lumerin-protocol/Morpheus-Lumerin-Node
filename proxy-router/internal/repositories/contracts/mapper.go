package contracts

import (
	"errors"
	"fmt"

	"github.com/MorpheusAIs/Morpheus-Lumerin-Node/proxy-router/internal/lib"
	"github.com/ethereum/go-ethereum/accounts/abi"
	"github.com/ethereum/go-ethereum/core/types"
)

var (
	ErrUnknownEvent = errors.New("unknown event")
)

type EventFactory func(name string) interface{}

func CreateEventMapper(eventFactory EventFactory, abi *abi.ABI) func(log types.Log) (interface{}, error) {
	return func(log types.Log) (interface{}, error) {
		namedEvent, err := abi.EventByID(log.Topics[0])
		if err != nil {
			return nil, err
		}
		concreteEvent := eventFactory(namedEvent.Name)

		if concreteEvent == nil {
			return nil, lib.WrapError(ErrUnknownEvent, fmt.Errorf("event: %s", namedEvent.Name))
		}

		return concreteEvent, UnpackLog(concreteEvent, namedEvent.Name, log, abi)
	}
}
