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

var (
	blksmtOffer string = "ignored"
	dk          string = "ignored"
	brave       string = "ignored"
	giveAid     string = "ignored"
)

type Player struct {
	Name      string
	Inventory Backpack
	Traits    Character
}
type Backpack struct {
	Weapon string
}
type Character struct {
	Brave bool
	Tired bool
}

var player1 Player

func main() {
	fmt.Println("Choose your own adventure - The Knight and the Dragon")
	fmt.Println()
	createChar()
	questBegin()

	fmt.Println()

	showP1()

}

func createChar() {
	player1.Name = getInput("What is your name?")
	fmt.Printf("Greetings, %v\n", player1.Name)
	fmt.Println()

}

func questBegin() {
	fmt.Println("===========================================")
	fmt.Println()
	fmt.Println("Welcome, brave adventurer!")
	fmt.Println("On this journey, you will face many choices.")
	fmt.Println("Sometimes choices will appear to you 'LIKE THIS', and you are to respond 'like this'.")
	fmt.Println("Other times it will be a simple yes or no question.")
	fmt.Println("And now, be off! Your quest awaits!")
	fmt.Println()
	castle()

}
