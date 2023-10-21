package main

import (
	"fmt"
)

func main() {
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
}
