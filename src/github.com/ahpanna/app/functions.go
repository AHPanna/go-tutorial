package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 2; i++ {
		customFunc("Salut : ", i)
	}
	hello := "hello"
	name := "azures"
	changeName(&hello, &name)
	fmt.Println("Name changed : ", name)

	// return function
	fmt.Println("the sum is : ", *sum(1, 65, 36, 85, 69))

	// return with errors :
	res, err := errorCheck("hello", "ashd")
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(res)

	// annonymous function and isolated function
	func() {
		fmt.Println("Hello boi i am from lousiana")
	}() // the () is for invoking the function

	// sync function isolated
	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println("Hello boi i am from lousiana from : ", i)
		}(i) // passing the param inside of function for best practice
	}

	// calling by var a function
	f := func() {
		fmt.Println("i am a function inside of variable")
	} // no () for calling
	f()

	// function methods type:
	p := &Person{"azures", 21}
	//p = p.withName("Azures").withAge(21)
	fmt.Println(*p)
}

func customFunc(str string, i int) {
	fmt.Println(str, i)
}

// pointers
func changeName(name, newName *string) { // you can pass multiple varlues param if type is the same, and giving them a pointers
	fmt.Println(*name, *newName)
	*newName = "Simca"
	fmt.Println(*newName)
}

func sum(values ...int) *int { // first 'int'-> give as slice, the second 'int' return value as int
	fmt.Println(values)
	res := 0
	for _, v := range values {
		res += v
	}
	return &res
}

func errorCheck(a, b string) (string, error) { // the second parenthese show what values to return a the end
	if a == "hello" {
		return "Hello", nil
	} else {
		return "Wtf", fmt.Errorf("Wtf bro?")
	}
}

type Person struct {
	Name string
	Age  int
}

func (p *Person) withName(name string) *Person {
	p.Name = name
	return p
}

func (p *Person) withAge(age int) *Person {
	p.Age = age
	return p
}
