package main

import (
	"fmt"
	"math"
)

type shape interface {
	area() float64
}

// Embedded shape interface as field which composited area() method
type object interface {
	shape
	volume() float64
}

// Function signature of areaSum only accepting type shape value
func areaSum(shapes ...shape) float64 {
	var sum float64
	for _, s := range shapes {
		sum += s.area()
	}
	return sum
}

// Function signature of areaVolumeSum only accepting type object value
func areaVolumeSum(shapes ...object) float64 {
	var sum float64
	for _, s := range shapes {
		sum += s.area() + s.volume()
	}
	return sum
}

type square struct {
	length float64
}

func (s square) area() float64 {
	return math.Pow(s.length, 2)
}

type cube struct {
	square
}

func (c cube) volume() float64 {
	return math.Pow(c.length, 3)
}

func main() {
	s1 := square{length: 5}
	s2 := square{length: 6}
	c1 := cube{square: square{length: 3}}
	c2 := cube{square: square{length: 4}}
	fmt.Println(areaSum(s1, s2, c1, c2))
	fmt.Println(areaVolumeSum(c1, c2))
}
