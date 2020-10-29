package main

import (
	"fmt"
)

func myfun(a interface{}) {

	// Extracting the value of a
	val := a.(string)
	fmt.Println("Value: ", val)
}

func myfun1(a interface{}) {
	value, ok := a.(float64)
	fmt.Println(value, ok)
}

func myfun2(a interface{}) {

	// Using type switch
	switch a.(type) {

	case int:
		fmt.Println("Type: int, Value:", a.(int))
	case string:
		fmt.Println("\nType: string, Value: ", a.(string))
	case float64:
		fmt.Println("\nType: float64, Value: ", a.(float64))
	default:
		fmt.Println("\nType not found")
	}
}

func main() {

	var val interface {
	} = "Azures"
	myfun(val)

	var a1 interface {
	} = 98.09
	myfun1(a1)
	var a2 interface {
	} = "Azures"
	myfun1(a2)

	myfun2("Mylene")
	myfun2(67.9)
	myfun2(true)
}
