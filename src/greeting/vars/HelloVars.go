package main

import (
	"fmt"
)

type Salutation struct {
	name     string
	greeting string
}

const (
	PI       = 3.14
	Language = "Go"
)

const (
	First  = iota
	Second = iota
	Third  = iota
)

func main() {

	var message = Salutation{}
	message.name = "ken"
	message.greeting = "hello"

	fmt.Println(message.greeting, message.name)
	fmt.Println(First, Second, Third)

}

