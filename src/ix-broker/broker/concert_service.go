package broker

import (
	"fmt"
	"sync"
	"time"
)

type Concert struct {
	ID               string    `json:"id"`
	Name             string    `json:"name"`
	Date             time.Time `json:"date"`
	Venue            string    `json:"venue"`
	AvailableTickets int       `json:"available_tickets"`
	TicketPrice      float64   `json:"ticket_price"`
}

type ConcertService struct {
	broker   *Broker
	concerts sync.Map
}

func NewConcertService(broker *Broker) *ConcertService {
	service := &ConcertService{
		broker: broker,
	}
	broker.RegisterService("concert", service)
	return service
}

func (s *ConcertService) HandleRequest(req Request) (interface{}, error) {
	switch req.Type {
	case "CREATE_CONCERT":
		return s.createConcert(req.Data)
	case "GET_CONCERT":
		return s.getConcert(req.Data["id"].(string))
	case "LIST_CONCERTS":
		return s.listConcerts()
	default:
		return nil, fmt.Errorf("unknown request type: %s", req.Type)
	}
}

func (s *ConcertService) createConcert(data map[string]interface{}) (*Concert, error) {
	concert := &Concert{
		ID:               fmt.Sprintf("CONCERT-%d", time.Now().UnixNano()),
		Name:             data["name"].(string),
		Date:             data["date"].(time.Time),
		Venue:            data["venue"].(string),
		AvailableTickets: int(data["available_tickets"].(float64)),
		TicketPrice:      data["ticket_price"].(float64),
	}

	s.concerts.Store(concert.ID, concert)
	s.broker.Publish("concertCreated", concert)
	return concert, nil
}

func (s *ConcertService) getConcert(id string) (*Concert, error) {
	if concert, ok := s.concerts.Load(id); ok {
		return concert.(*Concert), nil
	}
	return nil, fmt.Errorf("concert not found: %s", id)
}

func (s *ConcertService) listConcerts() ([]*Concert, error) {
	var concerts []*Concert
	s.concerts.Range(func(_, value interface{}) bool {
		concerts = append(concerts, value.(*Concert))
		return true
	})
	return concerts, nil
} 