package main

import (
	"fmt"
	"sync"
)

// Blackboard represents the shared knowledge base
type Blackboard struct {
	mu              sync.RWMutex
	inventory       map[string]int
	cart            map[string]int
	pricing         map[string]float64
	orderStatus     string
	paymentStatus   string
	shippingDetails map[string]string
}

// NewBlackboard creates a new instance of Blackboard
func NewBlackboard() *Blackboard {
	return &Blackboard{
		inventory:       make(map[string]int),
		cart:            make(map[string]int),
		pricing:         make(map[string]float64),
		shippingDetails: make(map[string]string),
	}
}

// KnowledgeSource interface defines methods for experts
type KnowledgeSource interface {
	CanContribute(b *Blackboard) bool
	Contribute(b *Blackboard)
}

// InventoryExpert manages product inventory
type InventoryExpert struct{}

func (e *InventoryExpert) CanContribute(b *Blackboard) bool {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return len(b.cart) > 0
}

func (e *InventoryExpert) Contribute(b *Blackboard) {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Check and update inventory
	for product, quantity := range b.cart {
		if currentStock, exists := b.inventory[product]; exists {
			if currentStock >= quantity {
				b.inventory[product] -= quantity
				fmt.Printf("Updated inventory for %s: %d remaining\n", product, b.inventory[product])
			} else {
				fmt.Printf("Insufficient stock for %s\n", product)
			}
		}
	}
}

// PricingExpert calculates total price
type PricingExpert struct{}

func (e *PricingExpert) CanContribute(b *Blackboard) bool {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return len(b.cart) > 0 && len(b.pricing) > 0
}

func (e *PricingExpert) Contribute(b *Blackboard) {
	b.mu.Lock()
	defer b.mu.Unlock()

	total := 0.0
	for product, quantity := range b.cart {
		if price, exists := b.pricing[product]; exists {
			total += price * float64(quantity)
		}
	}
	fmt.Printf("Total order value: $%.2f\n", total)
}

// ShippingExpert handles shipping logistics
type ShippingExpert struct{}

func (e *ShippingExpert) CanContribute(b *Blackboard) bool {
	b.mu.RLock()
	defer b.mu.RUnlock()
	return b.orderStatus == "PAID" && len(b.shippingDetails) > 0
}

func (e *ShippingExpert) Contribute(b *Blackboard) {
	b.mu.Lock()
	defer b.mu.Unlock()

	// Process shipping
	address := b.shippingDetails["address"]
	method := b.shippingDetails["method"]
	fmt.Printf("Processing shipping to %s via %s\n", address, method)
	b.orderStatus = "SHIPPED"
}

// Controller orchestrates the experts
type Controller struct {
	blackboard *Blackboard
	experts    []KnowledgeSource
}

func NewController(b *Blackboard) *Controller {
	return &Controller{
		blackboard: b,
		experts: []KnowledgeSource{
			&InventoryExpert{},
			&PricingExpert{},
			&ShippingExpert{},
		},
	}
}

func (c *Controller) Run() {
	for _, expert := range c.experts {
		if expert.CanContribute(c.blackboard) {
			expert.Contribute(c.blackboard)
		}
	}
}

func main() {
	// Initialize the blackboard
	bb := NewBlackboard()

	// Setup initial data
	bb.inventory = map[string]int{
		"laptop": 10,
		"phone":  20,
		"tablet": 15,
	}

	bb.pricing = map[string]float64{
		"laptop": 999.99,
		"phone":  599.99,
		"tablet": 299.99,
	}

	// Simulate a customer order
	bb.cart = map[string]int{
		"laptop": 1,
		"phone":  2,
	}

	bb.orderStatus = "PAID"
	bb.shippingDetails = map[string]string{
		"address": "123 Main St, City",
		"method":  "Express",
	}

	// Create and run the controller
	controller := NewController(bb)
	controller.Run() 
}
