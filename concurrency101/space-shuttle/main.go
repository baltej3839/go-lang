package main

import "fmt"

type Shuttle struct {
	systemACheck bool
	systemBCheck bool
}

func SystemACheck(channel chan *Shuttle) {
	temp := <-channel
	temp.systemACheck = true
	channel <- temp

}

func SystemBCheck(channel chan *Shuttle) {
	temp := <-channel
	temp.systemBCheck = true
	channel <- temp

}

func main() {
	sysChan := make(chan *Shuttle)

	go SystemACheck(sysChan)
	go SystemBCheck(sysChan)

	shuttle := &Shuttle{
		systemACheck: false,
		systemBCheck: false,
	}

	sysChan <- shuttle

	finalResult := <-sysChan

	fmt.Println(finalResult.systemACheck, finalResult.systemBCheck)

}
