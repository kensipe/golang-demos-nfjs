package main

import "fmt"

func main() {

	x := 1
	y := 2

	fmt.Println("x,y", x, y)
	fmt.Println("swapped")
	swap(&x, &y)
	fmt.Println("x,y", x, y)

}

func swap(x *int, y *int) {
//	temp := *x
//	*x = *y
//	*y = temp
	*x, *y = *y, *x
}

//(x := 1; y := 2; swap(&x, &y)
