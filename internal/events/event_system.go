package events

import (
	"errors"
	"fmt"
	"github.com/google/uuid"
	"github.com/mustafaturan/bus/v3"
)
// Simplifying Eventing by scoping what we can do with it.
type EventSystem struct {
	// Making this private to reduce potential event bus modifications hidden in code
	Bus *bus.Bus
	RegisteredEvents []RegisteredEvent
}

func NewEventSystem(eventTopics []string) *EventSystem {
	bus := newBus(eventTopics)

	return &EventSystem{
		Bus: bus,
	}
}

func newBus(topics []string) *bus.Bus {
	// init an id generator
	var id bus.Next = uuid.NewString

	// create a new bus instance
	b, err := bus.NewBus(id)
	if err != nil {
		panic(err)
	}

	// Safe to initialize all topics here
	b.RegisterTopics(topics...)

	println("Bus created successfully.")
	return b
}

func (e *EventSystem) RegisterEvent(event RegisteredEvent) error {
	for _, value := range e.RegisteredEvents {
		if event.HandlerName == value.HandlerName {
			return errors.New(fmt.Sprintf("Event %s already exists within the event system", value.HandlerName))
		}
	}

	e.RegisteredEvents = append(e.RegisteredEvents, event)

	return nil
}

// Debugging purposes
func (e *EventSystem) PrintRegisteredTopics() {
	topics := e.Bus.Topics()

	for _, v := range topics {
		println(v)
	}
}

