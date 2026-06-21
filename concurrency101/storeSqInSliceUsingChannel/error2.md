fatal error: all goroutines are asleep - deadlock!

goroutine 1 [chan receive]:
main.main()
        C:/Users/HarleenKaur/Desktop/go-projects/mapRecursionTypes/main.go:35 +0xf8
exit status 2

code
package main 

import (
	"fmt"
	// "go/types"
	// "sync"
)

func calculateSquareAndPrintNumbers(num  int, results chan int) {
	// defer wg.Done()
	
	// num:=<-jobs
	sq:=num*num 
	// mu.Lock()
	results<-sq
 	// defer mu.Unlock()
	
	// fmt.Print("\n",number*number)
	
	
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	// var wg sync.WaitGroup
	// var mu sync.Mutex	
	// var squares = []int{};
	jobs:=make(chan int, len(numbers))
	results:=make(chan int, len(numbers))
	for _,v:= range numbers {
		jobs<- v 
	}


	for v:= range jobs {
		go calculateSquareAndPrintNumbers(v, results)
	}
	close(jobs)

	// for _, v := range numbers {
	// 	// wg.Add(1)
	// 	go calculateSquareAndPrintNumbers(v, &messages)

	// }

		var squares = []int{};

		for i := 0; i < len(numbers); i++ {
			squares=append(squares,<-results)
		}

		close(results)
		fmt.Println("\nFinal Slice:", squares)
	


}

