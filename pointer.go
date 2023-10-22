package main

import (
	"fmt"
)

func gopointer() {
	name := "Ahmed"
	fmt.Println(name)

	// sending copy of variable value
	updateCopy(name)
	fmt.Println(name)

	// sending pointer to the variable by adding &
	updateReference(&name)
	fmt.Println(name)

	// print a pointer will print the memory address
	fmt.Println("will print the memory address", &name)

	// if parameter is pointer we can send nil as argument
	defaultvaluepointer(nil)

	var Pvalue *string
	Pvalue = &name
	fmt.Println("pointer", *Pvalue)

}

func updateCopy(s string) {
	s = "mark"
}

func updateReference(s *string) {
	*s = "mark s"
	n := "name"
	s = &n
	fmt.Println("up fun n:", &n, "s:", s)
}

func defaultvaluepointer(s *string) {
	fmt.Println("nil val", s)
}
