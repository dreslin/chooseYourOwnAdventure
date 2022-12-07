package main

import (
	"fmt"
)

func chooseDestination(CurrentLocation string) {
	fmt.Println("ChooseDest is called, and available choices shown")
	showAtlas(CurrentLocation)
	whereto := getInput("Where do you want to go or would you like to explore the area")
	toCheck = whereto
	result := isAvailable(player1.KnownLocations.Locations, toCheck)
	if result {
		locale[whereto]()
	} else {
		fmt.Println("you don't know that place yet")
		fmt.Println("next function called: chooseDestination")

		chooseDestination(player1.CurrentLocation.Location)
	}
}

// func createMap() {
// 	locale = make(map[string]func())
// 	locale["home"] = func() { home() }
// 	locale["cave"] = func() { cave() }
// 	locale["cave interior"] = func() { caveInterior() }
// 	locale["eastern wasteland"] = func() { easternwasteland() }
// 	locale["explore"] = func() { explore() }
// 	locale["cave entrance"] = func() { caveEntrance() }
// 	locale["crossroads"] = func() { crossroads() }
// 	locale["village"] = func() { village() }
// 	locale["hamlet"] = func() { hamlet() }
// locale["hamlet"] = func() { hamlet() }
// locale["hamlet"] = func() { hamlet() }
// locale["hamlet"] = func() { hamlet() }
// locale["hamlet"] = func() { hamlet() }
// locale["hamlet"] = func() { hamlet() }
// locale["hamlet"] = func() { hamlet() }
// locale["hamlet"] = func() { hamlet() }
// locale["hamlet"] = func() { hamlet() }

//}
