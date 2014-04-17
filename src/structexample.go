package main

import (
	"fmt"
	"math"
)

type Circle struct {
	x, y, r float64
}

// Circle here is a receiver

func (c *Circle) area() float64 {
     return math.Pi * c.r*c.r
}

func main() {

	//	c := new(Circle)
	c := Circle {1, 2, 5}

	fmt.Println(c.area())


}
