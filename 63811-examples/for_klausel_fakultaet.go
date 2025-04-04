package main

import (
	"fmt"
	"math"
)

func main() {
	eingabe := 0
	überlauf := false

	fmt.Println("Fakultät berechnen (natürliche Zahl eingeben):")
	fmt.Scan(&eingabe)

	if eingabe < 0 {
		fmt.Printf("Eingabe %v ist ungültig.\n", eingabe)
	} else {
		fak := 1
		for i := 1; i <= eingabe; i++ {
			// vorab prüfen, ob Multiplikation ohne Überlauf möglich ist
			if math.MaxInt/fak >= i {
				fak *= i
			} else {
				fmt.Printf("Achtung: %v * %v führt zu Überlauf: %v\n", fak, i, fak*i)
				überlauf = true
				break // Schleife sofort verlassen
			}
		}
		if überlauf {
			fmt.Printf("Fakultät von %v ist zu groß, um mit int berechnet werden zu können.", eingabe)
		} else {
			fmt.Printf("Fakultät von %v ist %v.\n", eingabe, fak)
		}
	}

}
