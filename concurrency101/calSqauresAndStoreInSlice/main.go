
package main

import (
	"fmt"
	"sync"
)

func calculateSquareAndPrint(number int, wg *sync.WaitGroup, squares *[]int, mu *sync.Mutex) {
	defer wg.Done()
	
	mu.Lock()
	*squares =append(*squares,number*number)
	defer mu.Unlock()
	
	fmt.Print("\n",number*number)
	
	
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	var wg sync.WaitGroup
	var mu sync.Mutex	
	var squares = []int{};

	for _, v := range numbers {
		wg.Add(1)
		go calculateSquareAndPrint(v, &wg, &squares, &mu )
	}

	wg.Wait()

	fmt.Println(squares)



}