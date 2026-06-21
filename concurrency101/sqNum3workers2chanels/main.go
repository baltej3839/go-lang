package main

import (
	"fmt"
	"sync"
)

func calcualateSquareOfNumber(jobs, results chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for num := range jobs {
		sq := num * num
		results <- sq
	}
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	jobs := make(chan int, len(numbers))
	results := make(chan int, len(numbers))
	var wg sync.WaitGroup

	for _, v := range numbers {
		jobs <- v
	}
	close(jobs)

	wg.Add(3)
	for i := 0; i < 3; i++ {
		go calcualateSquareOfNumber(jobs, results, &wg  )
	}
	
	go func(){
		wg.Wait()
		close(results)
	}()
	

	for v := range results {
		fmt.Println("squareed number is ", v)
	}

}
