package main

import (
	"fmt"
	"os"
)

// using variables in strings
// two ways of concatenating strings
// options array function

func town() {
	fmt.Println()
	if brave != "ignored" {
		fmt.Println("Since you're here again anyway, you consider taking some more time for another villager.")
		fmt.Println("But then again, the longer you delay, the more others will suffer.")

		query := getInput("Would you like to talk to another villager?")
		if query == "no" {
			fmt.Println("Realizing you've spent a fair bit of time chasing tales already, you feel you must be getting on with it.")
			roadOne()
		} else if query == "yes" {
			switch getInput("Which one will it be this time? The TOWN CRIER, the TRAVELING MERCHANT or the BLACKSMITH?") {
			case "town crier":
				crier()
				roadOne()
			case "traveling merchant":
				merchant()
				roadOne()
			case "blacksmith":
				fmt.Println()
				blacksmith()
				roadOne()
			default:
				wrongInput()
				town()
			}
		} else {
			wrongInput()
			town()
		}
	} else {
		fmt.Println()
		townsfolk()
		fmt.Println("Unfortunately, you only have time for one conversation.")
		switch getInput("Would you like to talk to the TOWN CRIER, the INNKEEPER, the TRAVELING MERCHANT or the BLACKSMITH?") {
		case "town crier":
			crier()
			roadOne()
		case "innkeeper":
			innkeep()
		case "traveling merchant":
			merchant()
			roadOne()
		case "blacksmith":
			blacksmith()
			roadOne()
		default:
			wrongInput()
			town()
		}
	}
}

func blacksmith() {
	fmt.Println()
	fmt.Println("The dragon killed his daughter. He has created a spear designed to pierce dragon hide.")
	blksmtOffer = getInput("He offers it to you for your quest. Do you offer him money? ")
	if blksmtOffer == "yes" {
		fmt.Println("The blacksmith refuses your money, and gives you the spear with a blessing.")
		player1.Inventory.Weapon = "trueDragonKiller"
	} else {
		fmt.Println("The blacksmith says he doesn't want any money.")
		blksmtOffer = getInput("He begs you to take the spear to avenge his daughter. Do you accept?")
		if blksmtOffer == "yes" {
			fmt.Println("You accept the spear and promise him the death of the dragon.")
			player1.Inventory.Weapon = "trueDragonKiller"
		} else if blksmtOffer == "no" {
			fmt.Println("You insist that you can not carry such a weapon.")
			fmt.Println("You leave town with just your sword.")
		}
	}
}
func merchant() {
	fmt.Println()
	dk = getInput("He offers you a weapon he deems a ‘dragon killer’. Will you buy it?")
	if dk == "no" {
		fmt.Println("You don't trust him, and decline the weapon. You leave town with just your sword.")
	} else if dk == "yes" {
		fmt.Println("Hoping you can trust him, you pay a small fortune for the spear.")
		player1.Inventory.Weapon = "fakeDragonKiller"
	}

}
func innkeep() {
	fmt.Println()
	fmt.Println("You decide on the innkeeper.")
	fmt.Println("Surely, someone has told him something of the dragon.")
	fmt.Println("You sidle up to the bar and order a drink.")
	fmt.Println("As he sets the glass in front of you, you ask,")
	fmt.Println("'Tell me, good sir, know you anything about that fearsome dragon?'")
	fmt.Println("He tells of a man in a small village to the south who survived a dragon attack.")
	fmt.Println("You rush out of town through the south gate in search of the village.")
	fmt.Println("After many stops in every village you finally meet the man who survived.")
	fmt.Println("He tells you that he was one of 20 soldiers that attacked the dragon. None of them even scratched the beast.")
	brave = getInput("Do you still want to face the dragon? ")
	if brave == "yes" {
		fmt.Println("Knowing it will likely mean death, you decide to continue your quest.")
		roadTwo()
	} else {
		fmt.Println()
		fmt.Println("You are a coward.")
		fmt.Println("In hopeless despair, you ride as far away as possible, never to be heard from again.")
		fmt.Println("The End")
		showP1()
		os.Exit(0)
	}
}
func crier() {
	fmt.Println()
	fmt.Println("He talks for an hour, but provides no useful information.")
	fmt.Println("You feel your time has been wasted, but you must carry on.")
}
func townsfolk() {
	fmt.Println("As you wander around the town, you see many people that might be able to help you.")
	fmt.Println("There's the town crier over there.")
	fmt.Println("Because he has to make announcements for the benefit of the town,")
	fmt.Println("he would likely be well informed on a diverse number of topics.")
	fmt.Println("You consider the blacksmith.")
	fmt.Println("Recently he has been less fun to be around, but he might have a useful tool lying about.")
	fmt.Println("The innkeeper has probably heard a lot of stories from various travelers.")
	fmt.Println("But how many of them are actually true?")
	fmt.Println("The most compelling stories are usually the most embellished.")
	fmt.Println("And then there's the merchant. He certainly gets around.")
	fmt.Println("What has he seen and heard in his travels?")
	fmt.Println("More importantly, what has he got tucked away in that cart of his?")
}
