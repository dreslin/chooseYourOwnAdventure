package main

import (
	"fmt"
	"strings"
)

// using variables in strings
// two ways of concatenating strings
// options array function
var options []string
var option string

func town2() {
	player1.Traits.Brave = false
	fmt.Println()
	fmt.Println("===========================================")

	options = []string{" Town Crier", " Traveling Merchant", " Blacksmith"}

	//fmt.Print(strings.ToUpper(strings.Trim(fmt.Sprint(options), "[]")))
	fmt.Println()
	if !player1.Traits.Brave {
		options = append(options, " or the innkeeper")
		for _, options := range options {
			fmt.Print(strings.ToUpper(fmt.Sprint(options)))
		}

		fmt.Println()

		townfolk()
		fmt.Println("Unfortunately, you only have time for one conversation.")
		option = getInput("Would you like to talk to the", options...)
		fmt.Printf("You decide it would be worth your while to talk to %v", option)
	} else if player1.Traits.Brave {

		fmt.Println()
		fmt.Println("Since you're here again anyway, you consider taking some more time for another villager.")
		fmt.Println("But then again, the longer you delay, the more others will suffer.")
		query := getInput("Would you like to talk to another villager?")
		if query == "no" {
			fmt.Println("Realizing you've spent a fair bit of time chasing tales already, you feel you must be getting on with it.")
			roadOne()
		} else if query == "yes" {
			option = getInput("Would you like to talk to the TOWN CRIER, the TRAVELING MERCHANT or the BLACKSMITH?")
			fmt.Printf("You decide it would be worth your while to talk to %v", option)

		}

		switch option {
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
			town2()
		}
	}
}
