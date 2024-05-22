package main

import (
	"fmt"
)

func main() {
	var name string
	fmt.Print("Enter your name: ")
	fmt.Scanln(&name)

	if name == "" {
		fmt.Println("Name cannot be empty. Please enter your name next time")
		return
	}

	fmt.Println("Hello " + name)
}
