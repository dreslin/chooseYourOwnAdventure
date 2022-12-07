package main

import (
	"fmt"
	"math/rand"
	"time"
)

type Player struct {
	Name            string
	Inventory       Backpack
	Traits          Character
	KnownLocations  Atlas
	CurrentLocation Atlas
}
type Backpack struct {
	Weapons []string
}

func (b *Backpack) Acquire(weapon string) {
	b.Weapons = append(b.Weapons, weapon)
	// example - player1.Inventory.Acquire("sword")
}

type Character struct {
	Brave bool
	Tired bool
}

type Atlas struct {
	Locations []string
	Location  string
}

func (a *Atlas) Discover(location string) {
	a.Locations = append(a.Locations, location)
	//example - player1.KnownLocations.Discover("town")
}

var player1 Player
var locale map[string]func()
var toCheck string

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	locale = make(map[string]func())
	locale["home"] = func() { home() }
	locale["cave"] = func() { cave() }
	locale["cave interior"] = func() { caveInterior() }
	locale["eastern wasteland"] = func() { easternwasteland() }
	locale["explore"] = func() { explore() }
	locale["cave entrance"] = func() { caveEntrance() }
	locale["crossroads"] = func() { crossroads() }
	locale["village"] = func() { village() }
	locale["hamlet"] = func() { hamlet() }

	fmt.Println("Choose your own adventure - The Knight and the Dragon")
	fmt.Println()
	createChar()
	questBegin()
	fmt.Println()

	//showP1()

	// fmt.Println("showing atlas before adding new areas")
	// showAtlas()
	// player1.KnownLocations.Discover("eastern wasteland")
	// player1.KnownLocations.Discover("cave")
	// player1.KnownLocations.Discover("hamlet")
	// player1.KnownLocations.Discover("village")

	// fmt.Println("showing atlas after adding new areas")
	// showAtlas()
	// chooseDestination()

}

func createChar() {

	player1.Name = getInput("What is your name")
	player1.KnownLocations.Discover("home")
	player1.KnownLocations.Discover("explore")
	player1.Inventory.Acquire("sword")
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
	fmt.Println("next function called: castle")

	castle()

}
