PS C:\Users\HarleenKaur\Desktop\go-projects\mapRecursionTypes> go run .
fatal error: all goroutines are asleep - deadlock!

goroutine 1 [sync.WaitGroup.Wait]:
sync.runtime_SemacquireWaitGroup(0x1d2a493ea2a0?, 0x0?)
        C:/Program Files/Go/src/runtime/sema.go:114 +0x2e
sync.(*WaitGroup).Wait(0x1d2a493d0090)
        C:/Program Files/Go/src/sync/waitgroup.go:206 +0x85
main.main()
        C:/Users/HarleenKaur/Desktop/go-projects/mapRecursionTypes/main.go:77 +0x1af

goroutine 7 [sync.Mutex.Lock]:
internal/sync.runtime_SemacquireMutex(0x0?, 0x0?, 0x0?)
        C:/Program Files/Go/src/runtime/sema.go:95 +0x25
internal/sync.(*Mutex).lockSlow(0x1d2a493d00a0)
        C:/Program Files/Go/src/internal/sync/mutex.go:149 +0x15d
internal/sync.(*Mutex).Lock(...)
        C:/Program Files/Go/src/internal/sync/mutex.go:70
sync.(*Mutex).Lock(...)
        C:/Program Files/Go/src/sync/mutex.go:46
main.calculateSquareAndPrint(0x0?, 0x0?, 0x0?, 0x1d2a493d00a0, 0x1d2a493ca060)
        C:/Users/HarleenKaur/Desktop/go-projects/mapRecursionTypes/main.go:53 +0x7f
created by main.main in goroutine 1
        C:/Users/HarleenKaur/Desktop/go-projects/mapRecursionTypes/main.go:72 +0xf0

goroutine 8 [sync.Mutex.Lock]:
internal/sync.runtime_SemacquireMutex(0x0?, 0x0?, 0x0?)
        C:/Program Files/Go/src/runtime/sema.go:95 +0x25
internal/sync.(*Mutex).lockSlow(0x1d2a493d00a0)
        C:/Program Files/Go/src/internal/sync/mutex.go:149 +0x15d
internal/sync.(*Mutex).Lock(...)
        C:/Program Files/Go/src/internal/sync/mutex.go:70
sync.(*Mutex).Lock(...)
        C:/Program Files/Go/src/sync/mutex.go:46
main.calculateSquareAndPrint(0x0?, 0x0?, 0x0?, 0x1d2a493d00a0, 0x1d2a493ca060)
        C:/Users/HarleenKaur/Desktop/go-projects/mapRecursionTypes/main.go:53 +0x7f
created by main.main in goroutine 1
        C:/Users/HarleenKaur/Desktop/go-projects/mapRecursionTypes/main.go:72 +0xf0

goroutine 9 [sync.Mutex.Lock]:
internal/sync.runtime_SemacquireMutex(0x0?, 0x0?, 0x0?)
        C:/Program Files/Go/src/runtime/sema.go:95 +0x25
internal/sync.(*Mutex).lockSlow(0x1d2a493d00a0)
        C:/Program Files/Go/src/internal/sync/mutex.go:149 +0x15d
internal/sync.(*Mutex).Lock(...)
        C:/Program Files/Go/src/internal/sync/mutex.go:70
sync.(*Mutex).Lock(...)
        C:/Program Files/Go/src/sync/mutex.go:46
main.calculateSquareAndPrint(0x0?, 0x0?, 0x0?, 0x1d2a493d00a0, 0x1d2a493ca060)
        C:/Users/HarleenKaur/Desktop/go-projects/mapRecursionTypes/main.go:53 +0x7f
created by main.main in goroutine 1
        C:/Users/HarleenKaur/Desktop/go-projects/mapRecursionTypes/main.go:72 +0xf0

goroutine 10 [sync.Mutex.Lock]:
internal/sync.runtime_SemacquireMutex(0x0?, 0x0?, 0x0?)
        C:/Program Files/Go/src/runtime/sema.go:95 +0x25
internal/sync.(*Mutex).lockSlow(0x1d2a493d00a0)
        C:/Program Files/Go/src/internal/sync/mutex.go:149 +0x15d
internal/sync.(*Mutex).Lock(...)
        C:/Program Files/Go/src/internal/sync/mutex.go:70
sync.(*Mutex).Lock(...)
        C:/Program Files/Go/src/sync/mutex.go:46
main.calculateSquareAndPrint(0x0?, 0x0?, 0x0?, 0x1d2a493d00a0, 0x1d2a493ca060)
        C:/Users/HarleenKaur/Desktop/go-projects/mapRecursionTypes/main.go:53 +0x7f
created by main.main in goroutine 1
        C:/Users/HarleenKaur/Desktop/go-projects/mapRecursionTypes/main.go:72 +0xf0

goroutine 11 [chan send]:
main.calculateSquareAndPrint(0x0?, 0x0?, 0x0?, 0x1d2a493d00a0, 0x0?)
        C:/Users/HarleenKaur/Desktop/go-projects/mapRecursionTypes/main.go:55 +0xa7
created by main.main in goroutine 1
        C:/Users/HarleenKaur/Desktop/go-projects/mapRecursionTypes/main.go:72 +0xf0
exit status 2


for this code 
// package main

// import (
// 	"fmt"
// 	"go/types"
// 	"sync"
// )

// func calculateSquareAndPrint(number int, wg *sync.WaitGroup, squares *[]int, mu *sync.Mutex, messages *chan int) {
// 	// defer wg.Done()
	
// 	// mu.Lock()
// 	*messages<-number*number
// 	// defer mu.Unlock()
	
// 	fmt.Print("\n",number*number)
	
	
// }

// func main() {
// 	numbers := []int{1, 2, 3, 4, 5}
// 	// var wg sync.WaitGroup
// 	// var mu sync.Mutex	
// 	// var squares = []int{};
// 	messages:=make(chan int)

// 	for _, v := range numbers {
// 		wg.Add(1)
// 		go calculateSquareAndPrint(v, &messages )

// 	}

	

// 	for  {
// 		fmt.Println(messages)
// 		if messages.(types)!=int {
// 			break
// 		}
// 	}


// }



solution
package main

import (
	"fmt"
)

func calculateSquareAndPrint(number int, messages chan int) {
	square := number * number
	fmt.Printf("%d ", square)
	
	messages <- square // Main will wait for this
}

func main() {
	numbers := []int{1, 2, 3, 4, 5}
	
	// Create a buffered channel
	messages := make(chan int, len(numbers))

	for _, v := range numbers {
		go calculateSquareAndPrint(v, messages)
	}

	// Instead of wg.Wait(), we just read exactly 5 times.
	// This naturally blocks main until all 5 goroutines send their data.
	var squares []int
	for i := 0; i < len(numbers); i++ {
		squares = append(squares, <-messages)
	}

	fmt.Println("\nFinal Slice:", squares)
}