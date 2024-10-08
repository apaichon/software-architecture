package main

import (
	"sync"
)

// Event interface
type Event interface {
	EventType() string
}

// Concrete events
type TicketPurchasedEvent struct {
	TicketID  string
	ConcertID string
	UserID    string
}

func (e TicketPurchasedEvent) EventType() string {
	return "TicketPurchased"
}

type TicketCancelledEvent struct {
	TicketID string
}

func (e TicketCancelledEvent) EventType() string {
	return "TicketCancelled"
}

// EventHandler type
type EventHandler func(event Event)

// EventBus interface
type EventBus interface {
	Publish(event Event)
	Subscribe(eventType string, handler EventHandler)
}

// EventBus implementation
type InMemoryEventBus struct {
	handlers map[string][]EventHandler
	mu       sync.RWMutex
}

func NewInMemoryEventBus() *InMemoryEventBus {
	return &InMemoryEventBus{
		handlers: make(map[string][]EventHandler),
	}
}

func (b *InMemoryEventBus) Publish(event Event) {
	b.mu.RLock()
	defer b.mu.RUnlock()

	if handlers, ok := b.handlers[event.EventType()]; ok {
		for _, handler := range handlers {
			go handler(event)
		}
	}
}

func (b *InMemoryEventBus) Subscribe(eventType string, handler EventHandler) {
	b.mu.Lock()
	defer b.mu.Unlock()

	b.handlers[eventType] = append(b.handlers[eventType], handler)
}

// Usage example
func main() {
	eventBus := NewInMemoryEventBus()

	// Subscribe to TicketPurchased events
	eventBus.Subscribe("TicketPurchased", func(event Event) {
		if e, ok := event.(TicketPurchasedEvent); ok {
			println("Ticket purchased:", e.TicketID)
		}
	})

	// Subscribe to TicketCancelled events
	eventBus.Subscribe("TicketCancelled", func(event Event) {
		if e, ok := event.(TicketCancelledEvent); ok {
			println("Ticket cancelled:", e.TicketID)
		}
	})

	// Publish events
	eventBus.Publish(TicketPurchasedEvent{TicketID: "T1", ConcertID: "C1", UserID: "U1"})
	eventBus.Publish(TicketCancelledEvent{TicketID: "T2"})
}
