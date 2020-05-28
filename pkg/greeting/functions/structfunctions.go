package main

import "fmt"

func main() {
// what would a functional example be without factorials?
	example3()
}

func example1() {
	add := func(x, y int) int {
		return x + y
	}

	fmt.Println(add(1, 1))
}

func example2() {
	x := 0
	increment := func() int {
		x++
		return x
	}
	fmt.Println(increment())
	fmt.Println(increment())
}

func example3() {
	fmt.Println(factorial(uint(2)))
	fmt.Println(factorial(uint(4)))
}

func factorial(x uint) uint {
	if x == 0 {
		return 1
	}
	return x * factorial(x - 1)
}
