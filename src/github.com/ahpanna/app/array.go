package main

import (
	"fmt"
)

func main() {
	grades := [3]int{100, 20, 30}
	for i, k := range grades {
		fmt.Printf("Grade student %v : %v\n", i, k)
	}
	fmt.Printf("total student %v\n", len(grades))

	a := [...]int{1, 2, 43, 10}
	b := a // copy array
	fmt.Println(a)
	fmt.Println(b[2])

	// continues here https://youtu.be/YS4e4q9oBaU?t=7018
}
