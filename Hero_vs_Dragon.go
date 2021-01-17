package main

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"strconv"
	"strings"
	"time"
)

type player struct {
	health, armor, damage int
}

type enemy struct {
	health, damage int
}

type weapon struct {
	name      string
	damage    int
	hitChance int // successfull hit chance
	left      int // how many times it can be used
}

// creating weapons
var crossbow = weapon{"crossbow", 30, 80, 5}
var bfg = weapon{"BFG9000", 100, 60, 3}
var ancientSpell = weapon{"Ancient Spell", 250, 20, 1}

var dragon = enemy{400, 40}     // creating dragon
var hero = player{100, 100, 40} // creating player

var moves int = 0        // moves counter
var regenAbility int = 5 //number of times player can regenerate his health

var rageDamage int = 0    // if player's health is low, he deals amped damage
var rageHitChance int = 0 // if player's health is low, he's got more chances to hit dragon

// gameStatus show current game stats
func gameStatus() {
	if hero.health < 0 { // hero's health can't be below 0
		hero.health = 0
	}
	if dragon.health < 0 { // dragon's health can't be below 0 too
		dragon.health = 0
	}
	// stats table
	fmt.Printf("\n")
	fmt.Println("# +++++++++++++++++++++++++++++++")
	fmt.Printf("# Move: %v\n", moves)
	fmt.Printf("# Hero's hp/armor: %v/%v \n", hero.health, hero.armor)
	fmt.Printf("# Dragon's hp: %v\n", dragon.health)
	fmt.Println("# +++++++++++++++++++++++++++++++")
	fmt.Printf("# BFG shots left: %v\n", bfg.left)
	fmt.Printf("# Crossbow shots left: %v\n", crossbow.left)
	fmt.Printf("# Ancient Spell: %v\n", ancientSpell.left)
	fmt.Println("# +++++++++++++++++++++++++++++++")
	fmt.Printf("\n")
}

func printInventory() {
	// inventory table/input instructions
	fmt.Println("\t\t\tInput:")
	fmt.Println("1 to attack with your sword. -5 damage with every successful hit.")
	fmt.Printf("2 to fire crossbow. %v%% hit chance, %v hp damage\n", crossbow.hitChance, crossbow.damage)
	fmt.Printf("3 to fire BFG9000. %v%% hit chance, %v hp damage\n", bfg.hitChance, bfg.damage)
	fmt.Printf("4 to use Ancient Spell. %v%% hit chance, 250hp damage!\n", ancientSpell.hitChance)
	fmt.Println("5 to regen 20hp")
	fmt.Printf("\n")
}

// userInput reads input from keyboard and returns an int number. Prints a error message in case of non-number input
func userInput() int {

	reader := bufio.NewReader(os.Stdin)   //read input from keyboard
	input, err := reader.ReadString('\n') //reads string until \n
	if err != nil {
		fmt.Println(err)
	}
	input = strings.TrimSpace(input)             //remove \n from the string
	number, err := strconv.ParseFloat(input, 64) //convert string to float64
	if err != nil {
		fmt.Println("not a number")
	}
	return int(number)
}

// heroTurn makes player's move based on the choice. He can attack with 50% chance or heal himself.
func heroTurn() {
	var action int        //action choice: 1 - 4 attack, 5 - regen
	var damageDone int    //damage value depending on weapon choice
	var hitChance int     //weapon hit chance depending on weapon choice
	var weaponName string //weapon hit chance depending on weapon choice
	var tired int         //hero get's more tired with every move. Substracted from the hit chance

	printInventory()

	for {
		for { //action choise. input repeats until number from 1 to 5 is entered
			action = userInput()
			if action >= 1 || action <= 5 {
				break
			}

			continue
		}
		// switch case for weapon choice
		switch action {
		case 1: // standard sword attack
			damageDone = hero.damage
			hitChance = 50 - tired
			weaponName = "your sword"
			break
		case 2: //crossbow
			weaponName = crossbow.name
			damageDone = crossbow.damage
			hitChance = crossbow.hitChance
			crossbow.left-- // -1 use
			break
		case 3: //BFG 9000
			if bfg.left > 0 {
				weaponName = bfg.name
				damageDone = bfg.damage
				hitChance = bfg.hitChance - tired
				bfg.left-- // -1 use
				break
			} else {
				fmt.Println("can't use BFG9000 again, choose another weapon")
				continue
			}

		case 4: //ancient spell
			if ancientSpell.left > 0 {
				weaponName = ancientSpell.name
				damageDone = ancientSpell.damage
				hitChance = ancientSpell.hitChance - tired
				ancientSpell.left-- // -1 use
				break
			} else {
				fmt.Println("can't use Ancient Spell again, choose another weapon")
				continue
			}

		case 5: // regen health
			if hero.health <= 80 && regenAbility > 0 {
				hero.health += 20
				regenAbility-- // -1 to regen ability
				if hero.health >= 50 {
					rageDamage = 0
					rageHitChance = 0
					fmt.Println("Hero healed, no rage damage")
				}
				fmt.Println("You've regenerated 20hp!", regenAbility, "regens left")
			} else if regenAbility <= 0 {
				fmt.Println("You can't heal anymore")
			} else {
				fmt.Println("You can't heal now")
			}
			break
		default: // input again
			fmt.Println("Enter a number from 1 to 5")
			continue
		}
		break
	}
	//seeding random number generator, result is random a number between 1 and 100
	seconds := time.Now().Unix()
	rand.Seed(seconds)
	result := rand.Intn(100)

	//damaging dragon according to player's choice and weapon's hit chance
	if result <= hitChance {
		dragon.health -= damageDone
		fmt.Printf("You've hit the Dragon with %v! He lost %v hp!\n", weaponName, damageDone)
	} else if damageDone != 0 {
		fmt.Printf("You attacked dragon with %v and missed!\n", weaponName)
	}

	// everytime hero hits dragon, his sword deals less damage
	if dragon.health-damageDone == dragon.health-hero.damage {
		hero.damage -= 5
		if hero.damage < 0 { //damage can't be negative
			hero.damage = 0
		}
	}
	moves++    // moves counter
	tired += 5 // hero get's tired
}

