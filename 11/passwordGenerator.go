package main

import (
	"crypto/rand"
	"fmt"
	"math/big"
)

const (
	Lowercase    = "abcdefghijklmnopqrstuvwxyz"
	Uppercase    = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Numbers      = "0123456789"
	SpecialChars = "!@#$%^&*()-_=+[]{}|;:,.<>/?"
)

func main() {
	var length int
	var lowercase, uppercase, numbers, special bool

	fmt.Println("Enter password length:")
	fmt.Scan(&length)
	if length <= 0 {
		fmt.Printf("Length Cannot be zero or less")
		return
	}

	fmt.Print("Include lowercase letters? (true/false): ")
	fmt.Scan(&lowercase)

	fmt.Print("Include uppercase letters? (true/false): ")
	fmt.Scan(&uppercase)

	fmt.Print("Include numbers? (true/false): ")
	fmt.Scan(&numbers)

	fmt.Print("Include special characters? (true/false): ")
	fmt.Scan(&special)

	password := createPassword(length, lowercase, uppercase, numbers, special)
	fmt.Println("Password: ", password)
}

func createPassword(length int, lowercase, uppercase, numbers, special bool) string {
	var chars string

	if lowercase {
		chars += Lowercase
	}
	if uppercase {
		chars += Uppercase
	}
	if numbers {
		chars += Numbers
	}
	if special {
		chars += SpecialChars
	}

	password := make([]byte, length)
	for i := range password {
		randomIndex, err := rand.Int(rand.Reader, big.NewInt(int64(len(chars))))
		if err != nil {
			return "Error"
		}
		password[i] = chars[randomIndex.Int64()]
	}

	return string(password)
}
