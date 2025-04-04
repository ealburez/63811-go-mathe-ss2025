package main

import "fmt"

// Deklaration einer Funktion zur
// Berechnung der Fläche eines Rechtecks aus seinen Seitenlängen
func rechteck(seiteA int, seiteB int) int {
	var fläche int = seiteA * seiteB
	return fläche
}

// Deklaration einer Funktion zur
// Berechnung des Volumens eines Quaders aus Grundfläche und Höhe
func quader(grundfläche int, höhe int) int {
	var volumen int = grundfläche * höhe
	return volumen
}

// Deklaration der obligatorischen Funktion main()
func main() {
	var a int = 3
	var b int = 4
	var c int = 5
	fmt.Println("Fläche eines", a, "x", b, "Rechtecks:", rechteck(a, b))
	fmt.Println("Volumen eines", a, "x", b, "x", c, "Quaders:", quader(rechteck(a, b), c))
}
