package main


import (
	"fmt"
	"time"

	"ix-broker/broker"
)

func main() {
	// Create broker and services
	b := broker.NewBroker()
	broker.NewConcertService(b)
	broker.NewTicketService(b)
	broker.NewNotificationService(b)

	// Create a concert
	concertResult, err := b.SendRequest("concert", broker.Request{
		Type: "CREATE_CONCERT",
		Data: map[string]interface{}{
			"name":              "Rock Festival 2024",
			"date":              time.Date(2024, 8, 15, 0, 0, 0, 0, time.Local),
			"venue":             "Central Park",
			"available_tickets": float64(1000),
			"ticket_price":      float64(99.99),
		},
	})
	if err != nil {
		fmt.Printf("Error creating concert: %v\n", err)
		return
	}
	concert := concertResult.(*broker.Concert)

	// Purchase a ticket
	ticketResult, err := b.SendRequest("ticket", broker.Request{
		Type: "PURCHASE_TICKET",
		Data: map[string]interface{}{
			"concert_id":     concert.ID,
			"customer_name":  "John Doe",
			"customer_email": "john@example.com",
		},
	})
	if err != nil {
		fmt.Printf("Error purchasing ticket: %v\n", err)
		return
	}
	ticket := ticketResult.(*broker.Ticket)

	// List all concerts
	concerts, err := b.SendRequest("concert", broker.Request{
		Type: "LIST_CONCERTS",
	})
	if err != nil {
		fmt.Printf("Error listing concerts: %v\n", err)
		return
	}
	fmt.Printf("All concerts: %+v\n", concerts)

	// Get specific ticket
	retrievedTicket, err := b.SendRequest("ticket", broker.Request{
		Type: "GET_TICKET",
		Data: map[string]interface{}{
			"id": ticket.ID,
		},
	})
	if err != nil {
		fmt.Printf("Error retrieving ticket: %v\n", err)
		return
	}
	fmt.Printf("Retrieved ticket: %+v\n", retrievedTicket)
} 