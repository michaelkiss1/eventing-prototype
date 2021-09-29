package events

import (
	"context"
	"github.com/mustafaturan/bus/v3"
	"sync"
)

type AsyncEvent interface {
	Start(wg *sync.WaitGroup)
	Stop()
	SendToChannel(ctx context.Context, e bus.Event)
}

// Base event type so that we have a good reference to the same bus.
type RegisteredEvent struct {
	Bus *bus.Bus
	HandlerName string
}

func NewRegisteredEvent(bus *bus.Bus, handlerName string) RegisteredEvent {
	return RegisteredEvent{
		Bus: bus,
		HandlerName: handlerName,
	}
}
