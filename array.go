package main

import (
	"fmt"
)

func goarray() {
	// define arrays
	var arr [4]string
	fmt.Println("arr", arr)

	var a [3]int
	fmt.Println("int arr", a)

	arr[0] = "Ahmed"
	arr[3] = "John"
	fmt.Println("arr", arr)

	// reassign val in arr
	arr = [4]string{"jan", "feb", "mar", "apr",}
	fmt.Println("arr", arr)

	var num = [...]int{1, 2, 3, 4, 5}
	fmt.Println("num array", num)

	arrLen := len(arr)
	fmt.Println(arrLen)
}
