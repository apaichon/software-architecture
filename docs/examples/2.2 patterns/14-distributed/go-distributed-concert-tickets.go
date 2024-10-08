package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net"
	"net/http"
	"net/rpc"
	"sync"
	"time"
)

// Ticket represents a concert ticket
type Ticket struct {
	ID     string
	Price  float64
	Status string
}

// TicketStore manages the tickets
type TicketStore struct {
	sync.RWMutex
	tickets map[string]*Ticket
}

// BookingService handles ticket booking operations
type BookingService struct {
	store *TicketStore
}

// BookingRequest represents a ticket booking request
type BookingRequest struct {
	TicketID string
}

// BookingResponse represents a ticket booking response
type BookingResponse struct {
	Success bool
	Message string
}

// Book attempts to book a ticket
func (s *BookingService) Book(req *BookingRequest, res *BookingResponse) error {
	s.store.Lock()
	defer s.store.Unlock()

	ticket, exists := s.store.tickets[req.TicketID]
	if !exists {
		res.Success = false
		res.Message = "Ticket not found"
		return nil
	}

	if ticket.Status != "Available" {
		res.Success = false
		res.Message = "Ticket not available"
		return nil
	}

	ticket.Status = "Booked"
	res.Success = true
	res.Message = "Ticket booked successfully"
	return nil
}

// LoadBalancer distributes requests among booking servers
type LoadBalancer struct {
	servers []string
	current int
}

func (lb *LoadBalancer) NextServer() string {
	server := lb.servers[lb.current]
	lb.current = (lb.current + 1) % len(lb.servers)
	return server
}

func main() {
	// Initialize ticket store
	store := &TicketStore{
		tickets: make(map[string]*Ticket),
	}

	// Create some sample tickets
	for i := 1; i <= 100; i++ {
		id := fmt.Sprintf("T%d", i)
		store.tickets[id] = &Ticket{
			ID:     id,
			Price:  50.0 + float64(rand.Intn(50)),
			Status: "Available",
		}
	}

	// Start booking servers
	serverAddresses := []string{":8081", ":8082", ":8083"}
	for _, address := range serverAddresses {
		go startBookingServer(address, store)
	}

	// Initialize load balancer
	lb := &LoadBalancer{
		servers: serverAddresses,
	}

	// Start HTTP server for client requests
	http.HandleFunc("/book", func(w http.ResponseWriter, r *http.Request) {
		ticketID := r.URL.Query().Get("ticket")
		if ticketID == "" {
			http.Error(w, "Missing ticket ID", http.StatusBadRequest)
			return
		}

		// Distribute request to a booking server
		server := lb.NextServer()
		client, err := rpc.DialHTTP("tcp", server)
		if err != nil {
			http.Error(w, "Failed to connect to booking server", http.StatusInternalServerError)
			return
		}

		req := &BookingRequest{TicketID: ticketID}
		var res BookingResponse
		err = client.Call("BookingService.Book", req, &res)
		if err != nil {
			http.Error(w, "Booking failed", http.StatusInternalServerError)
			return
		}

		json.NewEncoder(w).Encode(res)
	})

	// Start load testing
	go loadTest(serverAddresses)

	log.Println("Starting HTTP server on :8080")
	log.Fatal(http.ListenAndServe(":8080", nil))
}

func startBookingServer(address string, store *TicketStore) {
	service := &BookingService{store: store}
	rpc.Register(service)
	rpc.HandleHTTP()

	listener, err := net.Listen("tcp", address)
	if err != nil {
		log.Fatal("Listen error:", err)
	}

	log.Printf("Starting RPC server on %s\n", address)
	http.Serve(listener, nil)
}

func loadTest(servers []string) {
	time.Sleep(2 * time.Second) // Wait for servers to start

	for i := 0; i < 1000; i++ {
		go func() {
			ticketID := fmt.Sprintf("T%d", rand.Intn(100)+1)
			server := servers[rand.Intn(len(servers))]
			client, err := rpc.DialHTTP("tcp", server)
			if err != nil {
				log.Printf("Failed to connect to server %s: %v", server, err)
				return
			}

			req := &BookingRequest{TicketID: ticketID}
			var res BookingResponse
			err = client.Call("BookingService.Book", req, &res)
			if err != nil {
				log.Printf("Booking failed: %v", err)
				return
			}

			log.Printf("Booking result for %s: %v - %s", ticketID, res.Success, res.Message)
		}()

		time.Sleep(time.Millisecond * 10) // Simulate some delay between requests
	}
}
