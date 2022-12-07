package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func getInput(question string, options ...string) string {
	fmt.Print(question)
	for _, options := range options {
		fmt.Print(strings.ToUpper(fmt.Sprint(options)))
	}
	fmt.Println("?")

	inputReader := bufio.NewReader(os.Stdin)
	input1, err := inputReader.ReadString('\n')
	if err != nil {
		panic(err)
	}
	input1 = strings.TrimSuffix(input1, "\n")
	if err != nil {

		return getInput(question)
	}
	return input1

}

// func getInput2(question string, options ...string) string {
// 	fmt.Printf("%v?\n", question)
// 	for _, option := range options {
// 		fmt.Print(strings.ToUpper(fmt.Sprint(option)))
// 		fmt.Print(" ")
// 	}

// 	inputReader := bufio.NewReader(os.Stdin)
// 	input1, err := inputReader.ReadString('\n')
// 	if err != nil {
// 		panic(err)
// 	}
// 	input1 = strings.TrimSuffix(input1, "\n")
// 	if err != nil {

// 		return getInput(question)
// 	}
// 	return input1

// }

func wrongInput() {
	fmt.Println()
	fmt.Println("Alas! Thy speech is unintelligible!")
	fmt.Println()

}

func showP1() {
	fmt.Println(player1)
}
func explore() {
	fmt.Println("this choice will perform a random search of the area.")
	fmt.Println("try something else.")
	chooseDestination("explore")
}

func isAvailable(a []string, str string) bool {

	// a is the array I want to check through
	for i := 0; i < len(a); i++ {
		if a[i] == str {
			return true
		}
	}
	return false
}

// 	whereto := getInput("Where do you want to go or would you like to explore the area")
// result := isAvailable(player1.KnownLocations.Locations, whereto)
// 	toCheck = whereto
// 	result := isAvailable(player1.KnownLocations.Locations, toCheck)
// 	if result {
// 		locale[whereto]()
// 	} else {
// 		fmt.Println("you don't know that place yet")

func showAtlas(CurrentLocation string) {
	loci := deleteElement(player1.KnownLocations.Locations, CurrentLocation)
	fmt.Println(strings.ToUpper(strings.Trim(fmt.Sprint(loci), "[]")))
}
func showWeapons() {
	arms := player1.Inventory.Weapons
	fmt.Println(strings.ToUpper(strings.Trim(fmt.Sprint(arms), "[]")))
}

func deleteElement(l []string, local string) []string {
	for k, v := range l {
		if v == local {
			return append(l[:k], l[k+1:]...)

		}
	}

	panic("Could not find element to delete")
}
