package main

// refactor into scenes
// globals are inventory possibilities or knowledge - booleans?
// globals open new options
// one scene calls next scene depending on choice
// locations
// tracking state of users
// struct for inventory

import (
	"fmt"
)

type Player struct {
	Name           string
	Inventory      Backpack
	Traits         Character
	KnownLocations Atlas
}
type Backpack struct {
	Weapon string
}

// type Backpack struct {
// 	Weapons []string
// }

// func (b *Backpack) Acquire(weapon string) {
// 	b.Weapons = append(b.Weapons, weapon)
//example - player1.Inventory.Acquire("dk")
// }

type Character struct {
	Brave bool
	Tired bool
}

type Atlas struct {
	Locations []string
}

func (a *Atlas) Discover(location string) {
	a.Locations = append(a.Locations, location)
	//example - player1.KnownLocations.Discover("town")
}

var player1 Player

func main() {

	fmt.Println("Choose your own adventure - The Knight and the Dragon")
	fmt.Println()
	createChar()
	fmt.Println(player1)

	questBegin()

	fmt.Println()

	//showP1()

}

func createChar() {
	player1.Name = getInput("What is your name")
}

func questBegin() {
	fmt.Println("===========================================")
	fmt.Println()

	fmt.Printf("Welcome, brave %v!\n", player1.Name)
	fmt.Println("On this journey, you will face many choices.")
	fmt.Println("Sometimes choices will appear to you 'LIKE THIS', and you are to respond 'like this'.")
	fmt.Println("Other times it will be a simple yes or no question.")
	fmt.Println("Now, be off! Your quest awaits!")
	fmt.Println()
	castle()

}
