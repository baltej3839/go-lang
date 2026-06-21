// Create an unbuffered jobs := make(chan int) and start one worker goroutine that continuously receives numbers and prints their squares.

// From main, send 1,2,3,4,5 one at a time and observe that each jobs <- n blocks until the worker performs <-jobs.

package main

import "fmt"

func printSquares(channel chan int){
	number:=<-channel
	fmt.Println(number*number)
}

func main(){
	jobs:=make(chan int)

	numbers:=[]int{1,2,3,4,5}

	go printSquares(jobs)

	for _, v:= range numbers {
		jobs<-v
	}

}