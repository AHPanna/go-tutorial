package main

import (
	"fmt"
	"strconv"
)

// init global scope level declaration

var str1 string = "Hello"
var str2 string = "World"

// deaclare a packs of variables.
var (
	count    int    = 0
	season   int    = 11
	episode  string = "The death Curse"
	actorAge int    = 45
)

/* notiice you can reassigne variables
 * var i int = 0
 * i := "something" ->error compile
 * and you will get errors when some varibales are not used so go like to clean.
 * [NAMING] : package level global decalaration variable in uppercase can be used outside.
 * keep the variables length short, don't be that java retarded jerk.
 *
 */

var outsideExport int = 69

func variable() {
	var i int // init a var
	i = 10
	var j float32 = 42.69 // init with values
	x := "wassup"         // this is auto known by go for type
	fmt.Println(i)
	fmt.Println(j)
	fmt.Println(x)
	fmt.Println(str1)
	fmt.Println(str2)
	fmt.Println(count)
	fmt.Println(season)
	fmt.Println(episode)
	fmt.Println(actorAge)
	fmt.Println(outsideExport)
	fmt.Printf("%v, %T\n", i, i)
	fmt.Printf("%v, %T\n", j, j)
	fmt.Printf("%v, %T\n", x, x)

	// Conversion
	fmt.Println("#####Conversion#######")
	var varA float32 = 10.12
	fmt.Printf("%v, %T\n", varA, varA)
	var varB int
	varB = int(varA)
	fmt.Printf("%v, %T\n", varB, varB)
	var varC string
	varC = strconv.Itoa(i) // get ascii string of result and not the ascii table values.
	fmt.Printf("%v, %T\n", varC, varC)

}
