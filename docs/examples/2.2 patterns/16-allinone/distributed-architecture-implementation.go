package main

import (
	"fmt"
	"math/rand"
	"sync"
	"time"
)

// TicketNode represents an individual node in the distributed system
type TicketNode struct {
	ID           string
	TicketSpace  *TicketSpace
	Peers        map[string]*TicketNode
	EventBus     EventBus
	mu           sync.RWMutex
}

func NewTicketNode(id string, eventBus EventBus) *TicketNode {
	return &TicketNode{
		ID:          id,
		TicketSpace: NewTicketSpace(),
		Peers:       make(map[string]*TicketNode),
		EventBus:    eventBus,
	}
}

func (n *TicketNode) AddPeer(peer *TicketNode) {
	n.mu.Lock()
	defer n.mu.Unlock()
	n.Peers[peer.ID] = peer
}

func (n *TicketNode) RemovePeer(peerID string) {
	n.mu.Lock()
	defer n.mu.Unlock()
	delete(n.Peers, peerID)
}

func (n *TicketNode) PurchaseTicket(concertID, userID string) (*Ticket, error) {
	ticket := &Ticket{
		ID:        generateID(),
		ConcertID: concertID,
		UserID:    userID,
		Price:     calculatePrice(concertID),
	}

	err := n.TicketSpace.Write(ticket)
	if err != nil {
		return nil, err
	}

	// Propagate the change to peers
	n.propagateChange("purchase", ticket)

	return ticket, nil
}

func (n *TicketNode) CancelTicket(ticketID string) error {
	err := n.TicketSpace.Delete(ticketID)
	if err != nil {
		return err
	}

	// Propagate the change to peers
	n.propagateChange("cancel", &Ticket{ID: ticketID})

	return nil
}

func (n *TicketNode) propagateChange(operation string, ticket *Ticket) {
	n.mu.RLock()
	defer n.mu.RUnlock()

	for _, peer := range n.Peers {
		go func(p *TicketNode) {
			switch operation {
			case "purchase":
				p.TicketSpace.Write(ticket)
			case "cancel":
				p.TicketSpace.Delete(ticket.ID)
			}
		}(peer)
	}

	// Publish event
	if operation == "purchase" {
		n.EventBus.Publish(TicketPurchasedEvent{TicketID: ticket.ID, ConcertID: ticket.ConcertID, UserID: ticket.UserID})
	} else if operation == "cancel" {
		n.EventBus.Publish(TicketCancelledEvent{TicketID: ticket.ID})
	}
}

// Simulated network latency
func simulateNetworkLatency() {
	time.Sleep(time.Duration(rand.Intn(100)) * time.Millisecond)
}

// Usage
func main() {
	eventBus := NewInMemoryEventBus()

	// Create nodes
	node1 := NewTicketNode("Node1", eventBus)
	node2 := NewTicketNode("Node2", eventBus)
	node3 := NewTicketNode("Node3", eventBus)

	// Connect nodes
	node1.AddPeer(node2)
	node1.AddPeer(node3)
	node2.AddPeer(node1)
	node2.AddPeer(node3)
	node3.AddPeer(node1)
	node3.AddPeer(node2)

	// Subscribe to events
	eventBus.Subscribe("TicketPurchased", func(event Event) {
		if e, ok := event.(TicketPurchasedEvent); ok {
			fmt.Printf("Ticket purchased: %s\n", e.TicketID)
		}
	})

	// Purchase a ticket on node1
	ticket, err := node1.PurchaseTicket("C1", "U1")
	if err != nil {
		fmt.Printf("Error purchasing ticket: %v\n", err)
	} else {
		fmt.Printf("Ticket purchased on Node1: %+v\n", ticket)
	}

	// Simulate network latency
	simulateNetworkLatency()

	// Verify the ticket exists on node2
	readTicket, err := node2.TicketSpace.Read(ticket.ID)
	if err != nil {
		fmt.Printf("Error reading ticket from Node2: %v\n", err)
	} else {
		fmt.Printf("Ticket read from Node2: %+v\n", readTicket)
	}

	// Cancel the ticket on node3
	err = node3.CancelTicket(ticket.ID)
	if err != nil {
		fmt.Printf("Error cancelling ticket on Node3: %v\n", err)
	} else {
		fmt.Println("Ticket cancelled successfully on Node3")
	}

	// Simulate network latency
	simulateNetworkLatency()

	// Verify the ticket is cancelled on node1
	_, err = node1.TicketSpace.Read(ticket.ID)
	if err != nil {
		fmt.Println("Ticket successfully removed from Node1")
	} else {
		fmt.Println("Error: Ticket still exists on Node1")
	}
}
