package main

import (
	"database/sql"
	"errors"
	"fmt"
	_ "github.com/lib/pq" // PostgreSQL driver
)

// TicketRepository interface (already defined in previous steps)
type TicketRepository interface {
	Save(ticket *Ticket) error
	FindByID(id string) (*Ticket, error)
	Delete(id string) error
}

// TicketRepositoryImpl implements TicketRepository
type TicketRepositoryImpl struct {
	db *sql.DB
}

func NewTicketRepositoryImpl(connectionString string) (*TicketRepositoryImpl, error) {
	db, err := sql.Open("postgres", connectionString)
	if err != nil {
		return nil, err
	}
	return &TicketRepositoryImpl{db: db}, nil
}

func (r *TicketRepositoryImpl) Save(ticket *Ticket) error {
	query := `
		INSERT INTO tickets (id, concert_id, user_id, price, created_at)
		VALUES ($1, $2, $3, $4, $5)
		ON CONFLICT (id) DO UPDATE
		SET concert_id = $2, user_id = $3, price = $4, created_at = $5
	`
	_, err := r.db.Exec(query, ticket.ID, ticket.ConcertID, ticket.UserID, ticket.Price, ticket.CreatedAt)
	return err
}

func (r *TicketRepositoryImpl) FindByID(id string) (*Ticket, error) {
	query := `SELECT id, concert_id, user_id, price, created_at FROM tickets WHERE id = $1`
	row := r.db.QueryRow(query, id)

	var ticket Ticket
	err := row.Scan(&ticket.ID, &ticket.ConcertID, &ticket.UserID, &ticket.Price, &ticket.CreatedAt)
	if err == sql.ErrNoRows {
		return nil, errors.New("ticket not found")
	}
	if err != nil {
		return nil, err
	}
	return &ticket, nil
}

func (r *TicketRepositoryImpl) Delete(id string) error {
	query := `DELETE FROM tickets WHERE id = $1`
	result, err := r.db.Exec(query, id)
	if err != nil {
		return err
	}
	rowsAffected, err := result.RowsAffected()
	if err != nil {
		return err
	}
	if rowsAffected == 0 {
		return errors.New("ticket not found")
	}
	return nil
}
