package main

import (
	"fmt"
	"time"
)

type order struct {
	orderNumber int
	cardNumber  int
	address     string
	isValidated bool
	isCharged   bool
	isShipped   bool
}

// 1. Receives raw order, validates it, and passes it to the card channel
func ValidateOrder(inChan chan *order, outChan chan *order) {
	orderDetail := <-inChan
	fmt.Printf("[Stage 1] Validating order #%d...\n", orderDetail.orderNumber)
	time.Sleep(200 * time.Millisecond) // Simulate work
	
	orderDetail.isValidated = true
	outChan <- orderDetail // Pass to ChargeCard
}

// 2. Receives validated order, charges it, and passes it to the ship channel
func ChargeCard(inChan chan *order, outChan chan *order) {
	orderDetail := <-inChan
	fmt.Printf("[Stage 2] Charging card for order #%d...\n", orderDetail.orderNumber)
	time.Sleep(200 * time.Millisecond)
	
	orderDetail.isCharged = true
	outChan <- orderDetail // Pass to ShipItem
}

// 3. Receives charged order, ships it, and passes it to the final finish channel
func ShipItem(inChan chan *order, outChan chan *order) {
	orderDetail := <-inChan
	fmt.Printf("[Stage 3] Shipping order #%d to %s...\n", orderDetail.orderNumber, orderDetail.address)
	time.Sleep(200 * time.Millisecond)
	
	orderDetail.isShipped = true
	outChan <- orderDetail // Pass back to main (Finish line)
}

func main() {
	// We need 3 channels to connect 3 stages sequentially + 1 for the finish line
	validateChan := make(chan *order)
	cardChan := make(chan *order)
	shipChan := make(chan *order)
	finishChan := make(chan *order)

	// Wire them up sequentially
	go ValidateOrder(validateChan, cardChan)
	go ChargeCard(cardChan, shipChan)
	go ShipItem(shipChan, finishChan)

	// Create our single customer order
	newOrder := &order{
		orderNumber: 4509,
		cardNumber:  12345678,
		address:     "123 Go Lane",
	}

	fmt.Println("--- Pipeline Started ---")
	
	// Inject the order into the start of the pipeline
	validateChan <- newOrder

	// Main blocks here, waiting for the order to make it through all 3 stages
	finalOrder := <-finishChan

	fmt.Println("\n--- Pipeline Finished ---")
	fmt.Printf("Order Status -> Validated: %t | Charged: %t | Shipped: %t\n", 
		finalOrder.isValidated, finalOrder.isCharged, finalOrder.isShipped)
}