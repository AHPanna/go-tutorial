package main

import (
	"fmt"
)

func main() {
	var a int = 42
	var b *int = &a // b pointing to a address contains values
	fmt.Println(a, b)
	fmt.Println(&a, b) // same addresse
	fmt.Println(a, *b) // print the value of b pointer to address
	*b = 69            // let's change the value of a
	fmt.Println(a, *b) // print the value of b pointer to address

	// array pointers
	c := []int{15, 266, 36, 9, 3}
	d := &c[2] // address pointing to the c array of 2 attributed to d
	e := &c[1] // address pointing to the c array of 1 attributed to d
	fmt.Printf("%v %p %p\n", c, d, e)
	var ms *myStruct
	ms = new(myStruct) // new pointing the myStruct address
	fmt.Println(ms)    // you should see '&' for pointing address
	ms.foo = 69
	fmt.Println(ms.foo) // derreference pointer

	// array value wont change and slice have pointers and map
	arr := [3]int{1, 2, 3}
	arrb := arr
	fmt.Println(arr, arrb)
	arr[1] = 69
	fmt.Println(arr, arrb) // this kept the original values and show new one

	//slice
	arr1 := []int{1, 2, 3}
	arrb1 := arr1
	fmt.Println(arr1, arrb1)
	arr1[1] = 69
	fmt.Println(arr1, arrb1) // the value is changes for both original and new, because of pointer

	// map
	myMap := map[string]string{"name": "panna", "name1": "azures", "name3": "simca"}
	myMapb := myMap
	fmt.Println(myMap, myMapb)
	myMapb["name3"] = "mylene"
	fmt.Println(myMap, myMapb) // the value is changes for both original and new, because of pointer
}

type myStruct struct {
	foo int
}
