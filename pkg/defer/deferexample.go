package main

import (
	"fmt"
	"os"
)

func first() {
	fmt.Println("1st")
}
func second() {
	fmt.Println("2nd")
}
func main() {

	// defer until the end of main
	defer second()
	first()

	f, _ := os.Open(".")

	if f == nil {
		fmt.Println("file is nil")
	} else {
		// hey yo, defer until the end of the code block
		defer f.Close()
		fmt.Println(f.Name())
		fmt.Println(f.Readdirnames(10))
	}
}
