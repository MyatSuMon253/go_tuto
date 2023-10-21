package main

import (
	"fmt"
)

func gomap() {
	// init a map with initial values
	m := map[string]bool{
		"x": true,
		"y": false,
	}

	fmt.Println(m)

	// init a map, no need for length
	players := make(map[string]int)

	fmt.Println("p", players)

	// add val to the map
	players["Ahmed"] = 100
	players["Jack"] = 99
	players["Sara"] = 10

	// get data from map
	// fmt.Println("players", players["Sara"])

	// del val from map
	delete(players, "Jack")
	fmt.Println("p", players)

	// delete wrong key does not affect the map
	// delete(players, "Jack")
	// fmt.Println("p", players)

	//check if key exist and get val
	// value will be default if key does not exist
	v, ok := players["y"]
	fmt.Println("missing key, v", v, "ok", ok)

	v, ok = players["Sara"]
	fmt.Println("valid key, v", v, "ok", ok)

	// map of slices
	mp := map[string][]string{
		"admins": []string{"ahmed", "jack"},
		"users":  []string{"x", "y", "z"},
	}
	fmt.Println("mp", mp)
}
