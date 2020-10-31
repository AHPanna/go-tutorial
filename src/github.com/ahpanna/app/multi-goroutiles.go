package main

import (
	"fmt"
	"time"
)

// simple function print names on go routines
func getName() {
	arr := [...]string{"Azures", "Panna", "Simca", "Faucon"}
	for t := 0; t < len(arr); t++ {
		time.Sleep(150 * time.Millisecond)
		fmt.Println(arr[t], "\n")
	}
}

// function should be executed and stay longer than getName
func someIds() {
	arr := [...]int{16, 54566, 69, 78}
	for t := 0; t < len(arr); t++ {
		time.Sleep(500 * time.Millisecond)
		fmt.Println(arr[t], "\n")
	}
}

func main() {
	fmt.Println("********** Start **************\n")
	go getName()
	go someIds()
	time.Sleep(3500 * time.Millisecond)
	fmt.Println("-------- End ------------\n")
}
