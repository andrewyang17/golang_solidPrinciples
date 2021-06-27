// Open Closed Principle
// Objects/entities should be open for extension but closed for modification

package main

import (
	"fmt"
	"math"
)

type circle struct {
	radius float64
}

type square struct {
	length float64
}

type calculator struct {
}

// Function signature of areaSum takes empty interface value
func (a calculator) areaSum(shapes ...interface{}) float64 {
	var sum float64
	for _, shape := range shapes {
		switch shape.(type) {
		// Against open closed principle
		// If we add a new shape means we need to add new case
		// basically we need to modify this function
		case circle:
			r := shape.(circle).radius
			sum += math.Pi * r * r
		case square:
			l := shape.(square).length
			sum += l * l
		}
	}
	return sum
}

func main() {
	c := circle{radius: 5}
	s := square{length: 7}
	calc := calculator{}
	fmt.Println("area sum:", calc.areaSum(c, s))
}
