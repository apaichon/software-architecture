package queries

import (
	"database/sql"
	"vi-cqrs/domain"
)

type QueryHandler struct {
	db *sql.DB
}

func NewQueryHandler(db *sql.DB) *QueryHandler {
	return &QueryHandler{db: db}
}

func (h *QueryHandler) HandleGetConcertByID(query GetConcertByIDQuery) (*domain.Concert, error) {
	concert := &domain.Concert{}
	err := h.db.QueryRow(`
		SELECT id, name, date, venue, available_seats, ticket_price 
		FROM concerts 
		WHERE id = ?`, query.ID).Scan(
		&concert.ID, &concert.Name, &concert.Date, &concert.Venue,
		&concert.AvailableSeats, &concert.TicketPrice)
	if err != nil {
		return nil, err
	}
	return concert, nil
}

func (h *QueryHandler) HandleGetAvailableConcerts(query GetAvailableConcertsQuery) ([]domain.Concert, error) {
	rows, err := h.db.Query(`
		SELECT id, name, date, venue, available_seats, ticket_price 
		FROM concerts 
		WHERE available_seats >= ?`, query.MinAvailableSeats)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var concerts []domain.Concert
	for rows.Next() {
		var c domain.Concert
		err := rows.Scan(&c.ID, &c.Name, &c.Date, &c.Venue, &c.AvailableSeats, &c.TicketPrice)
		if err != nil {
			return nil, err
		}
		concerts = append(concerts, c)
	}
	return concerts, nil
}

func (h *QueryHandler) HandleGetStudentTickets(query GetStudentTicketsQuery) ([]domain.Ticket, error) {
	rows, err := h.db.Query(`
		SELECT id, concert_id, student_name, student_class, datetime(purchase_date)
		FROM tickets 
		WHERE student_name = ?`, query.StudentName)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var tickets []domain.Ticket
	for rows.Next() {
		var t domain.Ticket
		err := rows.Scan(&t.ID, &t.ConcertID, &t.StudentName, &t.StudentClass, &t.PurchaseDate)
		if err != nil {
			return nil, err
		}
		tickets = append(tickets, t)
	}
	return tickets, nil
}
