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

func Greet(message []Salutation, do Printer, isFormal bool, times int) {

	for _, s := range message {

		greeting, altgreeting := CreateMessage(s.Greeting, s.Name, "Sarah")
		if title := GetPrefix(s.Name); isFormal {
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

func CreateMessage(greeting string, name ...string) (message string, altgreeting string) {

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

	slice := []Salutation {
		{ "Bob", "Hello"},
		{ "Joe" , "Hi"},
		{"Mary", "What's up"},
	}

	Greet(slice, CreatePrintFunction("!!!"), true, 5)

}

