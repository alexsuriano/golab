package main

import (
	"fmt"

	"github.com/alexsuriano/golab/design-patterns/creational/singleton/single"
)

func main() {

	s := single.GetInstance()
	s.Name = "My single"
	s.Number = 3

	fmt.Printf("Name: %s, Number: %d\n\n", s.Name, s.Number)

	s2 := single.GetInstance()

	fmt.Printf("This Name: %s and this number: %d is of S2 variable\n\n", s2.Name, s2.Number)
}
