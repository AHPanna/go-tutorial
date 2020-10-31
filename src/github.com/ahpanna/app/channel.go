package main

import (
	"fmt"
	"sync"
)

var wg sync.WaitGroup // for syncing

/**
 * see the channel as socket on c, or a communication platform can communicate with multiple threads.
 * so here which a goroutine communicates with another goroutine and this communication is lock-free
 * in a nutshell this read and write with multiple goroutines.
 */

func operation(channel chan string) {
	//Decreasing the counter. This must be done at the exit of goroutine. Using deferred call,
	//we make sure that it will be called whenever function ends no matter but no matter how it ends.
	defer wg.Done()
	arr := [...]string{"Azures", "Panna", "Simca", "Faucon"}
	for i := range arr {
		channel <- arr[i]
	}
	close(channel)
}

func printChannel(channel chan string) {
	defer wg.Done()
	for i := range channel {
		fmt.Println(i)
	}
}

func operation1(channel chan string) {
	//Decreasing the counter. This must be done at the exit of goroutine. Using deferred call,
	//we make sure that it will be called whenever function ends no matter but no matter how it ends.
	defer wg.Done()
	arr := [...]string{"Mylene", "ab", "my", "caroline"}
	for i := range arr {
		channel <- arr[i]
	}
	close(channel)
}

func main() {
	myChannel := make(chan string)
	myChannel1 := make(chan string)
	wg.Add(4) //Increasing the counter. This must be done in main goroutine because there is no guarantee that newly started goroutine
	go operation(myChannel)
	printChannel(myChannel)
	go operation1(myChannel1)
	printChannel(myChannel1)
	wg.Wait() // Waiting for the counter to reach 0. This must be done in main goroutine to prevent program exit.
}