// dragonTurn makes dragon's move. Dragon attacks with 40% chance
func dragonTurn() {
	nanoseconds := time.Now().UnixNano()
	rand.Seed(nanoseconds) //seed
	chance := rand.Intn(100)

	if chance <= 40 { // attacks with 40%c hance
		if hero.armor > 0 {
			hero.armor -= dragon.damage
			if hero.armor <= 0 {
				hero.armor = 0
			}
			fmt.Printf("Dragon hits you! Shield -%v hp\n", dragon.damage) // armor to be destroyed first
		} else if hero.armor == 0 {

			hero.health -= dragon.damage
			fmt.Printf("Dragon hits you! Health -%v hp\n", dragon.damage)
		}
	} else if dragon.health <= 0 { // in case if dragon dies, he can't make a move
		fmt.Println("Dragon screams and falls in fear...")
	} else {
		fmt.Println("Dragon missed!")
	}
}

// random events! Hero's diarrhea etc.
func randomEvents() {
	seconds := time.Now().Unix()
	rand.Seed(seconds / 2) // seeding randomizer
	eventChance := rand.Intn(8)

	switch eventChance {
	case 0:
		dragon.health -= 50
		fmt.Println("Lightning strikes the Dragon! He loses 50hp")
	case 3:
		if bfg.left == 0 && crossbow.left == 0 && ancientSpell.left == 0 {
			hero.health -= 30
			fmt.Println("Dragon hits hero with his tail! Hero loses 30hp")
		} else {
			bfg.left = 0
			crossbow.left = 0
			ancientSpell.left = 0
			fmt.Println("Dragon hits hero with his tail! Hero loses all his weapons but sword!")
		}
	case 6:
		fmt.Println("Player finds a better sword with 70hp damage!")
		hero.damage = 70
	}

}

//some game comments
func gameComments() {
	if hero.armor == 0 {
		fmt.Println("Dragon destroyed hero's armor!")
	}

	if hero.health <= 50 {
		fmt.Println("Hero is about to die! Heal him!")
	} else if hero.health <= 40 {
		rageDamage = 50
		rageHitChance = 10
		fmt.Println("Hero is in rage mode now! +50 to sword damage")
	}

	switch {
	case dragon.health <= 200:
		fmt.Println("Dragon is half-alive! Keep fighting him")
	case dragon.health <= 100:
		fmt.Println("Dragon's health is below 100, keep fighting!")
	case dragon.health <= 50:
		fmt.Println("Dragon is about to die!")
	}

}

func main() {
	fmt.Println("You are face to face with an ancient dragon!")
	fmt.Println("Defeat him!")

	// game cycle
	for {
		gameStatus()
		gameComments()
		heroTurn()
		dragonTurn()
		randomEvents()

		if hero.health == 0 {
			fmt.Println("You died!")
			fmt.Printf("Fight lasted for %v moves", moves)
			gameStatus()
			break
		} else if dragon.health == 0 {
			fmt.Println("Congrats, Hero! You have defeated the ancient dragon!")
			fmt.Printf("Fight lasted for %v moves\n", moves)
			gameStatus()
			break
		} else if dragon.health == 0 && hero.health == 0 {
			fmt.Println("You are dead. But you have defeated the ancient dragon!")
			fmt.Printf("Fight lasted for %v moves\n", moves)
			gameStatus()
			break
		}

		continue
	}
}
