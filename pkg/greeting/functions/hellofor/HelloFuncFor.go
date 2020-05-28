package main

import (
	"fmt"
)

type Salutation struct {
	Name     string
	Greeting string
}

type Printer func(string) ()

func CreatePrintFunction(custom string) Printer {

	return func(s string) {
		fmt.Println(s + custom)
	}
}

type Renamable interface {
	Rename(newName string)
}

func RenameToJoe(r Renamable) {
	r.Rename("Joe")
}

func (message *Salutation) Rename(newName string) {
	message.Name = newName
}

func (message Salutation) Greet(do Printer, isFormal bool, times int) {

	greeting, altgreeting := CreateMessage(message.Greeting, message.Name, "Sarah")


//	only loops in Go are for... this fo
	for i := 0; i < times; i ++ {

		if title := GetPrefix(message.Name); isFormal {
			do(title + greeting)
		} else {
			do(altgreeting)
		}
	}
}

func GetPrefix(name string) (prefix string) {

	switch name {
	case "Bob" : prefix = "Mr. "
	case "Joe", "Amy" : prefix = "Dr. "
	case "Mary" : prefix = "Mrs. "
	default : prefix = "Dude "
	}

	return
}

func TypeSwitch(x interface{}) {
	switch x.(type) {
	case int:
		fmt.Println("int")
	case string:
		fmt.Println("int")
	case Salutation:
		fmt.Println("Salutation")
	default:
		fmt.Println("unknown")
	}
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

	var message = Salutation{ "Bobby", "hello" }

	message.Greet(CreatePrintFunction("!!!"), true, 5)

}

