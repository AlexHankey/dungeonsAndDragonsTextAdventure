package main

import (
	"fmt"
	"math/rand"
	"os"
	"strings"
	"time"

	"github.com/common-nighthawk/go-figure"
)

/*
Player character
Name, Age, Height, Class, Backstory,
*/
var (
	characterName      string
	characterGender    string
	characterAge       uint
	characterHeight    uint
	characterClass     string
	characterBackstory string
)

/*
Enemy character
Health, Race, Weapon, ArmourClass
*/
type Enemy struct {
	Health      uint
	Race        string
	Weapon      string
	ArmourClass uint
}

func getEnemy() Enemy {
	return Enemy{
		Health:      20,
		Race:        "Goblin",
		Weapon:      "Club",
		ArmourClass: 10,
	}
}

func (e Enemy) getEnemyDamage() int {
	return 20
}

func diceRoll() {
	min := 0
	max := 20
	// set seed
	rand.Seed(time.Now().UnixNano())
	// generate random number and print on console
	fmt.Println(rand.Intn(max-min) + min)
}

func combat() {
	fmt.Println("What do you want to do?")
	var playerCombatChoice string
	fmt.Scan(&playerCombatChoice)
	/* BUG/ ISSUE:
	entering string of text is for some reason causing the programme
	to suddenly exit out when i want it to perform the dice roll function
	*/

	var action string
	fmt.Println("Roll combat damage dice? (y/n)")
	fmt.Scan(&action)
	var toLowString_action = strings.ToLower(action)
	if toLowString_action == "y" {
		diceRoll()
	} else {
		return
	}

	getEnemy()
}

func characterBuilder() {
	intro1 := figure.NewFigure("Welcome adventurer to...", "", true)
	intro2 := figure.NewFigure("Dungeons & Dragons", "", true)
	intro1.Print()
	intro2.Print()
	fmt.Println("Please enter your character Name:")
	fmt.Scan(&characterName)
	fmt.Printf("Is %v a male or a female?\n", characterName)
	fmt.Scan(&characterGender)
	fmt.Printf("Welcome %v, Happy to see you have joined us on this dangerous adventure. Why don't we get introductions out of the way and you tell me a little about yourself?\n", characterName)
	fmt.Printf("How old is %v?:\n", characterName)
	fmt.Scan(&characterAge)

	/* TO-DO:
	Character Races could be older than 100(Such as elfs)
	Include character Race structs to easily do checks against in-game universe races
	*/
	if characterAge < 18 {
		fmt.Printf("Sorry but %v is too young to be fighting monsters, come back and play when you're a bit older\n", characterName)
		os.Exit(404)
	} else if characterAge > 100 {
		if characterGender == "male" {
			fmt.Println("Sorry Gramps but you're just too old to be going around fighting monsters.")
			os.Exit(404)
		} else if characterGender == "female" {
			fmt.Println("Sorry Grandma but you're just too old to be going around fighting monsters")
			os.Exit(404)
		}
	}

	fmt.Printf("How tall is %v?:\n", characterName)
	fmt.Scan(&characterHeight)
	switch {
	case characterHeight > 7:
		fmt.Printf("WOW. %v you're mighty tall!!!\n", characterName)
	case characterHeight < 4:
		fmt.Printf("%v you sure you won't struggle on this adventure?\n", characterName)
	default:
		fmt.Println("You know what they say. Average height, you'll fit in just right.")
	}
	fmt.Println("What is your character class?:")
	fmt.Scan(&characterClass)

	/* TO-DO:
	Way to handler multiple lines of user input without the
	terminal / programme having a heart attack and dieing on me
	*/

	trimmed := strings.TrimSpace(characterClass)
	fmt.Printf("Ah! so you're a %v. Well someone like you should do just fine on this adventure.\n", trimmed)
	fmt.Println("Well then, Let's begin!...")
	fmt.Println("-----------------------------------------------------")
}

func firstFight(e Enemy) {
	// To-Do, add story about engaging with a Slime
	fmt.Println("Oh No, A Slime!")
	fmt.Println("Do you wish to engage? (y/n):")
	var userInput string
	fmt.Scan(&userInput)
	var convertedAnswer = strings.ToLower(userInput)

	if convertedAnswer == "y" {
		combatText := figure.NewFigure("Engaged in combat!", "", true)
		fmt.Println(combatText)
	} else if convertedAnswer == "n" {
		fmt.Println("You fled from the fight & your adventure is over!")
	} else {
		fmt.Println("Please enter a valid input adventurer")
	}
	combat()

	// e.getEnemyDamage()

	// goblin1 := Enemy{Health: 20, Race: "Goblin", Weapon: "Club", ArmourClass: 10}

	// todo, once i enter my backstory the terminal just exits out for some reason :/
}

func main() {
	characterBuilder()
	fmt.Println("You enter a dark and slimy cave when all of a sudden you hear a noise, you turn around and notice a large object in front of you...")
	firstFight(getEnemy())
}
