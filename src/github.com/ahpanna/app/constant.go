package main

import (
	"fmt"
)

/**
 * Constants Bois
 * a cost which a value can't be changed
 */
func main() {

	const myConst int = 1
	// myConst = 23 // --> compiler boss will show you error
	fmt.Printf("%v , %T\n", myConst, myConst)

	const myConst1 float32 = 4.26
	fmt.Printf("%v , %T\n", myConst1, myConst1)

	fmt.Printf("%v , %T\n", myConst1+float32(myConst), myConst1+float32(myConst))

	// counter const
	const (
		a = iota
		b = iota
		c
	)

	const (
		_ = iota // the '_' is not kept in memory or ignore.
		theHomelessGuy
	)
	fmt.Printf("%v , %T\n", a, a)
	fmt.Printf("%v , %T\n", b, b)
	fmt.Printf("%v , %T\n", c, c)
	fmt.Printf("%v , %T\n", theHomelessGuy, theHomelessGuy)

	// look for bit shift ioat algorithme in one byte checkout 1h42 : https://www.youtube.com/watch?v=YS4e4q9oBaU

}
