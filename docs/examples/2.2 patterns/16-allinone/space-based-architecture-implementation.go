package main

import (
	"errors"
	"fmt"
	"sync"
)

// TicketSpace represents a distributed in-memory data grid
type TicketSpace struct {
	tickets map[string]*Ticket
	mu      sync.RWMutex
}

func NewTicketSpace() *TicketSpace {
	return &TicketSpace{
		tickets: make(map[string]*Ticket),
	}
}

func (s *TicketSpace) Write(ticket *Ticket) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	s.tickets[ticket.ID] = ticket
	return nil
}

func (s *TicketSpace) Read(id string) (*Ticket, error) {
	s.mu.RLock()
	defer s.mu.RUnlock()

	ticket, exists := s.tickets[id]
	if !exists {
		return nil, errors.New("ticket not found")
	}
	return ticket, nil
}

func (s *TicketSpace) Delete(id string) error {
	s.mu.Lock()
	defer s.mu.Unlock()

	delete(s.tickets, id)
	return nil
}

// SpaceBasedTicketService implements TicketService using TicketSpace
type SpaceBasedTicketService struct {
	space *TicketSpace
}

func NewSpaceBasedTicketService(space *TicketSpace) *SpaceBasedTicketService {
	return &SpaceBasedTicketService{space: space}
}

func (s *SpaceBasedTicketService) PurchaseTicket(concertID, userID string) (*Ticket, error) {
	ticket := &Ticket{
		ID:        generateID(),
		ConcertID: concertID,
		UserID:    userID,
		Price:     calculatePrice(concertID),
	}
	err := s.space.Write(ticket)
	if err != nil {
		return nil, err
	}
	return ticket, nil
}

func (s *SpaceBasedTicketService) CancelTicket(ticketID string) error {
	return s.space.Delete(ticketID)
}

// Helper functions (to be implemented)
func generateID() string {
	return "T1" // Simplified for this example
}

func calculatePrice(concertID string) float64 {
	return 50.0 // Simplified for this example
}

// Usage
func main() {
	space := NewTicketSpace()
	service := NewSpaceBasedTicketService(space)

	// Purchase a ticket
	ticket, err := service.PurchaseTicket("C1", "U1")
	if err != nil {
		fmt.Printf("Error purchasing ticket: %v\n", err)
	} else {
		fmt.Printf("Ticket purchased: %+v\n", ticket)
	}

	// Read the ticket
	readTicket, err := space.Read(ticket.ID)
	if err != nil {
		fmt.Printf("Error reading ticket: %v\n", err)
	} else {
		fmt.Printf("Read ticket: %+v\n", readTicket)
	}

	// Cancel the ticket
	err = service.CancelTicket(ticket.ID)
	if err != nil {
		fmt.Printf("Error cancelling ticket: %v\n", err)
	} else {
		fmt.Println("Ticket cancelled successfully")
	}
}
