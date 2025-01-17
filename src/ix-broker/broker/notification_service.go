package broker

import "fmt"

type NotificationService struct {
	broker *Broker
}

func NewNotificationService(broker *Broker) *NotificationService {
	service := &NotificationService{
		broker: broker,
	}
	broker.Subscribe("ticketPurchased", service.onTicketPurchased)
	return service
}

func (s *NotificationService) onTicketPurchased(data interface{}) {
	purchaseData := data.(map[string]interface{})
	ticket := purchaseData["ticket"].(*Ticket)
	concertID := purchaseData["concertId"].(string)

	fmt.Printf("Ticket purchased for concert %s. Sending confirmation email to %s\n",
		concertID, ticket.CustomerEmail)
} 