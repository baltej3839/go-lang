package main

import (
	"fmt"
	"math/rand/v2"
	"sync"
	"time"
)

type Notification struct {
	ID      int
	Type    string // "SMS", "Email", or "Push"
	UserID  int
	Message string
}

type Receipt struct {
	NotificationID int
	WorkerID       int
	Status         string // "SUCCESS" or "FAILED"
}

func TypeSelectorRandom() string {
	types := []string{"SMS", "Email", "Push"}
	randomIndex := rand.IntN(len(types))
	randomVal := types[randomIndex]
	return randomVal
}

const charset = "abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ0123456789"

func generateRandomString(length int) string {
	b := make([]byte, length)
	for i := range b {
		// Select a random index from the charset
		b[i] = charset[rand.IntN(len(charset))]
	}
	return string(b)
}

func getStatus() string {
	// rand.Float64() returns a random float between 0.0 and 1.0
	if rand.Float64() < 0.9 {
		return "success"
	}
	return "failure"
}

func workers(incomingRequests chan Notification, receipts chan Receipt, workerId int, wg *sync.WaitGroup) {
	defer wg.Done()
	for notification := range incomingRequests {

		time.Sleep(time.Millisecond * 10)
		receipt := Receipt{
			NotificationID: notification.ID,
			WorkerID:       workerId,
			Status:         getStatus(),
		}
		receipts <- receipt

	}
}

func logger(receipts chan Receipt) {
	var TotalProcessed int = 0
	var successCount int = 0
	var failedCount int = 0
	for r := range receipts {
		TotalProcessed++
		if r.Status == "success" {
			successCount++
		} else {
			failedCount++
		}
		fmt.Printf("[Worker %s] %s sent Notification #%s (SMS) \n", r.WorkerID, r.Status, r.NotificationID)
	}
	defer fmt.Printf("Total Processed: %d | Success: %d | Failed: %d", TotalProcessed, successCount, failedCount)
}

func main() {
	incomingRequests := make(chan Notification, 20)
	receipts := make(chan Receipt, 20)
	var wg sync.WaitGroup
	for i := 0; i < 20; i++ {
		notification := Notification{
			ID:      i + 1,
			Type:    TypeSelectorRandom(),
			UserID:  (i + 1) * 1000,
			Message: generateRandomString(12),
		}
		incomingRequests <- notification
	}
	close(incomingRequests)
	wg.Add(3)
	for i := 0; i < 3; i++ {
		go workers(incomingRequests, receipts, i+1, &wg)
	}

	go logger(receipts)

	go func() {
		wg.Wait()
		close(receipts)
	}()
	



}
