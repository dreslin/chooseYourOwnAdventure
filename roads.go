package main

import (
	"fmt"
	"math/rand"
	"time"
)

func roadEvent() {
	// randomly chooses road
	event := eventSelect()
	switch event {
	case 1:
		goblins()
	case 2:
		trolls()
	case 3:
		elves()
	case 4:
		dwarves()
	case 5:
		unicorn()
	}
}

func eventSelect() int {
	var v [5]int
	rand.Seed(time.Now().UnixNano())
	for i := 0; i < 5; i++ {
		v[i] = rand.Intn(5) + 1
	}
	return v[1]

}

func goblins() { fmt.Println("This is a random road event with goblins.") }
func trolls()  { fmt.Println("This is a random road event with trolls.") }
func elves()   { fmt.Println("This is a random road event with elves.") }
func dwarves() { fmt.Println("This is a random road event with dwarves.") }
func unicorn() { fmt.Println("This is a random road event with a unicorn.") }
