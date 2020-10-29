package main

import (
	"fmt"
)

func main() {
	// let's visite the house
	var visiteur house = &visitor{}
	fmt.Println("Nombre de Chambre : ", visiteur.Rooms(5))
	fmt.Println("Personne : ", visiteur.Person("Azures"))
}

// create interface
type house interface {
	// methods available for that house interface
	Rooms(int) int
	Person(string) string
}

type visitor struct {
	person int
	name   string
}

func (m *visitor) Rooms(n int) int {
	// only 1 person or add param :)
	m.person = n
	return m.person
}

func (m *visitor) Person(str string) string {
	m.name = str
	return m.name
}
