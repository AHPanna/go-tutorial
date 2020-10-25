package main

import (
	"fmt"
)

/**
 * primitives types
 *
 *
 */
func main() {
	var n bool = true
	var m bool = false
	fmt.Printf("%v , %T\n", n, n)
	fmt.Printf("%v , %T\n", m, m)

	x := 1 == 1
	y := 1 == 2
	fmt.Printf("%v , %T\n", x, x)
	fmt.Printf("%v , %T\n", y, y)

	var empty bool // init bool var as false
	fmt.Printf("%v , %T\n", empty, empty)

	// big boy numbers
	var bigBoy int16 = 1286
	fmt.Printf("%v , %T\n", bigBoy, bigBoy)

	// diffrent hoods
	var bigHood2 int16 = 1286
	var bigHood1 int8 = 3
	// fmt.Println(bigHood1 + bigHood2) // --> error compile wrong hoood bro |fix : invite the little hood man to the big hood
	fmt.Println(int16(bigHood1) + bigHood2)

	// there are bits calculator, complexe, big numbers you can see this later
	// strings
	s := "My Name is Cartman"
	s1 := "and respect my authoritah"
	fmt.Printf("%v , %T\n", s, s)
	fmt.Printf("%v , %T\n", string(s[3]), s[3])
	fmt.Printf("%v , %T\n", s+", "+s1, s+", "+s1)

	// get all bytes printed out:
	b := []byte(s)
	fmt.Printf("%v , %T\n", b, b)

}
