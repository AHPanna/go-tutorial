package main

import (
	"fmt"
	"time"
)

func display(str string) {
	for w := 0; w < 6; w++ {
		time.Sleep(1 * time.Second)
		fmt.Println(str)
	}
}

func main() {

	// Calling Goroutine
	go display("Welcome")

	// Calling normal function
	display("Azures S/U/C")

	fmt.Println("Welcome!! to Main function")

	// Creating Anonymous Goroutine
	go func() {

		fmt.Println("Welcome!! to Azures")
	}()

	time.Sleep(1 * time.Second)
	fmt.Println("GoodBye!! to Main function")
}

//https://www.geeksforgeeks.org/select-statement-in-go-language/?ref=lbp
//https://www.geeksforgeeks.org/multiple-goroutines/?ref=lbp
//https://www.geeksforgeeks.org/channel-in-golang/?ref=lbp
//https://www.geeksforgeeks.org/unidirectional-channel-in-golang/?ref=lbp
