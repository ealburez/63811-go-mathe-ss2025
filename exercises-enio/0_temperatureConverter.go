package main

import (
	"fmt"
)

func farenheitToCelsius(farenheit float64) float64 {
	var celsius float64

	// to have a float division it was needed to add the numbers as float "9.0"
	celsius = 5.0/9.0 * (farenheit - 32)
	return celsius
}

func main() {
	var userInput float64 = 200

	fmt.Scan(&userInput)

	fmt.Println("Farenheit: ", userInput)
	fmt.Println("Celsius: ", farenheitToCelsius(userInput))

}
