package main

import (
	"fmt"
	"sync"
)

// TicketRequest represents a request for a ticket operation
type TicketRequest struct {
	Operation string // e.g., "purchase", "cancel"
	TicketID  string
	UserID    string
	ConcertID string
}

// TicketResponse represents the response to a ticket operation
type TicketResponse struct {
	Success bool
	Message string
	Ticket  *Ticket
}

// TicketBroker acts as an intermediary for ticket operations
type TicketBroker struct {
	ticketService TicketService
	eventBus      EventBus
	requests      chan TicketRequest
	responses     chan TicketResponse
}

func NewTicketBroker(ticketService TicketService, eventBus EventBus) *TicketBroker {
	broker := &TicketBroker{
		ticketService: ticketService,
		eventBus:      eventBus,
		requests:      make(chan TicketRequest),
		responses:     make(chan TicketResponse),
	}
	go broker.processRequests()
	return broker
}

func (b *TicketBroker) processRequests() {
	for request := range b.requests {
		var response TicketResponse
		switch request.Operation {
		case "purchase":
			ticket, err := b.ticketService.PurchaseTicket(request.ConcertID, request.UserID)
			if err != nil {
				response = TicketResponse{Success: false, Message: err.Error()}
			} else {
				response = TicketResponse{Success: true, Message: "Ticket purchased successfully", Ticket: ticket}
				b.eventBus.Publish(TicketPurchasedEvent{TicketID: ticket.ID, ConcertID: ticket.ConcertID, UserID: ticket.UserID})
			}
		case "cancel":
			err := b.ticketService.CancelTicket(request.TicketID)
			if err != nil {
				response = TicketResponse{Success: false, Message: err.Error()}
			} else {
				response = TicketResponse{Success: true, Message: "Ticket cancelled successfully"}
				b.eventBus.Publish(TicketCancelledEvent{TicketID: request.TicketID})
			}
		default:
			response = TicketResponse{Success: false, Message: "Unknown operation"}
		}
		b.responses <- response
	}
}

func (b *TicketBroker) RequestTicketOperation(request TicketRequest) TicketResponse {
	b.requests <- request
	return <-b.responses
}

// Usage
func main() {
	// Create dependencies
	ticketService := &MockTicketService{}
	eventBus := NewInMemoryEventBus()

	// Create broker
	broker := NewTicketBroker(ticketService, eventBus)

	// Subscribe to events
	eventBus.Subscribe("TicketPurchased", func(event Event) {
		if e, ok := event.(TicketPurchasedEvent); ok {
			fmt.Printf("Ticket purchased: %s\n", e.TicketID)
		}
	})

	// Use broker to purchase a ticket
	response := broker.RequestTicketOperation(TicketRequest{
		Operation: "purchase",
		ConcertID: "C1",
		UserID:    "U1",
	})

	fmt.Printf("Operation result: %v, Message: %s\n", response.Success, response.Message)
	if response.Ticket != nil {
		fmt.Printf("Purchased ticket ID: %s\n", response.Ticket.ID)
	}
}

// MockTicketService for demonstration
type MockTicketService struct{}

func (s *MockTicketService) PurchaseTicket(concertID, userID string) (*Ticket, error) {
	return &Ticket{ID: "T1", ConcertID: concertID, UserID: userID}, nil
}

func (s *MockTicketService) CancelTicket(ticketID string) error {
	return nil
}
