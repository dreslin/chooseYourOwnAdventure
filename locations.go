package main

import (
	"fmt"
	"os"
)

func castle() {
	fmt.Println("===========================================")
	fmt.Println()
	fmt.Println("You have been tasked by the king to kill a dragon that has been terrorizing the land.")
	fmt.Println("Upon exiting the castle, you consider your options.")
	fmt.Println("The situation with the dragon is dire. Every day, someone suffers from the dragons assaults.")
	fmt.Println("In order to waste no time, you consider going home to make preparations for an immediate departure.")
	fmt.Println("You also consider that there may be someone in town that could help you in some way.")

	switch getInput("Do you want to go HOME or into TOWN? ") {
	case "home":
		fmt.Println("Your heart aches for those that are suffering.")
		fmt.Println("You go home to get ready.")
		fmt.Println("You pack your supplies and eat a hearty meal.")
		fmt.Println("The road will be long, it may be some time before you get a decent meal and night's sleep again.")
		fmt.Println("Before the sun peeks over the horizon, you are on your way.")
		fmt.Println()
		roadOne()
	case "town":
		fmt.Println("You head to town to talk to the townsfolk to get info that might help kill the dragon.")
		town()
	default:
		wrongInput()
		castle()
	}
}
func crossroads() {
	fmt.Println("You arrive at the crossroads.")
	fmt.Println("You can go north, east, south, west or back home.")
	switch getInput("which way will you go") {
	case "north":
		fmt.Println("change of coords to the north")
		fmt.Println("Meh, you actually want to go east for now")
		crossroads()
	case "east":
		fmt.Println("change of coords to the east")
		fmt.Println("You decide to go east for now")
		roadEvent()
		roadEvent()
		switch getInput("The road splits. Continue straight or turn north") {
		case "straight":
			roadEvent()
			roadEvent()
			westernmountains()
		case "north":
			roadEvent()
			roadEvent()
			player1.KnownLocations.Discover("village")
			village()
		}
	case "south":
		fmt.Println("change of coords to the south")
		fmt.Println("Meh, you actually want to go east for now")
		crossroads()
	case "west":
		fmt.Println("change of coords to the west")
		fmt.Println("Meh, you actually want to go east for now")
		crossroads()
	case "home":
		fmt.Println("go home for new options")
		fmt.Println("Meh, you actually want to go east for now")
		crossroads()
	default:
		wrongInput()
		crossroads()

	}

}

func roadOne() {
	fmt.Println("===========================================")
	fmt.Println()
	var giveAid string
	fmt.Println("And so your journey begins. The road ahead seems long.")
	fmt.Println("You come across a vagrant lying on the side of the road.")
	giveAid = getInput("Do you get off your horse to CHECK ON HIM or RIDE ON and leave him be?")
	if giveAid == "check on him" {
		fmt.Println("You get off your horse and as you bend down to roll the vagrant over to see how he is")
		fmt.Println("As you crouch down with concern, you are ambushed from behind.")
		fmt.Println("The bandits kill you and take your horse and belongings.")
		fmt.Println("The End")
		//showP1()
		os.Exit(0)
	} else if giveAid == "ride on" {
		fmt.Println("You ignore the vagrant and continue on your quest.")
		crossroads()
	} else {
		wrongInput()
		crossroads()
	}
}

func roadTwo() {
	fmt.Println("===========================================")
	fmt.Println()
	fmt.Println("Time to go back the way you came.")
	fmt.Println("But you realize that it is getting late.")
	delay := getInput("Do you find an INN for the night or HEAD BACK since there's no time to lose?")
	if delay == "inn" {
		fmt.Println()
		fmt.Println("Deciding it best not to exhaust yourself, you seek out an inn.")
		fmt.Println("The innkeeper has nothing but a pallet in the main hall, but it's better than stony ground under a bush.")
		fmt.Println("You awaken somewhat refreshed, and head back to town")
		player1.Traits.Tired = false
		town()
	} else if delay == "head back" {
		fmt.Println()
		fmt.Println("The sun is getting low as you head out, but the dragon must be stopped.")
		fmt.Println("No one else will die if you can help it.")
		town()
		fmt.Println()

	} else {
		wrongInput()
		roadTwo()

	}
	fmt.Println()

}

func village() {
	fmt.Println("===========================================")
	fmt.Println()
	fmt.Println("It's been a long ride, but you finally come upon a small village.")

	switch getInput("Will you REST for the night or KEEP GOING") {
	case "rest":
		fmt.Println("the innkeeper offers you a room and meal and you awake refreshed.")
		player1.Traits.Tired = false
		fmt.Println("from your upper room, you can see a cave in the distance.")
		player1.KnownLocations.Discover("cave")
		chooseDestination("village")
	case "keep going":
		fmt.Println("You ride on, hoping to make it just a bit further.")
		fmt.Println("However, you and your horse tire before reaching the next village.")
		fmt.Println("You lay out your bedroll under the stars.")
		fmt.Println("It rains. You sleep very little.")
		player1.Traits.Tired = true
		fmt.Println("Now that there's some light to see by, you notice a cave not too far from your sleeping place.")
		player1.KnownLocations.Discover("cave")
		chooseDestination("village")
	default:
		wrongInput()
		village()
	}
}

func cave() {
	fmt.Println("===========================================")
	fmt.Println("It's time to make the trip to the dragon's lair.")
	player1.KnownLocations.Discover("cave entrance")
	caveEntrance()
}

