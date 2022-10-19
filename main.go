package main

// refactor into scenes
// globals are inventory possibilities or knowledge - booleans?
// globals open new options
// one scene calls next scene depending on choice
// locations
// tracking state of users
// struct for inventory

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var (
	blksmtOffer string = "ignored"
	dk          string = "ignored"
	brave       string = "ignored"
	giveAid     string = "ignored"
	rest        string = "ignored"
	rested      string = "ignored"

	//enter       string = "ignored"
)

type Player struct {
	Name string
}

func main() {
	fmt.Println("Choose your own adventure - The Knight and the Dragon")
	questBegin()

	fmt.Println()

}

func checkVariables() {
	fmt.Println()
	fmt.Println()
	fmt.Println("VARIABLE VALUE CHANGES")

	fmt.Println("blksmtOffer-", blksmtOffer)
	fmt.Println("dk----------", dk)
	fmt.Println("brave-------", brave)
	fmt.Println("giveAid-----", giveAid)
	fmt.Println("rest--------", rest)
	fmt.Println("rested------", rested)

	fmt.Println()

}
func getInput(question string) string {
	fmt.Println(question)
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

func wrongInput() {
	fmt.Println()
	fmt.Println()
	fmt.Println("Alas! Thy speech is unintelligible!")
	fmt.Println()

}
func questBegin() {
	fmt.Println()
	fmt.Println()
	fmt.Println("Welcome, brave adventurer!")
	fmt.Println("On this journey, you will face many choices.")
	fmt.Println("Sometimes choices will appear to you 'LIKE THIS', and you are to respond 'like this'.")
	fmt.Println("Other times it will be a simple yes or no question.")
	fmt.Println("And now, be off! Your quest awaits!")
	fmt.Println()
	castle()

}

func castle() {
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("You have been tasked by the king to kill a dragon that has been terrorizing the land.")
	switch getInput("Upon exiting the castle, do you want to go HOME or into TOWN? ") {
	case "home":
		fmt.Println("You go home to eat a hearty meal before setting out. Then you saddle your horse and ride for the dragon's lair.")
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

func town() {
	fmt.Println()
	fmt.Println()
	fmt.Println()
	if brave != "ignored" {
		query := getInput("Since you're here again anyway, would you like to talk to another villager?")
		if query == "no" {
			fmt.Println("Realizing you've spent a fair bit of time chasing tales already, you feel you must be getting on with it.")
			roadOne()
		} else if query == "yes" {
			switch getInput("Would you like to talk to the TOWN CRIER, the TRAVELING MERCHANT or the BLACKSMITH?") {

			case "town crier":
				fmt.Println()
				fmt.Println("He talks for an hour, but provides no useful information.")
				fmt.Println("You feel your time has been wasted, but you must carry on.")
				roadOne()

			case "traveling merchant":
				fmt.Println()
				dk = getInput("He offers you a weapon he deems a ‘dragon killer’. Will you buy it?")
				if dk == "no" {
					fmt.Println("You don't trust him, and decline the weapon. You leave town with just your sword.")
				} else if dk == "yes" {
					fmt.Println("Hoping you can trust him, you pay a small fortune for the spear.")
				}

				roadOne()

			case "blacksmith":
				fmt.Println()
				fmt.Println("The dragon killed his daughter. He has created a spear designed to pierce dragon hide.")
				blksmtOffer = getInput("He offers it to you for your quest. Do you offer him money? ")
				if blksmtOffer == "yes" {
					fmt.Println("The blacksmith refuses your money, and gives you the spear with a blessing.")
				} else {
					fmt.Println("The blacksmith says he doesn't want any money.")
					blksmtOffer = getInput("He begs you to take the spear to avenge his daughter. Do you accept?")
					if blksmtOffer == "yes" {
						fmt.Println("You accept the spear and promise him the death of the dragon.")
					} else if blksmtOffer == "no" {
						fmt.Println("You insist that you can not carry such a weapon.")
						fmt.Println("You leave town with just your sword.")
					}
				}
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

		fmt.Println("As you wander around the town, you see many people that might be able to help you.")
		fmt.Println("Unfortunately, you only have time for one conversation.")
		switch getInput("Would you like to talk to the TOWN CRIER, the INNKEEPER, the TRAVELING MERCHANT or the BLACKSMITH?") {
		case "town crier":
			fmt.Println()
			fmt.Println("He talks for an hour, but provides no useful information.")
			fmt.Println("You feel your time has been wasted, but you must carry on.")
			roadOne()
		case "innkeeper":
			fmt.Println()
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
				checkVariables()
				os.Exit(0)
			}
		case "traveling merchant":
			fmt.Println()
			dk = getInput("He offers you a weapon he deems a ‘dragon killer’. Will you buy it?")
			if dk == "no" {
				fmt.Println("You don't trust him, and decline the weapon. You leave town with just your sword.")
			} else if dk == "yes" {
				fmt.Println("Hoping you can trust him, you pay a small fortune for the spear.")
			}

			roadOne()

		case "blacksmith":
			fmt.Println()
			fmt.Println("The dragon killed his daughter. He has created a spear designed to pierce dragon hide.")
			blksmtOffer = getInput("He offers it to you for your quest. Do you offer him money? ")
			if blksmtOffer == "yes" {
				fmt.Println("The blacksmith refuses your money, and gives you the spear with a blessing.")
			} else {
				fmt.Println("The blacksmith says he doesn't want any money.")
				blksmtOffer = getInput("He begs you to take the spear to avenge his daughter. Do you accept?")
				if blksmtOffer == "yes" {
					fmt.Println("You accept the spear and promise him the death of the dragon.")
				} else if blksmtOffer == "no" {
					fmt.Println("You insist that you can not carry such a weapon.")
					fmt.Println("You leave town with just your sword.")
				}
			}
			roadOne()
		default:
			wrongInput()
			town()
		}

	}
}

func roadTwo() {
	fmt.Println()
	fmt.Println("Time to go back the way you came.")
	fmt.Println("But you realize that it is getting late.")
	delay := getInput("Do you find an INN for the night or HEAD BACK since there's no time to lose?")
	if delay == "inn" {
		fmt.Println()
		fmt.Println("Deciding it best not to exhaust yourself, you seek out an inn.")
		fmt.Println("The innkeeper has nothing but a pallet in the main hall, but it's better than stony ground under a bush.")
		fmt.Println("You awaken somewhat refreshed, and head back to town")
		rested = "yes"
		town()
	} else if delay == "head back" {
		fmt.Println()
		fmt.Println("The sun is getting low as you head out, but the dragon must be stopped.")
		fmt.Println("No one else will die if you can help it.")
		rested = "no"
		town()
		fmt.Println()

	} else {
		wrongInput()
		roadTwo()

	}
	fmt.Println()
	fmt.Println()

}
func roadOne() {
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("And so your journey begins. The road ahead seems long.")
	fmt.Println("You come across a vagrant lying on the side of the road.")
	giveAid = getInput("Do you get off your horse to CHECK ON HIM or RIDE ON and leave him be?")
	if giveAid == "check on him" {
		fmt.Println("You get off your horse and as you bend down to roll the vagrant over to see how he is")
		fmt.Println("As you crouch down with concern, you are ambushed from behind.")
		fmt.Println("The bandits kill you and take your horse and belongings.")
		fmt.Println("The End")
		checkVariables()
		os.Exit(0)
	} else if giveAid == "ride on" {
		fmt.Println("You ignore the vagrant and continue on your quest.")
		village()
	} else {
		wrongInput()
		roadOne()
	}
}

func village() {
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("After riding for what seems like weeks, you come upon a small village.")
	switch getInput("Will you REST for the night or KEEP GOING?") {
	case "rest":
		fmt.Println("the innkeeper offers you a room and meal and you awake refreshed.")
		rest = "yes"
		caveEntrance()
	case "keep going":
		fmt.Println("You ride on, hoping to make it just a bit further.")
		fmt.Println("However, you and your horse tire before reaching the next village.")
		fmt.Println("You lay out your bedroll under the stars.")
		fmt.Println("It rains. You sleep very little.")
		rest = "no"
		caveEntrance()
	default:
		wrongInput()
		village()
	}
}

func caveEntrance() {
	fmt.Println()
	fmt.Println()
	fmt.Println()
	fmt.Println("You finally arrive at what must be the dragon's lair.")

	switch getInput("Do you ENTER immediately or LOOK AROUND? ") {
	case "enter":
		fmt.Println()
		fmt.Println("You decide that it's been a long day, and you just want to get this over with.")
		fmt.Println("Immediately, you come face to face with the dragon, as though it's been waiting for you.")
		fmt.Println("It unleashes a blast of flaming breath.")
		if rest == "yes" && rested == "yes" {
			fmt.Println()
			fmt.Println("you quickly dodge into a small side tunnel.")
			fmt.Println("While the dragon tries to find you, you manage to sneak back out of the tunnel.")
			fmt.Println("Maybe looking for another way in would be a good idea after all.")
			caveEntrance()
		} else if rest != "yes" || rested != "yes" {
			fmt.Println()
			fmt.Println("Exhausted from your bad sleep, you move too slowly.")
			fmt.Println("The dragon's flaming breath engulfs you.")
			fmt.Println("You die horribly, cooking in your own armor.")
			checkVariables()
			os.Exit(0)
		}
	case "look around":
		fmt.Println("You find a small entry tunnel that may lead into the dragon's lair.")
		if rest != "yes" || rested != "yes" {
			fmt.Println()
			fmt.Println("Exhausted from a lack of sleep, you wander for hours and get lost in the tunnels.")
			fmt.Println("Unable to find your way out again before collapsing from exhaustion, you die a cold death on the dirty stone ground.")
			checkVariables()
			os.Exit(0)
		} else if rest == "yes" && rested == "yes" {
			fmt.Println()
			fmt.Println("Your alertness helps you find your way through the maze of tunnels.")
			fmt.Println("Eventually, you come upon the dragon from behind.")
			caveInterior()
		}
	default:
		wrongInput()
		caveEntrance()
	}
}
func caveInterior() {
	fmt.Println()
	fmt.Println()
	fmt.Println()
	if blksmtOffer == "yes" {
		fmt.Println()
		fmt.Println("You level your weapon at the dragon and creep into striking distance.")
		fmt.Println("The weapon slides through the dragonhide smoothly and pierces the heart of the monster.")
		fmt.Println("Relief washes over you as you realize you've succeeded.")
	} else if dk == "yes" {
		fmt.Println()
		fmt.Println("You level your weapon at the dragon and creep into striking distance.")
		fmt.Println("As you attempt to stab the beast, the weapon snaps in two.")
		fmt.Println("The monstrous beast turns upon you.")
		fmt.Println("Snapped up into it's jaws, at least your death is quick.")
	} else {
		fmt.Println()
		fmt.Println("Trusting to your sword, you attack the beast with a battle cry.")
		fmt.Println("Unfortunately, your sword is inadequate to pierce the monster's hide.")
		fmt.Println("You die as a gout of flame erupts from the dragon's maw.")
	}
}
