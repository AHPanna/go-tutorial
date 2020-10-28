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
	aDoctor := Doctor{ // you can leave the key mapping but it's recommended for readability
		number:    69,
		actorName: "Azures",
		companions: []string{
			"Mylene",
			"Caroline",
			"Ab",
			"Mai",
		},
	}
	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)
	fmt.Println(aDoctor.companions[0])

	// volatile struct short live. -- not recommended
	aShortLiveStruct := struct{ name string }{name: "Azures"}
	fmt.Println(aShortLiveStruct)
	anotherShortLiveStruct := aShortLiveStruct
	anotherShortLiveStruct.name = "Mylene"
	fmt.Println(anotherShortLiveStruct)

	// advanced struct inheritance
	wildBird := Bird{}
	wildBird.Name = "Abid"
	wildBird.Origin = "Mars"
	wildBird.speed = 300.5
	wildBird.canFly = true
	fmt.Println(wildBird)
	fmt.Println(wildBird.Name)
	// reommended
	aBird := Bird{
		Animal: Animal{Name: "Nusrat", Origin: "Asia"},
		speed:  24.2,
		canFly: false,
	}
	fmt.Println(aBird)
	fmt.Println(aBird.Name)
}

// struct can only be declared outside of a function
type Doctor struct {
	number     int
	actorName  string
	companions []string
}

// advanced struct inheritance
type Animal struct {
	Name   string
	Origin string
}
type Bird struct {
	// let's add instance of animal
	Animal
	speed  float32
	canFly bool
}
