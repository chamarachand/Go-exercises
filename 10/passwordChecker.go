package main

import (
	"fmt"
	"unicode"
)

func main() {
	fmt.Print("Enter Password: ")
	var password string
	fmt.Scanln(&password)
	fmt.Println("Your password is " + passwordStrength(password) + "!")
}

func passwordStrength(password string) string {
	const LENGTH int = 8
	var hasAppropriateLength, hasUpper, hasLower, hasDigit bool

	if len(password) >= LENGTH {
		hasAppropriateLength = true
	}

	for _, ch := range password {
		switch {
		case unicode.IsUpper(ch):
			hasUpper = true
		case unicode.IsLower(ch):
			hasLower = true
		case unicode.IsDigit(ch):
			hasDigit = true
		}
	}

	var strength string

	if hasAppropriateLength && hasUpper && hasLower && hasDigit {
		strength = "Strong"
	} else {
		strength = "Weak"
	}

	return strength
}
