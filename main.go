package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

var firstDest string = "ignored"
var villager string = "ignored"
var blksmtOffer string = "ignored"
var dk string = "ignored"
var brave string = "ignored"
var giveAid string = "ignored"
var rest string = "ignored"
var enter string = "ignored"

func main() {
	fmt.Println("Choose your own adventure - The Knight and the Dragon")
	actOne()
	switch firstDest {
	case "home":
		actThree()
		fmt.Println()
	case "town":
		actTwo()
		fmt.Println()
		actThree()
		fmt.Println()
	}
	fmt.Println()
	actFour()
	fmt.Println()
	fmt.Println("Another half day's ride brings you to the dragon's lair.")
	fmt.Println()
	actSix()
	fmt.Println()
	actSeven()
	fmt.Println()

	// fmt.Println()
	// fmt.Println("VARIABLE VALUE CHANGES")
	// fmt.Println("firstDest---", firstDest)
	// fmt.Println("villager----", villager)
	// fmt.Println("blksmtOffer-", blksmtOffer)
	// fmt.Println("dk----------", dk)
	// fmt.Println("brave-------", brave)
	// fmt.Println("giveAid-----", giveAid)
	// fmt.Println("rest-----", rest)

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
		fmt.Println("I don't understand your choice. Please try again.")
		return getInput(question)
	}
	return input1

}

func actOne() {
	fmt.Println()
	fmt.Println("You have been tasked by the king to kill a dragon that has been terrorizing the land.")
	firstDest = getInput("Upon exiting the castle, do you want to go HOME or into TOWN? ")
	if firstDest == "home" {
		fmt.Println("You go home to eat a hearty meal before setting out. Then you saddle your horse and ride for the dragon's lair.")
	} else if firstDest == "town" {
		fmt.Println("You head to town to talk to the townsfolk to get info that might help kill the dragon.")
	} else {
		fmt.Println()
		fmt.Println("I don't understand your answer. Let's try again.")
		fmt.Println()
		actOne()
	}

}

func actTwo() {
	fmt.Println()
	fmt.Println("As you wander around the town, you see many people that might be able to help you.")
	fmt.Println("Unfortunately, you only have time for one conversation.")
	villager = getInput("Would you like to talk to the TOWN CRIER, the INNKEEPER, the TRAVELING MERCHANT or the BLACKSMITH?")
	if villager == "town crier" {
		fmt.Println("He talks for an hour, but provides no useful information.")
		fmt.Println("You feel your time has been wasted, but you must carry on.")
	} else if villager == "innkeeper" {
		fmt.Println("He tells of a man in a small village to the south who survived a dragon attack.")
		fmt.Println("You rush out of town through the south gate in search of the village.")
		fmt.Println("After many stops in every village you finally meet the man who survived.")
		fmt.Println("He tells you that he was one of 20 soldiers that attacked the dragon. None of them even scratched the beast.")
		brave = getInput("Do you still want to face the dragon? ")
		if brave == "yes" {
			fmt.Println("Knowing you might be riding to your death, you head back the way you came to face the beast.")
		} else {
			fmt.Println("In hopeless despair, you ride as far away as possible, never to be heard from again.")
			fmt.Println("The End")
			os.Exit(0)
		}
	} else if villager == "merchant" || villager == "traveling merchant" {
		dk = getInput("He offers you a weapon he deems a ‘dragon killer’. Will you buy it?")
		if dk == "no" {
			fmt.Println("You don't trust him, and decline the weapon. You leave town with just your sword.")
		} else if dk == "yes" {
			fmt.Println("Hoping you can trust him, you pay a small fortune for the spear.")
		}
	} else if villager == "blacksmith" {
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
	} else {
		fmt.Println()
		fmt.Println("I don't understand your answer. Let's try again.")
		fmt.Println()
		actTwo()
	}

}

