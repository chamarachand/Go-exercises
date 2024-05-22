package main

import "fmt"

func main() {
	fmt.Println("Celcius to Farenheit Convertor")

	fmt.Print("Temp (C): ")
	var tempCelcius float64
	_, err := fmt.Scanln(&tempCelcius)
	if err != nil {
		fmt.Println("Invalid input")
		return
	}

	tempFahrenheit := (tempCelcius * 5 / 9) + 32
	fmt.Printf("%.2f C: %.2f", tempCelcius, tempFahrenheit)
}
