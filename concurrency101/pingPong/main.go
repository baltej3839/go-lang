package main

import (
	"fmt"
	"time"
)

type Ball struct {
	Hit int
}

func hitBall(name string, ballCheck chan *Ball) {
	for {
		hit:=<-ballCheck
		hit.Hit++  
		
		time.Sleep(time.Millisecond*200)

		fmt.Println("Player", name, "hits and hit count is", hit.Hit)
		
		ballCheck<-hit 
	}
}

func mai() {
	ballCheck := make(chan *Ball)

	go hitBall("Alice",ballCheck)
	go hitBall("Bob",ballCheck)

	ballCheck<-&Ball{Hit: 0}

	time.Sleep(time.Second*2)

	finalBall:=<-ballCheck

	fmt.Print("matched ended", finalBall)

}
