package domain

import "time"

type Concert struct {
	ID             int       `json:"id"`
	Name           string    `json:"name"`
	Date           time.Time `json:"date"`
	Venue          string    `json:"venue"`
	AvailableSeats int       `json:"availableSeats"`
	TicketPrice    float64   `json:"ticketPrice"`
}

type Ticket struct {
	ID           int       `json:"id"`
	ConcertID    int       `json:"concertId"`
	StudentName  string    `json:"studentName"`
	StudentClass string    `json:"studentClass"`
	PurchaseDate time.Time `json:"purchaseDate"`
}
