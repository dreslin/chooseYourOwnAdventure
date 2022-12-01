package main

import (
	"fmt"
)

func chooseDestination() {
	fmt.Println("ChooseDest is called, and available choices shown")
	showAtlas()
	whereto := getInput("Where do you want to go or would you like to explore the area")
	// result := isAvailable(player1.KnownLocations.Locations, whereto)
	toCheck = whereto
	result := isAvailable(player1.KnownLocations.Locations, toCheck)
	if result {
		locale[whereto]()
	} else {
		fmt.Println("you don't know that place yet")
		chooseDestination()
	}
}

func createMap() {
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
	// locale["hamlet"] = func() { hamlet() }
	// locale["hamlet"] = func() { hamlet() }
	// locale["hamlet"] = func() { hamlet() }
	// locale["hamlet"] = func() { hamlet() }
	// locale["hamlet"] = func() { hamlet() }
	// locale["hamlet"] = func() { hamlet() }
	// locale["hamlet"] = func() { hamlet() }
	// locale["hamlet"] = func() { hamlet() }

}
