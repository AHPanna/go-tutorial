package main

import (
	"fmt"
	"time"
)

/**
 * select usage on go which is like switch statement but this append on channel.
 */
func window(channel1 chan string) {
	time.Sleep(3 * time.Second)
	channel1 <- "Welcome to Window with channel 1"
}

func door(channel2 chan string) {
	time.Sleep(9 * time.Second)
	channel2 <- "Welcome to door with channle 2"
}
func main() {
	R1 := make(chan string)
	R2 := make(chan string)

	// let run this on same time on multiple threads and channel can be used go routines
	go window(R1)

	go door(R2)
	select {
	// case 1 for window
	case op1 := <-R1:
		fmt.Println(op1)

	// case 2 for door
	case op2 := <-R2:
		fmt.Println(op2)
		//default:
		//fmt.Println("Not found")
	}
}
