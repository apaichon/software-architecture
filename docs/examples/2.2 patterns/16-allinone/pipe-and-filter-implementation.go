package main

import (
	"errors"
	"fmt"
)

// TicketProcessor interface
type TicketProcessor interface {
	Process(ticket *Ticket) error
}

// ValidationFilter
type ValidationFilter struct{}

func (f *ValidationFilter) Process(ticket *Ticket) error {
	if ticket.UserID == "" || ticket.ConcertID == "" {
		return errors.New("invalid ticket: missing user ID or concert ID")
	}
	return nil
}

// PricingFilter
type PricingFilter struct {
	BasePrice float64
}

func (f *PricingFilter) Process(ticket *Ticket) error {
	// Simple pricing logic, could be more complex
	ticket.Price = f.BasePrice
	return nil
}

// AvailabilityFilter
type AvailabilityFilter struct {
	ConcertRepository ConcertRepository
}

func (f *AvailabilityFilter) Process(ticket *Ticket) error {
	concert, err := f.ConcertRepository.FindByID(ticket.ConcertID)
	if err != nil {
		return err
	}
	if concert.Capacity <= 0 {
		return errors.New("concert is sold out")
	}
	return nil
}

// Pipeline struct to manage the sequence of filters
type Pipeline struct {
	filters []TicketProcessor
}

func NewPipeline() *Pipeline {
	return &Pipeline{filters: make([]TicketProcessor, 0)}
}

func (p *Pipeline) AddFilter(filter TicketProcessor) {
	p.filters = append(p.filters, filter)
}

func (p *Pipeline) Process(ticket *Ticket) error {
	for _, filter := range p.filters {
		if err := filter.Process(ticket); err != nil {
			return err
		}
	}
	return nil
}

// Usage
func main() {
	// Create filters
	validationFilter := &ValidationFilter{}
	pricingFilter := &PricingFilter{BasePrice: 50.0}
	availabilityFilter := &AvailabilityFilter{ConcertRepository: &MockConcertRepository{}}

	// Create and configure pipeline
	pipeline := NewPipeline()
	pipeline.AddFilter(validationFilter)
	pipeline.AddFilter(pricingFilter)
	pipeline.AddFilter(availabilityFilter)

	// Process a ticket
	ticket := &Ticket{ID: "T1", ConcertID: "C1", UserID: "U1"}
	err := pipeline.Process(ticket)
	if err != nil {
		fmt.Printf("Error processing ticket: %v\n", err)
	} else {
		fmt.Printf("Ticket processed successfully. Price: %.2f\n", ticket.Price)
	}
}

// Mock ConcertRepository for demonstration
type MockConcertRepository struct{}

func (r *MockConcertRepository) FindByID(id string) (*Concert, error) {
	return &Concert{ID: id, Capacity: 100}, nil
}
