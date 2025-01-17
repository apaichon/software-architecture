package broker

import (
	"fmt"
	"sync"
)

// Request represents a service request
type Request struct {
	Type string
	Data map[string]interface{}
}

// Service interface defines methods that all services must implement
type Service interface {
	HandleRequest(Request) (interface{}, error)
}

// Callback is a function type for event subscribers
type Callback func(interface{})

// Broker manages services and event subscriptions
type Broker struct {
	services      map[string]Service
	subscriptions map[string][]Callback
	mu            sync.RWMutex
}

// NewBroker creates a new broker instance
func NewBroker() *Broker {
	return &Broker{
		services:      make(map[string]Service),
		subscriptions: make(map[string][]Callback),
	}
}

// RegisterService registers a service with the broker
func (b *Broker) RegisterService(serviceName string, service Service) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.services[serviceName] = service
}

// Subscribe adds a callback for an event
func (b *Broker) Subscribe(event string, callback Callback) {
	b.mu.Lock()
	defer b.mu.Unlock()
	b.subscriptions[event] = append(b.subscriptions[event], callback)
}

// Publish sends an event to all subscribers
func (b *Broker) Publish(event string, data interface{}) {
	b.mu.RLock()
	defer b.mu.RUnlock()
	if callbacks, exists := b.subscriptions[event]; exists {
		for _, callback := range callbacks {
			callback(data)
		}
	}
}

// SendRequest sends a request to a specific service
func (b *Broker) SendRequest(serviceName string, request Request) (interface{}, error) {
	b.mu.RLock()
	service, exists := b.services[serviceName]
	b.mu.RUnlock()

	if !exists {
		return nil, fmt.Errorf("service %s not found", serviceName)
	}
	return service.HandleRequest(request)
} 