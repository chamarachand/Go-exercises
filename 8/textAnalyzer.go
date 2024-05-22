package main

import (
	"fmt"
	"strings"
)

func main() {
	var userText string
	fmt.Print("Enter text: ")
	fmt.Scanln(&userText)

	fmt.Println("Number of characters:", len(userText))
	fmt.Println("Number of words:", len(strings.Split(userText, " ")))

	var letter string
	fmt.Print("Enter the letter to count: ")
	fmt.Scanf("%s", &letter)

	if letter == "" || len(letter) > 1 {
		fmt.Println("Invalid input")
		return
	}

	count := 0
	for _, ch := range userText {
		if string(ch) == letter {
			count++
		}
	}

	fmt.Printf("Count of the letter %s: %d", letter, count)
}
