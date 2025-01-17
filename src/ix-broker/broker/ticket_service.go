package broker

import (
	"fmt"
	"sync"
	"time"
)

type Ticket struct {
	ID            string    `json:"id"`
	ConcertID     string    `json:"concert_id"`
	CustomerName  string    `json:"customer_name"`
	CustomerEmail string    `json:"customer_email"`
	PurchaseDate  time.Time `json:"purchase_date"`
}

type TicketService struct {
	broker  *Broker
	tickets sync.Map
}

func NewTicketService(broker *Broker) *TicketService {
	service := &TicketService{
		broker: broker,
	}
	broker.RegisterService("ticket", service)
	broker.Subscribe("concertCreated", service.onConcertCreated)
	return service
}

func (s *TicketService) HandleRequest(req Request) (interface{}, error) {
	switch req.Type {
	case "PURCHASE_TICKET":
		return s.purchaseTicket(req.Data)
	case "GET_TICKET":
		return s.getTicket(req.Data["id"].(string))
	default:
		return nil, fmt.Errorf("unknown request type: %s", req.Type)
	}
}

func (s *TicketService) purchaseTicket(data map[string]interface{}) (*Ticket, error) {
	concertID := data["concert_id"].(string)
	
	// Get concert through broker
	result, err := s.broker.SendRequest("concert", Request{
		Type: "GET_CONCERT",
		Data: map[string]interface{}{"id": concertID},
	})
	if err != nil {
		return nil, err
	}

	concert := result.(*Concert)
	if concert.AvailableTickets <= 0 {
		return nil, fmt.Errorf("no tickets available for concert: %s", concertID)
	}

	ticket := &Ticket{
		ID:            fmt.Sprintf("TICKET-%d", time.Now().UnixNano()),
		ConcertID:     concertID,
		CustomerName:  data["customer_name"].(string),
		CustomerEmail: data["customer_email"].(string),
		PurchaseDate:  time.Now(),
	}

	s.tickets.Store(ticket.ID, ticket)
	concert.AvailableTickets--

	s.broker.Publish("ticketPurchased", map[string]interface{}{
		"ticket":    ticket,
		"concertId": concert.ID,
	})

	return ticket, nil
}

func (s *TicketService) getTicket(id string) (*Ticket, error) {
	if ticket, ok := s.tickets.Load(id); ok {
		return ticket.(*Ticket), nil
	}
	return nil, fmt.Errorf("ticket not found: %s", id)
}

func (s *TicketService) onConcertCreated(data interface{}) {
	concert := data.(*Concert)
	fmt.Printf("New concert created: %s. Tickets available: %d\n",
		concert.Name, concert.AvailableTickets)
} 