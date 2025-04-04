package main

import "fmt"

var klima string // globale Variable (auf package-level)

func f() {
	wetter := "Regen" // lokale Variable
	fmt.Printf("Ausgabe f:    " + klima + ", " + wetter + "\n")
	klima = "warm"
	wetter = wetter + "schauer"
	fmt.Printf("Ausgabe f:    " + klima + ", " + wetter + "\n")
}

func main() {
	wetter := "Sonne" // lokale Variable
	klima = "kalt"
	fmt.Printf("Ausgabe main: " + klima + ", " + wetter + "\n")
	f()
	wetter = wetter + "nschein"
	fmt.Printf("Ausgabe main: " + klima + ", " + wetter + "\n")
}
