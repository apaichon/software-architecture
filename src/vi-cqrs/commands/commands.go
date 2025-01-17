package commands

import "time"

type PurchaseTicketCommand struct {
	ConcertID    int       `json:"concertId"`
	StudentName  string    `json:"studentName"`
	StudentClass string    `json:"studentClass"`
	PurchaseDate time.Time `json:"purchaseDate"`
}

type CreateConcertCommand struct {
	Name           string    `json:"name"`
	Date           time.Time `json:"date"`
	Venue          string    `json:"venue"`
	AvailableSeats int       `json:"availableSeats"`
	TicketPrice    float64   `json:"ticketPrice"`
}
