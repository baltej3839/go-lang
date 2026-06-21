package main

import (
	"fmt"
	"time"
)

type banton struct {
	tt time.Time
}

func Run(playerName string, race chan *banton) {
	for {
		timer := <-race
		fmt.Println(playerName, "hit the ball")
		time.Sleep(time.Millisecond * 400) // Changed to Millisecond so you can watch it log comfortably
		
		// FIXED: Used .Add() instead of +
		race <- &banton{tt: timer.tt.Add(time.Millisecond * 234)} 
	}
}

func main() {
	race := make(chan *banton)

	go Run("player1", race)
	go Run("player2", race)
	go Run("player3", race)
	go Run("player4", race)
	
	// Start the race
	race <- &banton{tt: time.Now()}

	// Let them play for 4 seconds
	time.Sleep(time.Second * 4)

	// Snatches the baton away to stop the game
	endTime := <-race

	fmt.Printf("\nTotal time is %s\n", endTime.tt)
}