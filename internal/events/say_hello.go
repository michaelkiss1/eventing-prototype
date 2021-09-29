package events

import (
	"context"
	"fmt"
	"sync"

	"github.com/mustafaturan/bus/v3"
)

type HelloEvent RegisteredEvent

var c chan bus.Event

var ctx context.Context
var cancel context.CancelFunc

func init() {
	c = make(chan bus.Event)
	ctx, cancel = context.WithCancel(context.Background())
}

func NewHelloEvent(bus *bus.Bus, handlerName string) *HelloEvent {
	return &HelloEvent{
		Bus: bus,
		HandlerName: handlerName,
	}
}

// Start registers the handler
func (h *HelloEvent) Start(wg *sync.WaitGroup) {
	// Matcher looks for topic names to trigger the handler func
	handler := bus.Handler{Handle: SendToChannel, Matcher: ".*"}
	h.Bus.RegisterHandler(h.HandlerName, handler)
	fmt.Printf("Registered %s handler...\n", h.HandlerName)

	wg.Add(1)
	go sayHello(wg)
}

func (h *HelloEvent) Stop() {
	defer fmt.Printf("Deregistered %s handler...\n", h.HandlerName)

	h.Bus.DeregisterHandler(h.HandlerName)
	cancel()
}

func SendToChannel(_ context.Context, e bus.Event) {
	c <- e
}

func sayHello(wg *sync.WaitGroup) {
	defer wg.Done()
	for {
		// Separating the logic from channels would be better. So, please
		// consider this is an example but do not consider as best practice.

		select {
		case <-ctx.Done():
			return
		case e := <-c:
			name := e.Data.(string)
			println("\n\nTopic name: " + e.Topic)
			println("Event ID: " + e.ID)
			fmt.Printf("EVENT: Hello %s, I'm an asychronous event!!\n\n", name)
		}
	}
}
