package main

import (
	"errors"
	"fmt"
)

// func that does not return anything
func p(s string) {
	fmt.Println("hello", s)
}

// func that does not return anything
// and take as many params as we send and print params length
func l(params ...interface{}) {
	fmt.Println(len(params))
}

// func that take 2 args and return the sum
func add(x int, y int) int {
	return x + y
}

// let's group types if it is the same
func sum(x, y int) int {
	return x + y
}

// return multiple results
func concat(a, b string) (string, int) {
	r := a + b
	return r, len(r)
}

// named return types
func concat2(a, b string) (r string, l int) {
	r = a + b
	l = len(r)
	// will return the named vars
	return
}

// func that return string & nil or default string value & error
// this pattern is widely used in Go
func returnError(b bool) (string, error) {
	if b {
		return "", errors.New("this is an go error")
	}
	return "okay", nil
}

func gofunc() {
	res, len := concat2("hi", "hey")
	fmt.Println(res, len)

	r, err := returnError(false) // false
	if err != nil {
		fmt.Println("we have error ", err.Error())
	} else {
		fmt.Println("we got ", r)
	}

	// func type
	var x func(int, int) int
	x = sum
	fmt.Println("x", x(3, 5))
}
