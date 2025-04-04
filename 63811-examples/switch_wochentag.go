package main

import "fmt"

func main() {
	var num int
	var text string

	fmt.Scan(&num)

	switch num {
	case 1:
		text = "Montag"
	case 2:
		text = "Dienstag"
	case 3:
		text = "Mittwoch"
	case 4:
		text = "Donnerstag"
	case 5:
		text = "Freitag"
	case 6, 7:
		text = "Wochenende"
	default:
		text = "unbekannt"
	}

	fmt.Println(num, "ist", text, "zugeordnet.")
}
