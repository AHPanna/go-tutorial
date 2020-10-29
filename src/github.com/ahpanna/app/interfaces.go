package main

import (
	"fmt"
)

func main() {
	var w Writer = &ConsoleWrite{}
	o, err := w.Write([]byte("Hello Mylene"))
	if err != nil {
		fmt.Println("Something went wrong : ", err)
	}
	fmt.Printf(" Results : %v\n", o)
}

// interface is behaviour
type Writer interface {
	Write([]byte) (int, error)
}

type ConsoleWrite struct{}

func (cw *ConsoleWrite) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
