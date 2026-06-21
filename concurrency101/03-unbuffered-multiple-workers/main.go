// Create two worker goroutines that both receive from the same unbuffered jobs channel and print:

// Worker 1 processed 3
// Worker 2 processed 4

// Observe how Go automatically distributes jobs between workers.

package main

import (
	"fmt"
)

func printSquares(channel chan int, done chan bool, worker string) {
	for number := range channel {
		fmt.Printf("worker %s processed number %d and square is %d \n", worker, number, number*number)
	}
	done <- true
}

func main() {
	jobs := make(chan int)
	done := make(chan bool)

	numbers := []int{1, 2, 3, 4, 5}

	go printSquares(jobs, done, "1")
	go printSquares(jobs, done, "2")

	for _, v := range numbers {
		jobs <- v
	}

	close(jobs)
	<-done
	<-done

}