package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/common-nighthawk/go-figure"
)

// Name, Age, Height, Class, Backstory,

// type Player struct {
// 	Name, Class, Backstory string
// 	Age, Height            uint
// }

// Player character
var characterName string
var characterAge uint
var characterHeight uint
var characterClass string
var characterBackstory string

// !!! Need to figure out how to give player option of what classes they can choose from

// Health, Race, Weapon, ArmourClass
type Enemy struct {
	Health      uint
	Race        string
	Weapon      string
	ArmourClass uint
}

func actOneStory() {

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
	var action string
	fmt.Println("Roll dice? (y/n)")
	fmt.Scan(&action)
	var toLowString_action = strings.ToLower(action)
	if toLowString_action == "y" {
		diceRoll()
	} else {
		return
	}

}

func welcomeIntro() {
	intro1 := figure.NewFigure("Welcome adventurer to...", "", true)
	intro2 := figure.NewFigure("Dungeons & Dragons", "", true)
	intro1.Print()
	intro2.Print()
	fmt.Println("Please enter your character Name:")
	fmt.Scan(&characterName)
	fmt.Printf("Welcome %v, Happy to see you have joined us on this dangerous adventure. Why don't we get introductions out of the way and you tell me a little about yourself?\n", characterName)
	fmt.Printf("How old is %v?:\n", characterName)
	fmt.Scan(&characterAge)
	fmt.Printf("How tall is %v?:\n", characterName)
	fmt.Scan(&characterHeight)
	fmt.Println("What is your character class?:")
	fmt.Scan(&characterClass)
	fmt.Printf("And what is your backstory %v?\n", characterName)
	fmt.Scan(&characterBackstory)
	fmt.Printf("Ah! so you're a %v. Well someone like you should do just fine on this adventure.\n", characterClass)
	fmt.Println("Well then, Let's begin!...")
	fmt.Println("-----------------------------------------------------")
}
func firstFight(e Enemy) {
	// To-Do, add story about engaging with a goblin
	fmt.Println("oh no a goblin!")
	fmt.Println("Do you wish to engage? (y/n):")
	var userInput string
	fmt.Scan(&userInput)
	var convertedAnswer = strings.ToLower(userInput)

	if convertedAnswer == "y" {
		combatText := figure.NewFigure("Engaged in combat!", "", true) // Not working at the moment
		fmt.Println(combatText)
	} else if convertedAnswer == "n" {
		fmt.Println("You fled from the fight & your adventure is over!")
	} else {
		fmt.Println("Please enter a valid input adventurer")
	}
	combat()

	e.getDamage()

	// goblin1 := Enemy{Health: 20, Race: "Goblin", Weapon: "Club", ArmourClass: 10}

	// todo, once i enter my backstory the terminal just exits out for some reason :/
}

func main() {
	welcomeIntro()
	firstFight(getEnemy()) // Not sure if this argument is correct?

}

func getEnemy() Enemy {
	return Enemy{
		Health:      20,
		Race:        "Goblin",
		Weapon:      "Club",
		ArmourClass: 10,
	}
}

func (e Enemy) getDamage() int {
	return 20
}
