package main

import (
	"fmt"
)

func gorange() {
	// range on array
	arr := [4]string{"A", "B", "C", "D"}
	for k, v := range arr {
		fmt.Println("range on arr key:", k, "val:", v)
	}

	// range on slice
	slc := []string{"A", "B", "C"}
	for k, v := range slc {
		fmt.Println("range on slc key:", k, "val:", v)
	}

	// range on map
	mp := map[string]string{
		"admin": "Ahmed",
		"user":  "Sara",
	}
	for key, value := range mp {
		fmt.Println("range on map key:", key, "value:", value)
	}

	// range on map of slices
	mapOfSlices := map[string][]string{
		"Admins": []string{"Ahmed", "Jack"},
		"Users":  []string{"X", "Y", "Z"},
	}
	for key, value := range mapOfSlices {
		fmt.Println("range on map of slices key:", key, "value:", value)
		for k, v := range value {
			fmt.Println("range on slices in map key:", k, "value:", v)
		}
	}
}
