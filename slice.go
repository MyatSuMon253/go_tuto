package main

import (
	"fmt"
)

func slice() {
	// init slice with initial values
	src := []string{"x", "y", "z"}
	src = append(src, "d")

	// another way to init slice with predefined size
	dest := make([]string, 2)

	// copy
	copy(dest, src)
	fmt.Println(dest)

	// change source after copy
	src[0] = "a"
	fmt.Println(src)
	fmt.Println(dest)

	// a slice point to same values in original array
	// start: end-1
	s := src[1:3]
	fmt.Println(s)

	// updating the source
	src[2] = "c"
	fmt.Println(src)
	fmt.Println(s)

	// src[:x] index 0 to x-1
	// src[x:y] index x to y-1
	// src[x:] index x to len -1
}
