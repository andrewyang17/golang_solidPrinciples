// Liskov Substitution Principle
// Ability to replace any instance of a parent class
// with an instance of one of its child classes
// without negative side effects

// Human type is not the only one who implements the person interface
// Teacher and student types are also implement the person interface
// by inheriting (in Go called composition) every field and method of this human type.
// So the teacher type and student type is considered a subtype of the human type
// that they both can be substituted with the human type.

package main

import "fmt"

type human struct {
	name string
}

func (h human) getName() string {
	return h.name
}

type teacher struct {
	human
	degree string
	salary float64
}

type student struct {
	human
	grades map[string]int
}

type person interface {
	getName() string
}

type printer struct {}

func (printer) info(p person) {
	fmt.Println("Name: ", p.getName())
}

func main() {
	h := human{name: "Andrew"}
	s := student{
		human:  human{
			name: "Julius",
		},
		grades: map[string]int{
			"Math": 10,
			"English": 10,
		},
	}
	t := teacher{
		human:  human{
			name: "Jay",
		},
		degree: "Computer Science",
		salary: 5000,
	}

	p := printer{}
	p.info(h)
	p.info(s)
	p.info(t)
}
