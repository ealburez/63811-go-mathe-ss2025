package main

import "fmt"

// fNachC rechnet eine Temperatur von Grad Fahrenheit nach Grad Celsius um.
func fNachC(tempF float64) float64 {
	return (tempF - 32) / 1.8 // Temperatur in Grad Celsius
}

func main() {
	var tempF0 float64 = -459.67 // absoluter Nullpunkt
	var tempF1 float64 = 0       // spezielle Kältemischung
	var tempF2 float64 = 32      // Gefrierpunkt von Wasser
	var tempF3 float64 = 96      // Körpertemperatur des Menschen

	fmt.Println(tempF0, "°F entsprechen", fNachC(tempF0), "°C.")
	fmt.Println(tempF1, "°F entsprechen", fNachC(tempF1), "°C.")
	fmt.Println(tempF2, "°F entsprechen", fNachC(tempF2), "°C.")
	fmt.Println(tempF3, "°F entsprechen", fNachC(tempF3), "°C.")
}
