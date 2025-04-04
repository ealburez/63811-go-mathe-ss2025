package main

import "fmt"

// Verbund-Typ für einen binären Knoten
type knoten struct {
	wert          int
	links, rechts *knoten
}

// Die rekursive Funktion fügt einen neuen Knoten mit dem Wert
// neuerWert so in einen binären Suchbaum ein,
// dass die Suchbaum-Eigenschaft danach weiterhin erfüllt ist.
// wurzel muss die Wurzel eines Suchbaums referenzieren,
// die bereits einen Wert enthält (d.h. wurzel darf nicht nil sein).
// neuerWert darf noch nicht im Suchbaum enthalten sein.
func einfügen(wurzel *knoten, neuerWert int) {
	if neuerWert < wurzel.wert {
		if wurzel.links == nil {
			// neuen Knoten erzeugen
			wurzel.links = &knoten{wert: neuerWert}
		} else {
			// rekursiv im linken Teilbaum nach
			// passender Einfüge-Position suchen
			einfügen(wurzel.links, neuerWert)
		}
	} else {
		if wurzel.rechts == nil {
			// neuen Knoten erzeugen
			wurzel.rechts = &knoten{wert: neuerWert}
		} else {
			// rekursiv im rechten Teilbaum nach
			// passender Einfüge-Position suchen
			einfügen(wurzel.rechts, neuerWert)
		}
	}
}

// Die rekursive Funktion sucht in einem Suchbaum
// nach dem Wert suchwert.
// wurzel muss die Wurzel eines Suchbaums referenzieren.
// Wenn suchwert im Suchbaum enthalten ist, wird ein Zeiger
// auf den entsprechenden Knoten zurückgeliefert, ansonsten nil.
func suchen(wurzel *knoten, suchwert int) *knoten {
	if wurzel == nil || wurzel.wert == suchwert {
		return wurzel
	}

	if suchwert < wurzel.wert {
		return suchen(wurzel.links, suchwert)
	} else {
		return suchen(wurzel.rechts, suchwert)
	}
}

// Die Funktion liefert true zurück, wenn der Wert suchwert
// in einem Suchbaum enthalten ist, ansonsten false.
// wurzel muss die Wurzel eines Suchbaums referenzieren.
func istWertInBaum(wurzel *knoten, suchwert int) bool {
	return suchen(wurzel, suchwert) != nil
}

// gibt die Werte des Suchbaums in aufsteigender Reihenfolge aus
func baumAusgeben(wurzel *knoten) {
	if wurzel != nil {
		// Traversierung des Baums in sog. In-Order (L-Knoten-R)
		baumAusgeben(wurzel.links)
		fmt.Printf("%d ", wurzel.wert)
		baumAusgeben(wurzel.rechts)
	}
}
