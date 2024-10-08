package main

import (
	"fmt"
	"log"
	"math/rand"
	"time"
)

// Ticket represents a concert ticket
type Ticket struct {
	ID     string
	Status string
}

// TicketService manages ticket operations
type TicketService struct {
	tickets map[string]*Ticket
}

func (ts *TicketService) ReserveTicket(id string) error {
	if ticket, exists := ts.tickets[id]; exists && ticket.Status == "Available" {
		ticket.Status = "Reserved"
		return nil
	}
	return fmt.Errorf("unable to reserve ticket %s", id)
}

func (ts *TicketService) CancelReservation(id string) error {
	if ticket, exists := ts.tickets[id]; exists && ticket.Status == "Reserved" {
		ticket.Status = "Available"
		return nil
	}
	return fmt.Errorf("unable to cancel reservation for ticket %s", id)
}

// PaymentService manages payment operations
type PaymentService struct{}

func (ps *PaymentService) ProcessPayment(amount float64) error {
	if rand.Float64() < 0.8 { // 80% chance of success
		return nil
	}
	return fmt.Errorf("payment processing failed")
}

func (ps *PaymentService) RefundPayment(amount float64) error {
	// In a real system, this would interact with a payment gateway
	log.Printf("Refunded $%.2f", amount)
	return nil
}

// NotificationService manages sending notifications
type NotificationService struct{}

func (ns *NotificationService) SendConfirmation(email string) error {
	// In a real system, this would send an actual email
	log.Printf("Sent confirmation email to %s", email)
	return nil
}

func (ns *NotificationService) SendCancellation(email string) error {
	// In a real system, this would send an actual email
	log.Printf("Sent cancellation email to %s", email)
	return nil
}

// Saga orchestrates the ticket booking process
type Saga struct {
	ticketService      *TicketService
	paymentService     *PaymentService
	notificationService *NotificationService
}

func (s *Saga) BookTicket(ticketID string, amount float64, email string) error {
	// Step 1: Reserve Ticket
	if err := s.ticketService.ReserveTicket(ticketID); err != nil {
		return err
	}

	// Step 2: Process Payment
	if err := s.paymentService.ProcessPayment(amount); err != nil {
		// Compensation: Cancel Ticket Reservation
		s.ticketService.CancelReservation(ticketID)
		return err
	}

	// Step 3: Send Confirmation
	if err := s.notificationService.SendConfirmation(email); err != nil {
		// Compensation: Refund Payment and Cancel Ticket Reservation
		s.paymentService.RefundPayment(amount)
		s.ticketService.CancelReservation(ticketID)
		return err
	}

	log.Printf("Successfully booked ticket %s for $%.2f", ticketID, amount)
	return nil
}

func main() {
	// Initialize services
	ticketService := &TicketService{
		tickets: map[string]*Ticket{
			"T1": {ID: "T1", Status: "Available"},
			"T2": {ID: "T2", Status: "Available"},
			"T3": {ID: "T3", Status: "Available"},
		},
	}
	paymentService := &PaymentService{}
	notificationService := &NotificationService{}

	// Create Saga
	saga := &Saga{
		ticketService:      ticketService,
		paymentService:     paymentService,
		notificationService: notificationService,
	}

	// Simulate multiple booking attempts
	for i := 0; i < 5; i++ {
		ticketID := fmt.Sprintf("T%d", rand.Intn(3)+1)
		amount := 50.0 + float64(rand.Intn(50))
		email := fmt.Sprintf("user%d@example.com", i+1)

		log.Printf("Attempting to book ticket %s for $%.2f", ticketID, amount)
		err := saga.BookTicket(ticketID, amount, email)
		if err != nil {
			log.Printf("Booking failed: %v", err)
		}
		log.Println("---")
		time.Sleep(time.Second) // Pause between attempts
	}
}
