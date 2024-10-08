package main

import (
	"errors"
	"fmt"
)

// SagaStep represents a step in the saga
type SagaStep interface {
	Execute() error
	Compensate() error
}

// TicketPurchaseSaga manages the ticket purchase transaction
type TicketPurchaseSaga struct {
	steps       []SagaStep
	completedSteps int
}

func NewTicketPurchaseSaga() *TicketPurchaseSaga {
	return &TicketPurchaseSaga{
		steps: make([]SagaStep, 0),
	}
}

func (s *TicketPurchaseSaga) AddStep(step SagaStep) {
	s.steps = append(s.steps, step)
}

func (s *TicketPurchaseSaga) Execute() error {
	for i, step := range s.steps {
		if err := step.Execute(); err != nil {
			// If a step fails, compensate all completed steps
			for j := i - 1; j >= 0; j-- {
				s.steps[j].Compensate()
			}
			return err
		}
		s.completedSteps++
	}
	return nil
}

// Concrete saga steps
type ReserveTicketStep struct {
	ticketSpace *TicketSpace
	ticket      *Ticket
}

func (s *ReserveTicketStep) Execute() error {
	return s.ticketSpace.Write(s.ticket)
}

func (s *ReserveTicketStep) Compensate() error {
	return s.ticketSpace.Delete(s.ticket.ID)
}

type ProcessPaymentStep struct {
	paymentService PaymentService
	amount         float64
	userID         string
}

func (s *ProcessPaymentStep) Execute() error {
	return s.paymentService.ProcessPayment(s.userID, s.amount)
}

func (s *ProcessPaymentStep) Compensate() error {
	return s.paymentService.RefundPayment(s.userID, s.amount)
}

type SendConfirmationStep struct {
	notificationService NotificationService
	userID              string
}

func (s *SendConfirmationStep) Execute() error {
	return s.notificationService.SendConfirmation(s.userID)
}

func (s *SendConfirmationStep) Compensate() error {
	return s.notificationService.SendCancellation(s.userID)
}

// Mock services for demonstration
type PaymentService struct{}

func (s *PaymentService) ProcessPayment(userID string, amount float64) error {
	fmt.Printf("Processing payment of %.2f for user %s\n", amount, userID)
	return nil
}

func (s *PaymentService) RefundPayment(userID string, amount float64) error {
	fmt.Printf("Refunding payment of %.2f for user %s\n", amount, userID)
	return nil
}

type NotificationService struct{}

func (s *NotificationService) SendConfirmation(userID string) error {
	fmt.Printf("Sending confirmation to user %s\n", userID)
	return nil
}

func (s *NotificationService) SendCancellation(userID string) error {
	fmt.Printf("Sending cancellation notice to user %s\n", userID)
	return nil
}

// Usage
func main() {
	ticketSpace := NewTicketSpace()
	paymentService := &PaymentService{}
	notificationService := &NotificationService{}

	ticket := &Ticket{ID: "T1", ConcertID: "C1", UserID: "U1", Price: 50.0}

	saga := NewTicketPurchaseSaga()
	saga.AddStep(&ReserveTicketStep{ticketSpace: ticketSpace, ticket: ticket})
	saga.AddStep(&ProcessPaymentStep{paymentService: paymentService, amount: ticket.Price, userID: ticket.UserID})
	saga.AddStep(&SendConfirmationStep{notificationService: notificationService, userID: ticket.UserID})

	err := saga.Execute()
	if err != nil {
		fmt.Printf("Saga failed: %v\n", err)
	} else {
		fmt.Println("Saga completed successfully")
	}
}
