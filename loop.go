package main

import (
	"fmt"
)

func goloop() {
	// loop
	// for i := 0; i <= 4; i++ {
	// 	fmt.Println("hi", i)
	// }

	x := 1 
	for x <= 5 {
		fmt.Println(x)
		x++
	}

	for {
		fmt.Println("will run one time")
		return
	}
	fmt.Println("it will never run because of return")
}