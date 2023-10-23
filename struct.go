package main

import (
	"fmt"
)

type player struct {
	firstname string
	lastname  string
	score     int
	any       bool
}

type game struct {
	players map[string]*player
}

func (g *game) getWinner() *player {
	var maxScore int
	var winner *player

	for _, value := range g.players {
		if value.score > maxScore {
			winner = value
			maxScore = value.score
		}
	}
	return winner
}

type chain struct {
	base
	value int
	next  *chain
}

func (c *chain) sayOK() {
	fmt.Println("Chain OK")
}

type base struct {
}

func (b *base) sayHi() {
	fmt.Println("base hiii")
}

func (c *base) sayOK() {
	fmt.Println("Chain OK")
}

func gostruct() {
	// g := &game{
	// 	players: make(map[string]*player),
	// }
	// p1 := player{}
	// g.players["p1"] = &p1

	// var p2 player
	// p2 = player{}
	// p2.score = 4
	// g.players["p2"] = &p2

	// p3 := player{
	// 	lastname:  "mon",
	// 	firstname: "su",
	// }
	// g.players["p3"] = &p3

	// fmt.Println("game", g)

	// winner := g.getWinner()
	// fmt.Println("winner", winner)
	// fmt.Println("*winner", *winner)

	// b := base{}
	// b.sayHi()
	// b.sayOK()

	c1 := chain{value: 100}
	fmt.Println("c1", c1)

	c1.sayHi()

	c2 := &chain{value: 200}
	c1.next = c2
	fmt.Println("c1", c1)

	c1.next.sayHi()
	c1.sayOK()
}
