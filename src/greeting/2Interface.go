package greeting

import (
	"fmt"
)

/**
example of greeting showing a custom type with a def func and interface
 */

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

func (message Salutation) Greet(do Printer, isFormal bool) {
	greeting, altgreeting := CreateMessage(message.Greeting, message.Name, "Vic")

	if title := "Mr. "; isFormal {
		do(title + greeting)
	} else {
		do(altgreeting)
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

