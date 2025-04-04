package main

import "fmt"

func f() bool {
	fmt.Println("<<Funktion f wird aufgerufen>>")
	return false
}

func main() {
	fmt.Println("AND Auswertung")
	fmt.Println("ergibt", true && f()) // f() wird ausgewertet
	fmt.Println()

	fmt.Println("OR Auswertung")
	fmt.Println("ergibt", true || f()) // f() wird nicht ausgewertet
}