func caveEntrance() {
	fmt.Println("===========================================")
	fmt.Println()
	fmt.Println("You finally arrive at what must be the dragon's lair.")

	switch getInput("Do you ENTER immediately or LOOK AROUND? ") {
	case "enter":
		fmt.Println()
		fmt.Println("You decide that it's been a long day, and you just want to get this over with.")
		fmt.Println("Immediately, you come face to face with the dragon, as though it's been waiting for you.")
		fmt.Println("It unleashes a blast of flaming breath.")
		if !player1.Traits.Tired {
			fmt.Println()
			fmt.Println("you quickly dodge into a small side tunnel.")
			fmt.Println("While the dragon tries to find you, you manage to sneak back out of the tunnel.")
			fmt.Println("Maybe looking for another way in would be a good idea after all.")
			fmt.Println("next function called: caveEntrance")

			caveEntrance()
		} else if player1.Traits.Tired {
			fmt.Println()
			fmt.Println("Exhausted from your bad sleep, you move too slowly.")
			fmt.Println("The dragon's flaming breath engulfs you.")
			fmt.Println("You die horribly, cooking in your own armor.")
			//showP1()
			os.Exit(0)
		}
	case "look around":
		fmt.Println("next function called: caveEntrance2")

		caveEntrance2()
	default:
		wrongInput()
		caveEntrance()
	}
}

func caveEntrance2() {
	fmt.Println("You find a small entry tunnel that may lead into the dragon's lair.")
	if player1.Traits.Tired {
		fmt.Println()
		fmt.Println("Exhausted from a lack of sleep, you wander for hours and get lost in the tunnels.")
		fmt.Println("Unable to find your way out again before collapsing from exhaustion, you die a cold death on the dirty stone ground.")
		//showP1()
		os.Exit(0)
	} else if !player1.Traits.Tired {
		fmt.Println()
		fmt.Println("Your alertness helps you find your way through the maze of tunnels.")
		fmt.Println("Eventually, you come upon the dragon from behind.")
		fmt.Println("next function called: caveInterior")
		caveInterior()
	}
}

func caveInterior() {
	fmt.Println("===========================================")
	fmt.Println()
	fmt.Println("next function called: showWeapons")

	showWeapons()
	chooseWeapon := getInput("What weapon will you use")
	toCheck = chooseWeapon
	result := isAvailable(player1.Inventory.Weapons, toCheck)
	if !result {
		fmt.Println("You don't have that weapon")
		fmt.Println("next function called: caveInterior")

		caveInterior()

	} else {
		if chooseWeapon == "blacksmith spear" {
			fmt.Println()
			fmt.Println("You level your weapon at the dragon and creep into striking distance.")
			fmt.Println("The weapon slides through the dragonhide smoothly and pierces the heart of the monster.")
			fmt.Println("Relief washes over you as you realize you've succeeded.")
		} else if chooseWeapon == "dragon killer" {
			fmt.Println()
			fmt.Println("You level your weapon at the dragon and creep into striking distance.")
			fmt.Println("As you attempt to stab the beast, the weapon snaps in two.")
			fmt.Println("The monstrous beast turns upon you.")
			fmt.Println("Snapped up into it's jaws, at least your death is quick.")
		} else if chooseWeapon == "sword" {
			fmt.Println()
			fmt.Println("Trusting to your sword, you attack the beast with a battle cry.")
			fmt.Println("Unfortunately, your scream and your sword are inadequate to pierce the monster's hide.")
			fmt.Println("You die as a gout of flame erupts from the dragon's maw.")
		} else {
			wrongInput()
		}
	}
}
func home() {
	//player1.CurrentLocation.Location = "home"

	fmt.Println("Your heart aches for those that are suffering.")
	fmt.Println("You go home to get ready.")
	fmt.Println("You pack your supplies and eat a hearty meal.")
	fmt.Println("The road will be long, it may be some time before you get a decent meal and night's sleep again.")
	fmt.Println("Before the sun peeks over the horizon, you are on your way.")
	fmt.Println()
	fmt.Println("more options to come")
	fmt.Println("next function called: chooseDestination")
	chooseDestination("home")
}

func hamlet() {
	fmt.Println("You rush out of town through the south gate in search of the hamlet.")
	fmt.Println("After many stops in every village you finally meet the man who survived.")
	fmt.Println("He tells you that he was one of 20 soldiers that attacked the dragon. None of them even scratched the beast.")
	brave := getInput("Do you still want to face the dragon ")
	if brave == "yes" {
		fmt.Println("Knowing it will likely mean death, you decide to continue your quest.")
		player1.Traits.Brave = true
		fmt.Println("next function called: roadTwo")
		roadTwo()
	} else {
		fmt.Println()
		fmt.Println("You are a coward.")
		fmt.Println("In hopeless despair, you ride as far away as possible, never to be heard from again.")
		fmt.Println("The End")
		//showP1()
		os.Exit(0)
	}
}
func easternwasteland() {
	fmt.Println("You've discovered the Eastern Wastes.")
	fmt.Println("options and adventure to come.")
	fmt.Println("For now, you must turn back.")

}

func westernmountains() {
	fmt.Println("You've discovered the Western Mountains.")
	fmt.Println("options and adventure to come.")
	fmt.Println("For now, you must turn back.")
	fmt.Println("You return to the last fork and head north")
	village()
}

// func southernseas() {
// 	fmt.Println("You've discovered the Southern Seas.")
// 	fmt.Println("options and adventure to come.")
// 	fmt.Println("For now, you must turn back.")
// }
// func northerntundra() {
// 	fmt.Println("You've discovered the Northern Tundra.")
// 	fmt.Println("options and adventure to come.")
// 	fmt.Println("For now, you must turn back.")
// }

// func lostwoods() {
// 	fmt.Println("options and adventure to come")
// }
