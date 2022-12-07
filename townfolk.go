package main

import (
	"fmt"
	"os"
)

func blacksmith() {
	fmt.Println()
	fmt.Println("-------------------")
	fmt.Println("You have always known the blacksmith to be a reliable man.")
	fmt.Println("While you aren't sure what happened to change his demeanor,")
	fmt.Println("you believe he's your best chance for reliable aid.")
	fmt.Println("The dragon killed his daughter. He has created a spear designed to pierce dragon hide.")
	blksmtOffer := getInput("He offers it to you for your quest. Do you offer him money? ")
	if blksmtOffer == "yes" {
		fmt.Println("The blacksmith refuses your money, and gives you the spear with a blessing.")
		player1.Inventory.Acquire("blacksmith spear")
	} else {
		fmt.Println("The blacksmith says he doesn't want any money.")
		blksmtOffer = getInput("He begs you to take the spear to avenge his daughter. Do you accept?")
		if blksmtOffer == "yes" {
			fmt.Println("You accept the spear and promise him the death of the dragon.")
			player1.Inventory.Acquire("blacksmith spear")
		} else if blksmtOffer == "no" {
			fmt.Println("You insist that you can not carry such a weapon.")
			fmt.Println("You leave town with just your sword.")
		}
	}
}
func merchant() {
	fmt.Println()
	fmt.Println("-------------------")
	fmt.Println("That merchant must have picked up something in his travels.")
	fmt.Println("You decide the go see what he's got for sale.")
	fmt.Println("The merchant's eyes light up as he sees you approach.")
	fmt.Println(`"Well met, sir knight!" he exclaims as you reach his cart.`)
	fmt.Println(`"What can the humble Travers do for thee?"`)
	fmt.Println(`"What do you know of the dragon?"`)
	fmt.Println(`"Ah! That is a truly fearsome beast"`)
	fmt.Println(`"Many have tried to vanquish the tyrant, but alas, few have survived."`)
	fmt.Println(`"And as I understand it, none have even managed to scratch it."`)
	fmt.Println(`"There is only one weapon that can pierce it's iron hide."`)
	fmt.Println(`"Oh?' you say, 'And where might I find this weapon?"`)
	fmt.Println(`"I have it here, of course! There was a knight in the North that
	was gathering a force for an expedition to rid the land of the dragon.
	The night before he was to set out, he choked on a chicken bone and died.
	His family asked that I take the weapon and find someone to finish what he started."`)
	fmt.Println(`"Are you setting out in search of the dragon?"`)
	fmt.Println(`"Indeed. It the task to which the King has set my sword."`)
	fmt.Println(`"Well then, good knight! You must not go inadequately armed to such a fearsome foe."`)
	fmt.Println(`"This lance is your only hope of success."`)
	fmt.Println(`"And you are giving it to me freely?" you inquire.`)
	fmt.Println(`"Alas, brave knight, it has taken up precious room on my cart for many a moon"`)
	fmt.Println(`"I'm afraid I must charge something for the lost income while it has been in my care."`)
	fmt.Println()
	dk := getInput("He offers you a weapon named ‘Dragon Killer’. Will you buy it?")
	if dk == "no" {
		fmt.Println("You don't trust him, and decline the weapon. You leave town with just your sword.")
	} else if dk == "yes" {
		fmt.Println("Hoping you can trust him, you pay a small fortune for the spear.")
		player1.Inventory.Acquire("dragon killer")
	}

}
func innkeep() {
	fmt.Println()
	fmt.Println("-------------------")
	fmt.Println("You decide on the innkeeper.")
	fmt.Println("Surely, someone has told him something of the dragon.")
	fmt.Println("You sidle up to the bar and order a drink.")
	fmt.Println("As he sets the glass in front of you, you ask,")
	fmt.Println("'Tell me, good sir, know you anything about that fearsome dragon?'")
	fmt.Println("He tells of a man in a small village to the south who survived a dragon attack.")
	fmt.Println("You rush out of town through the south gate in search of the village.")
	fmt.Println("After many stops in every village you finally meet the man who survived.")
	fmt.Println("He tells you that he was one of 20 soldiers that attacked the dragon. None of them even scratched the beast.")
	brave := getInput("Do you still want to face the dragon? ")
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
		fmt.Println("next function called: os.Exit(0)")

		//showP1()
		os.Exit(0)
	}
}
func crier() {
	fmt.Println()
	fmt.Println("-------------------")
	fmt.Println("He talks for an hour, but provides no useful information.")
	fmt.Println("You feel your time has been wasted, but you must carry on.")
}
func townfolk() {
	fmt.Println("As you wander around the town, you see many people that might be able to help you.")
	fmt.Println("There's the town crier over there.")
	fmt.Println("Because he has to make announcements for the benefit of the town,")
	fmt.Println("he would likely be well informed on a diverse number of topics.")
	fmt.Println("You consider the blacksmith.")
	fmt.Println("Recently he has been less fun to be around, but he might have a useful tool lying about.")
	fmt.Println("The innkeeper has probably heard a lot of stories from various travelers.")
	fmt.Println("But how many of them are actually true?")
	fmt.Println("The most compelling stories are usually the most embellished.")
	fmt.Println("And yet, sometimes crucial nuggets of truth are contained in stories.")
	fmt.Println("And then there's the merchant. He certainly gets around.")
	fmt.Println("What has he seen and heard in his travels?")
	fmt.Println("More importantly, what has he got tucked away in that cart of his?")
}
