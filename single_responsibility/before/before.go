// Single Responsibility Principle
// A class should have one and only one reason to change.
// In Go: a type of a function should have one and only one job responsibility

package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

func (c circle) area() {
	// against single responsibility principle
	// it does two things, output result and calculate area
	fmt.Printf("circle area: %f \n", math.Pi * c.radius * c.radius)
}

type square struct {
	length float64
}

func (s square) area() {
	fmt.Printf("square area: %f \n", s.length * s.length)
}

func main() {
	c := circle{radius: 5}
	c.area()

	s := square{length: 7}
	s.area()
}
