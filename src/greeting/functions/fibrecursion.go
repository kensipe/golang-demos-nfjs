package main

import "fmt"

func main() {

	fmt.Println("fib is", fib(7))

}

func fib(x int) int {

	switch x {
	case 0: return 0;
	case 1: return 1;
	default: return fib(x - 1) + fib(x - 2)
	}
}
