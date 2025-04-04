package main

import (
	"fmt"
	"math"
)

func main() {
	var eingabe float64 = 1.0

	for eingabe != 0 {
		fmt.Println("Bitte eine positive Zahl eingeben (0 zum Beenden):")
		fmt.Scan(&eingabe)
		if eingabe > 0 {
			fmt.Printf("Die Wurzel von %v ist %v.\n", eingabe, math.Sqrt(eingabe))
		}
	}
}