func actThree() {
	fmt.Println()
	fmt.Println("And so your journey begins. The road ahead seems long.")
	fmt.Println("You come across a vagrant lying on the side of the road.")
	giveAid = getInput("Do you get off your horse to CHECK ON HIM or RIDE ON and leave him be?")
	if giveAid == "check on him" {
		fmt.Println("You get off your horse and as you bend down to roll the vagrant over to see how he is")
		fmt.Println("As you crouch down with concern, you are ambushed from behind.")
		fmt.Println("The bandits kill you and take your horse and belongings.")
		fmt.Println("The End")
		os.Exit(0)
	} else if giveAid == "ride on" {
		fmt.Println("You ignore the vagrant and continue on your quest.")
	} else {
		fmt.Println()
		fmt.Println("I don't understand your answer. Let's try again.")
		fmt.Println()
		actThree()
	}
}

func actFour() {
	fmt.Println()
	fmt.Println("After riding for what seems like days, you come upon a small village.")
	rest = getInput("Will you REST for the night or KEEP GOING?")
	if rest == "rest" {
		fmt.Println("the innkeeper offers you a room and meal and you awake refreshed.")
	} else if rest == "keep going" {
		fmt.Println("You ride on, hoping to make it just a bit further.")
		fmt.Println("However, you and your horse tire before reaching the next village.")
		fmt.Println("You lay out your bedroll under the stars.")
		fmt.Println("It rains. You sleep very little.")
	} else {
		fmt.Println()
		fmt.Println("I don't understand your answer. Let's try again.")
		fmt.Println()
		actFour()
	}
}

func actSix() {
	fmt.Println()

	enter = getInput("Do you ENTER immediately or LOOK AROUND? ")
	if enter == "enter" {
		fmt.Println("You decide that it's been a long day, and you just want to get this over with.")
		fmt.Println("Immediately, you come face to face with the dragon, as though it's been waiting for you.")
		fmt.Println("It unleashes a blast of flaming breath.")
		if rest == "rest" {
			fmt.Println("you quickly dodge into a small side tunnel.")
			fmt.Println("While the dragon tries to find you, you manage to sneak back out of the tunnel.")
			fmt.Println("Maybe looking for another way in would be a good idea after all.")
			actSix()
		} else if rest == "keep going" {
			fmt.Println("Exhausted from your bad sleep, you move too slowly.")
			fmt.Println("The dragon's flaming breath engulfs you.")
			fmt.Println("You die horribly, cooking in your own armor.")
			os.Exit(0)
		}
	} else if enter == "look around" {
		fmt.Println("You find a small entry tunnel that may lead into the dragon's lair.")
		if rest == "keep going" {
			fmt.Println("Exhausted from a lack of sleep, you wander for hours and get lost in the tunnels.")
			fmt.Println("Unable to find your way out again before collapsing from exhaustion, you die a cold death on the dirty stone ground.")
			os.Exit(0)
		} else if rest == "rest" {
			fmt.Println("Your alertness helps you find your way through the maze of tunnels.")
			fmt.Println("Eventually, you come upon the dragon from behind.")
		}
	} else {
		fmt.Println()
		fmt.Println("I don't understand your answer. Let's try again.")
		fmt.Println()
		actSix()
	}
}
func actSeven() {
	fmt.Println()
	if blksmtOffer == "yes" {
		fmt.Println("You level your weapon at the dragon and creep into striking distance.")
		fmt.Println("The weapon slides through the dragonhide smoothly and pierces the heart of the monster.")
		fmt.Println("Relief washes over you as you realize you've succeeded.")
	} else if dk == "yes" {
		fmt.Println("You level your weapon at the dragon and creep into striking distance.")
		fmt.Println("As you attempt to stab the beast, the weapon snaps in two.")
		fmt.Println("The monstrous beast turns upon you.")
		fmt.Println("Snapped up into it's jaws, at least your death is quick.")
	} else {
		fmt.Println("Trusting to your sword, you attack the beast with a battle cry.")
		fmt.Println("Unfortunately, your sword is inadequate to pierce the monster's hide.")
		fmt.Println("You die as a gout of flame erupts from the dragon's maw.")
	}
}
