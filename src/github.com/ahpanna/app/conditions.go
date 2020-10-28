package main

import (
	"fmt"
)

func main() {
	if true {
		fmt.Println("Hello from true")
	}

	// little game basic conditions
	number := 50
	guess := 70
	if guess < 1 || guess > 100 {
		fmt.Println("Out of bound only 1-100")
	} else {
		if guess < number {
			fmt.Println("too low")
		} else if guess > number {
			fmt.Println("too high")
		} else if guess == number {
			fmt.Println("Yes !!!")
		}
	}
	fmt.Println(number <= guess, number >= guess, number != guess)

	// swithc case  --> no need a break or fallthrough
	victim := "glass"
	switch victim {
	case "Apple", "Car":
		fmt.Println("i should take some drugs")
	case "pencil", "glass", "cow":
		fmt.Println("i know the suspect is the cow")
	default:
		fmt.Println("you are the killer, arrest yourself.")
	}

	// type switch
	var i interface{} = [3]int{}
	switch i.(type) {
	case int:
		fmt.Println("it's int")
	case float32:
		fmt.Println("it's float")
	case string:
		fmt.Println("it's string")
	case [2]int:
		fmt.Println("it's array")
	default:
		fmt.Println("i want to believe")
		//break
		fmt.Println("might be a ufobool")

	}
}
