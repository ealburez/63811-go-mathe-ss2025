package main

import "fmt"

func main() {
	// Deklaration einer Variablen vom Typ float64
	// und implizite Initialisierung mit Null-Wert
	var tempC float64

	// Deklaration einer Konstanten vom Typ float64
	const nullpunkt float64 = -273.15

	tempC = 10
	fmt.Println("Temperatur zum Start     (°C):", tempC)
	tempC = tempC + 12
	fmt.Println("Temperatur nach Änderung (°C):", tempC)

	var tempF float64
	tempF = 32 + tempC*1.8 // Temperatur in Grad Fahrenheit

	// Deklaration einer Variablen und
	// direkte Initialisierung mit einem Wert
	var tempK float64 = tempC - nullpunkt // Temperatur in Kelvin

	/* Ausgabe der Temperatur in den Einheiten:
	   - Celsius
	   - Fahrenheit
	   - Kelvin
	*/
	var tempText string = "Temperatur in "
	fmt.Println()
	fmt.Println(tempText+"°C:", tempC)
	fmt.Println(tempText+"°F:", tempF)
	fmt.Println(tempText+" K:", tempK)
}
