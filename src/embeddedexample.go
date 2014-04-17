package main

import "fmt"

type Person struct {
	Name string
}

func (p *Person) Talk() {
	fmt.Println("Hello. My name is", p.Name)
}

type Android struct    {
	Person
	Model string
}

func main() {

	// structures and inheritance are different in Go...
/**
type Android struct    {
	Person Person
	Model string
}  would be a different ... the Person without a a var is treated different :)

	 */
	// we don't want a.Person.Talk()

	a := new(Android)
	a.Model = "XRT"
	a.Name = "Amanda"

	fmt.Println(a.Talk())
}

