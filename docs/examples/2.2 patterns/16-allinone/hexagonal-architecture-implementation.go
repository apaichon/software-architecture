package main

import (
	"errors"
	"time"
)

// Domain entities
type Ticket struct {
	ID        string
	ConcertID string
	UserID    string
	Price     float64
	CreatedAt time.Time
}

type Concert struct {
	ID       string
	Name     string
	Date     time.Time
	Capacity int
}

// Ports (interfaces)
type TicketService interface {
	PurchaseTicket(concertID, userID string) (*Ticket, error)
	CancelTicket(ticketID string) error
}

type TicketRepository interface {
	Save(ticket *Ticket) error
	FindByID(id string) (*Ticket, error)
	Delete(id string) error
}

type ConcertRepository interface {
	FindByID(id string) (*Concert, error)
	UpdateCapacity(id string, newCapacity int) error
}

// Adapter (implementation of TicketService)
type TicketServiceImpl struct {
	ticketRepo  TicketRepository
	concertRepo ConcertRepository
}

func NewTicketService(tr TicketRepository, cr ConcertRepository) TicketService {
	return &TicketServiceImpl{
		ticketRepo:  tr,
		concertRepo: cr,
	}
}

func (s *TicketServiceImpl) PurchaseTicket(concertID, userID string) (*Ticket, error) {
	concert, err := s.concertRepo.FindByID(concertID)
	if err != nil {
		return nil, err
	}

	if concert.Capacity <= 0 {
		return nil, errors.New("concert is sold out")
	}

	ticket := &Ticket{
		ID:        generateID(), // Implement this function to generate unique IDs
		ConcertID: concertID,
		UserID:    userID,
		Price:     calculatePrice(concert), // Implement this function to calculate the ticket price
		CreatedAt: time.Now(),
	}

	err = s.ticketRepo.Save(ticket)
	if err != nil {
		return nil, err
	}

	err = s.concertRepo.UpdateCapacity(concertID, concert.Capacity-1)
	if err != nil {
		// If updating capacity fails, we should roll back the ticket creation
		_ = s.ticketRepo.Delete(ticket.ID)
		return nil, err
	}

	return ticket, nil
}

func (s *TicketServiceImpl) CancelTicket(ticketID string) error {
	ticket, err := s.ticketRepo.FindByID(ticketID)
	if err != nil {
		return err
	}

	err = s.ticketRepo.Delete(ticketID)
	if err != nil {
		return err
	}

	concert, err := s.concertRepo.FindByID(ticket.ConcertID)
	if err != nil {
		return err
	}

	return s.concertRepo.UpdateCapacity(concert.ID, concert.Capacity+1)
}

// Helper functions (to be implemented)
func generateID() string {
	// Implementation to generate unique ID
	return ""
}

func calculatePrice(concert *Concert) float64 {
	// Implementation to calculate ticket price
	return 0
}
