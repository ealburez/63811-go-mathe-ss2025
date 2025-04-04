package main

import "fmt"

func main() {

	var jahr int
	var schaltjahr bool

	fmt.Print("Jahr eingeben: ")
	fmt.Scan(&jahr)

	if jahr < 1582 {
		fmt.Println("Hinweis: Die Gregorianische Schaltjahr-Regel gilt erst ab dem Jahr 1582.")
	}

	// Schaltjahr-Regel des Gregorianischen Kalenders
	if jahr%400 == 0 {
		schaltjahr = true
	} else if jahr%100 == 0 {
		schaltjahr = false
	} else if jahr%4 == 0 {
		schaltjahr = true
	} else {
		schaltjahr = false
	}

	if schaltjahr {
		fmt.Print(jahr, " ist ein Schaltjahr.")
	} else {
		fmt.Print(jahr, " ist kein Schaltjahr.")
	}

}
