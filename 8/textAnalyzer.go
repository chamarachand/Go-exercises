package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

func main() {
	fmt.Print("Enter text: ")

	reader := bufio.NewReader(os.Stdin)
	userText, _ := reader.ReadString('\n')
	userText = strings.TrimSpace(userText)

	fmt.Println("Number of characters:", len(userText))
	fmt.Println("Number of words:", len(strings.Split(userText, " ")))

	fmt.Print("Enter the letter to count: ")
	letter, _ := reader.ReadString('\n')
	letter = strings.TrimSpace(letter)

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
