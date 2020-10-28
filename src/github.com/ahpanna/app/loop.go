package main

import (
	"fmt"
)

func main() {
	// only "for" loop
	// basic loop

	for i := 0; i < 5; i++ {
		fmt.Println("Hey this is  : ", i)
	}

	for j := 0; j < 5; j++ {
		fmt.Println("Hey this is  : ", j)
		if j%2 == 0 {
			j /= 2
		} else {
			j = 2*j + 1
		}
	}

	// fuck boi loop 肏你妈
	i := 0
	for ; i < 5; i++ {
		fmt.Println("this is  : ", i)
	}
	fmt.Println(i)

	// while type loop
	x := 0
	for x < 5 {
		fmt.Println("this is  : ", x)
		x++
	}

	// infinite loop
	z := 0
	for {
		fmt.Println(z)
		z++
		if z == 10 {
			break // exit the loop
		}
		if z%2 == 0 { // continues infite loop but controlled
			fmt.Println("true")
			continue // this will restard from the begaining after this no sequence will be executed. so put on order
		}
	}

	// loop inside of slice
	mySlice := []int{1321, 3265, 65, 135, 356, 9, 89, 5, 65}
	for k, v := range mySlice {
		fmt.Println("key : ", k, " Values : ", v)
	}
}
