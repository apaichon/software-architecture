package main

import (
	"errors"
	"time"
)

// Command models
type PurchaseTicketCommand struct {
	ConcertID string
	UserID    string
}

type CancelTicketCommand struct {
	TicketID string
}

// Command handlers
type TicketCommandHandler interface {
	HandlePurchaseTicket(cmd PurchaseTicketCommand) error
	HandleCancelTicket(cmd CancelTicketCommand) error
}

type TicketCommandHandlerImpl struct {
	ticketRepo  TicketRepository
	concertRepo ConcertRepository
}

func (h *TicketCommandHandlerImpl) HandlePurchaseTicket(cmd PurchaseTicketCommand) error {
	concert, err := h.concertRepo.FindByID(cmd.ConcertID)
	if err != nil {
		return err
	}

	if concert.Capacity <= 0 {
		return errors.New("concert is sold out")
	}

	ticket := &Ticket{
		ID:        generateID(),
		ConcertID: cmd.ConcertID,
		UserID:    cmd.UserID,
		Price:     calculatePrice(concert),
		CreatedAt: time.Now(),
	}

	err = h.ticketRepo.Save(ticket)
	if err != nil {
		return err
	}

	return h.concertRepo.UpdateCapacity(cmd.ConcertID, concert.Capacity-1)
}

func (h *TicketCommandHandlerImpl) HandleCancelTicket(cmd CancelTicketCommand) error {
	ticket, err := h.ticketRepo.FindByID(cmd.TicketID)
	if err != nil {
		return err
	}

	err = h.ticketRepo.Delete(cmd.TicketID)
	if err != nil {
		return err
	}

	concert, err := h.concertRepo.FindByID(ticket.ConcertID)
	if err != nil {
		return err
	}

	return h.concertRepo.UpdateCapacity(concert.ID, concert.Capacity+1)
}

// Query models
type GetUserTicketsQuery struct {
	UserID string
}

type GetConcertQuery struct {
	ConcertID string
}

// Query handlers
type TicketQueryHandler interface {
	HandleGetUserTickets(query GetUserTicketsQuery) ([]Ticket, error)
	HandleGetConcert(query GetConcertQuery) (*Concert, error)
}

type TicketQueryHandlerImpl struct {
	ticketRepo  TicketRepository
	concertRepo ConcertRepository
}

func (h *TicketQueryHandlerImpl) HandleGetUserTickets(query GetUserTicketsQuery) ([]Ticket, error) {
	// Implementation to fetch user tickets
	// This would typically involve a new method on the TicketRepository
	return nil, errors.New("not implemented")
}

func (h *TicketQueryHandlerImpl) HandleGetConcert(query GetConcertQuery) (*Concert, error) {
	return h.concertRepo.FindByID(query.ConcertID)
}
