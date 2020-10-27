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

	a := [...]int{1, 2, 43, 10} // array init
	b := a                      // copy array
	fmt.Println(a)
	fmt.Println(b[2])

	// lets slice. -> this is usefull when working for data.
	slice := []int{65, 3265, 659, 897, 989, 324, 7898}
	fmt.Printf("Results : %v\n", slice)
	fmt.Printf("Length : %v\n", len(slice))
	fmt.Printf("Capacity : %v\n", cap(slice))

	/**
	 * Note : Slice and Array
	 * Array = [...]int{data,dateaz,fdazd}
	 * Slice = []int{data,dateaz,fdazd}
	 * At end if you print it's same results but different usage operation,
	 * the difference can be revealed by those `...` for array inside of the brackets.
	 * */

	// custom slice
	custSlice := []int{546, 65464, 68489, 87987, 32135, 54683, 468746, 1354856, 212, 345, 687}
	customSlice1 := custSlice[:]   // slice of all elements
	customSlice2 := custSlice[2:]  // slcie from 3rd elements to end
	customSlice3 := custSlice[:6]  // slice first 5th elements to end
	customSlice4 := custSlice[4:7] // slice the 5th, 6th, 7th elements

	//let's change some values that happened later some operations
	custSlice[5] = 69

	fmt.Println(custSlice)
	fmt.Println(customSlice1)
	fmt.Println(customSlice2)
	fmt.Println(customSlice3)
	fmt.Println(customSlice4)

	// new slice methods
	newSlice := make([]int, 10, 15) //type of instance, length and capacity > lenght
	fmt.Println(newSlice)           // should be only `0` as instance initiated.
	fmt.Printf("Lenght : %v\n", len(newSlice))
	fmt.Printf("Capacity : %v\n", cap(newSlice))

	// manual slice append capacity :
	manuelSlice := []int{}   // this instance is initiated as empty
	fmt.Println(manuelSlice) // should be only `0` as instance initiated.
	fmt.Printf("Init Lenght : %v\n", len(manuelSlice))
	fmt.Printf("Init Capacity : %v\n", cap(manuelSlice))

	manuelSlice = append(manuelSlice, 10) // adding first element in slice
	fmt.Printf("After Lenght : %v\n", len(manuelSlice))
	fmt.Printf("After Capacity : %v\n", cap(manuelSlice))

	// adding more elements inside slice
	manuelSlice = append(manuelSlice, 10, 3265, 569, 68486, 656, 32325, 648, 6564)
	fmt.Printf("After Lenght : %v\n", len(manuelSlice))
	fmt.Printf("After Capacity : %v\n", cap(manuelSlice)) // old capacity is still kept in memory

	// adding more elements inside slice via append of slice int
	manuelSlice = append(manuelSlice, []int{0, 25, 36, 48, 69, 788, 569, 577}...) // the ... will send each each number inside of slice, works only for slice.
	fmt.Printf("After Lenght : %v\n", len(manuelSlice))
	fmt.Printf("After Capacity : %v\n", cap(manuelSlice)) // old capacity is still kept in memory

	// stack operation slice
	stackSlice := []int{154, 335, 3658, 6985, 52}
	fmt.Printf("Original : %v\n", stackSlice)
	stackSlice1 := stackSlice[1:]
	fmt.Printf("No first Slice : %v\n", stackSlice1)

	// lets not get the end.
	stackSlice2 := stackSlice[:len(stackSlice)-1]
	fmt.Printf("No last Slice : %v\n", stackSlice2)

	// lets not get middle slice.
	stackSlice3 := append(stackSlice[:2], stackSlice[3:]...)
	fmt.Printf("No Middle Slice : %v\n", stackSlice3)
	// Wow?
	fmt.Printf("Original : %v\n", stackSlice) // make sure to copy a temp slice if you don't want to lose data on slices.
	// continues here https://youtu.be/YS4e4q9oBaU?t=7018
}
