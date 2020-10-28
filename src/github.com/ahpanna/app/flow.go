package main

import (
	"fmt"
)

func main() {
	// defer view as await async
	fmt.Println("first")
	defer fmt.Println("i will at last in this order")
	fmt.Println("last")
	// defer opposite from the last to start
	defer fmt.Println("first")
	defer fmt.Println("middle 肏你妈")
	defer fmt.Println("last")

	//panic -- view as exception
	fmt.Println("Woiuyyyiioo")
	//panic("something going on are oyu okay annie?")
	defer fmt.Println("Smooth Criminal")
	fmt.Println("end")

	// recover
	// to long
	f()
	fmt.Println("Returned normally from f.")
}

func f() {
	defer func() {
		if r := recover(); r != nil {
			fmt.Println("Recovered in f", r)
		}
	}()
	fmt.Println("Calling g.")
	g(0)
	fmt.Println("Returned normally from g.")
}

func g(i int) {
	if i > 3 {
		fmt.Println("Panicking!")
		panic(fmt.Sprintf("%v", i))
	}
	defer fmt.Println("Defer in g", i)
	fmt.Println("Printing in g", i)
	g(i + 1)
}
