package main

import (
	"fmt"
	"math"
)

func main() {
	var principal, interestRate float64
	var loanTerm int

	fmt.Println("Principle: ")
	_, err := fmt.Scanln(&principal)
	if err != nil {
		error("Invalid input")
		return
	}

	fmt.Println("Interest rate: ")
	_, err = fmt.Scanln(&interestRate)
	if err != nil {
		error("Invalid input")
		return
	}

	fmt.Println("loan term (months): ")
	_, err = fmt.Scanln(&loanTerm)
	if err != nil {
		error("Invalid input")
		return
	}

	monthlyInterestRate := interestRate / 12.0
	payment := principal * monthlyInterestRate * math.Pow(1+monthlyInterestRate, float64(loanTerm)) / (math.Pow(1+monthlyInterestRate, float64(loanTerm)) - 1)
	fmt.Println("Monthly loan payment: ", payment)
}

func error(msg string) {
	fmt.Println(msg)
}
