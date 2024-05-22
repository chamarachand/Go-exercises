package main

import (
	"fmt"
	"math/rand"
)

func main() {
	secretNumber := rand.Intn(10) + 1

	var userGuess int
	var correctGuess = false

	for !correctGuess {
		fmt.Print("Guess the number: ")
		fmt.Scanln(&userGuess)

		if userGuess == secretNumber {
			fmt.Println("You win!")
			correctGuess = true
		} else if userGuess < secretNumber {
			fmt.Println("Your guess is lower than the actual number")
		} else {
			fmt.Println("Your guess is higher than the actual number")
		}
	}
}
