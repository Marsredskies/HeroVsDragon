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
	health, damage int
}

var dragon = player{100, 30} // creating dragon

var hero = player{100, 20} // creating player

var moves int = 1 // moves counter

var healAbility int = 4 //number of times player can regenerate his health

// gameStatus show current game stats
func gameStatus() {
	if hero.health < 0 {
		hero.health = 0
	}
	if dragon.health < 0 {
		dragon.health = 0
	}
	fmt.Printf("\n")
	fmt.Println("# +++++++++++++++++++++++++++++++")
	fmt.Println("# Move:", moves)
	fmt.Println("# Hero's health:", hero.health)
	fmt.Println("# Dragon's health:", dragon.health)
	fmt.Println("# +++++++++++++++++++++++++++++++")
	fmt.Printf("\n")

}

// userInput reads input from keyboard and returns int number. Prints a error message in case of non-number input
func userInput() int {
	fmt.Println("Input 1 to attack, 2 to heal")
	reader := bufio.NewReader(os.Stdin)   //read input from keyboard
	input, err := reader.ReadString('\n') //reads string until \n
	if err != nil {
		fmt.Println(err)
	}
	input = strings.TrimSpace(input)             //remove \n from the string
	number, err := strconv.ParseFloat(input, 64) //convert string to float64
	if err != nil {
		fmt.Println("enter a number")
	}
	return int(number)
}

// heroTurn makes player's move based on the choice. He can attack with 50% chance or heal himself.
func heroTurn() {
	var action int //action choice: 1 - attack, 2 - regen

	for {
		action = userInput()
		if action == 1 || action == 2 {
			break
		}

		continue
	}

	switch action {
	case 1:
		seconds := time.Now().Unix() //seed
		rand.Seed(seconds)
		chance := rand.Intn(10)

		if chance <= 5 {
			dragon.health -= hero.damage //attacks with 50% chance
			fmt.Println("You've hit the Dragon! He lost 20hp")
		} else {
			fmt.Println("You missed!")
		}

	default:
		if hero.health <= 80 && healAbility > 0 {
			hero.health += 20
			healAbility--
			fmt.Println("You've regenerated 20hp!", healAbility, "regens left")
		} else {
			fmt.Println("You can't heal now")

		}
	}

}

// dragonTurn makes dragon's move. Dragon attacks with 40% chance
func dragonTurn() {
	nanoseconds := time.Now().UnixNano()
	rand.Seed(nanoseconds) //seed
	chance := rand.Intn(10)

	if chance <= 4 { //attacks with 40%chance
		hero.health -= dragon.damage
		fmt.Println("Dragon hits you! You lose 30hp")
	} else {
		fmt.Println("Dragon missed!")
	}
}

func main() {
	fmt.Println("You are face to face with an ancient dragon!")
	fmt.Println("Defeat him!")
	fmt.Printf("\n")
	// game cycle
	for {
		heroTurn()
		dragonTurn()
		gameStatus()

		if hero.health <= 0 {
			fmt.Println("You died!")
			fmt.Println("Fight lasted for ", moves, "moves")
			break
		} else if dragon.health <= 0 {
			fmt.Println("Congrats, Hero! You have defeated the ancient dragon!")
			fmt.Println("Fight lasted for ", moves, "moves")
			break
		} else if dragon.health <= 0 && hero.health <= 0 {
			fmt.Println("Congrats, Hero! You have defeated the ancient dragon!")
			fmt.Println("Fight lasted for ", moves, "moves")
			break
		}
		moves++
		continue
	}
}
