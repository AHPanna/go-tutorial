package main

import (
	"fmt"
)

func main() {

	/**
	 * the map is a key value format but it's fixed types;
	 * which is most part is retarded only java retards use this.
	 *
	 * Note this the order of map is random so becareful, which more retarded
	 * */

	myMap := map[string]int{
		"Chambre": 3,
		"Age":     24,
		"Rue":     10,
		"Family":  3,
	}
	fmt.Println(myMap)
	// let's add another map into our current map :
	customMap := map[[3]int]string{} // created 3 key with empty values of string
	fmt.Println(myMap, customMap)

	// create map with instance
	instMap := make(map[string]int)
	fmt.Println(myMap, instMap) // best practice

	// retrieve key value
	fmt.Println(myMap["Age"])
	// edit value :
	myMap["Family"] = 5
	fmt.Println(myMap)
	// delete key
	delete(myMap, "Chambre")
	fmt.Println(myMap)

	// let's check for key
	valueExist, isExist := myMap["Rue"]
	fmt.Println(valueExist, isExist)

	valueExist1, isExist1 := myMap["Chambre"]
	fmt.Println(valueExist1, isExist1)

	// struct
	// 	https: //youtu.be/YS4e4q9oBaU?t=8913

}
