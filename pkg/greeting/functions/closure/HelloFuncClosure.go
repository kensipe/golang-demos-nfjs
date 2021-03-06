package main

import (
	"fmt"
)

type Salutation struct {
	name     string
	greeting string
}

type Printer func(string) ()

func CreatePrintFunction(custom string) Printer {

	return func(s string) {
		fmt.Println(s + custom)
	}
}


func Greet(message Salutation, do Printer) {
	greeting, altgreeting := CreateMessage(message.greeting, message.name, "Sarah")

	do(greeting)
	do(altgreeting)
}

func CreateMessage(greeting string, name ...string) (message string, altgreeting string) {
	fmt.Println(len(name))

	message = greeting+" "+name[0]
	altgreeting = "hey! "+name[1]
	return
}

func Print(s string) {
	fmt.Print(s)
}

func PrintLine(s string) {
	fmt.Println(s)
}

func main() {

	var message = Salutation{}
	message.name = "ken"
	message.greeting = "hello"

	Greet(message, CreatePrintFunction("!!!"))

}

