


package main

import (
	"fmt"
	// "go/types"
	// "sync"
)

func calculateSquareAndPrint(number int,  messages *chan int) {
	// defer wg.Done()
	
	// mu.Lock()
	*messages<-number*number
	// defer mu.Unlock()
	
	fmt.Print("\n",number*number)
	
	
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	// var wg sync.WaitGroup
	// var mu sync.Mutex	
	// var squares = []int{};
	messages:=make(chan int, len(numbers))

	for _, v := range numbers {
		// wg.Add(1)
		go calculateSquareAndPrint(v, &messages)

	}

		var squares = []int{};

		for i := 0; i < len(numbers); i++ {
			squares=append(squares,<-messages)
		}

		fmt.Println("\nFinal Slice:", squares)
	


}