package main

import (
	"fmt"
)

func main() {
	var num1, num2 float64
	fmt.Print("Enter num 1: ")
	_, err := fmt.Scanln(&num1)
	if err != nil {
		error("Invalid number")
		return
	}

	fmt.Print("Enter num 2: ")
	_, err = fmt.Scanln(&num2)
	if err != nil {
		error("Invalid number")
		return
	}

	var operation string
	fmt.Print("Enter the operation (+ - x %): ")
	fmt.Scanln(&operation)

	var answer float64
	switch operation {
	case "+":
		answer = add(num1, num2)
	case "-":
		answer = subtract(num1, num2)
	case "x":
		answer = multiply(num1, num2)
	case "/":
		if num2 == 0 {
			error("Division by zero error")
			return
		}
		answer = divide(num1, num2)
	default:
		fmt.Println("Invalid Operation")
		return
	}

	fmt.Println("Answer: ", answer)
}

func add(num1 float64, num2 float64) float64 {
	return num1 + num2
}

func subtract(num1 float64, num2 float64) float64 {
	return num1 - num2
}

func multiply(num1 float64, num2 float64) float64 {
	return num1 * num2
}

func divide(num1 float64, num2 float64) float64 {
	return num1 / num2
}

func error(msg string) {
	fmt.Println(msg)
}
