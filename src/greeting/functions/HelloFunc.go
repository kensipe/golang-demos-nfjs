package main

import (
	"fmt"
)

type Salutation struct {
	Name     string
	Greeting string
}


// 2.
//func CreateMessage(greeting, name string) (string, string) {
//	return greeting + " " + name, "hey! " + name
//}

//3.
//func CreateMessage(greeting, name string) (message string, altgreeting string) {
//	message = greeting+" "+name
//	altgreeting = "hey! "+name
//	return
//}

func Greet(message Salutation) {
	_, altgreeting := CreateMessage(message.greeting, message.name, "Vic")

	fmt.Println(altgreeting)
}

func CreateMessage(greeting string, name ...string) (message string, altgreeting string) {
	fmt.Println(len(name))

	message = greeting+" "+name[0]
	altgreeting = "hey! "+name[1]
	return
}

// 1.
//func CreateMessage(greeting, name string) string {
//	return greeting + " " + name
//}
// 1.
//func Greet(message Salutation) {
//	fmt.Println(CreateMessage(message.greeting, message.name))
//}

func main() {

	var message = Salutation{}
	message.name = "ken"
	message.greeting = "hello"

	Greet(message)

}

