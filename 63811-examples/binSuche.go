package main

const arrlen = 10 // Länge des Arrays

type person struct {
	name          string
	telefonnummer int
}

// Die rekursive Funktion sucht im Verbund-Array nach einem Element,
// dessen Feld name den Wert suchwert hat und gibt bei Erfolg
// den Wert des Felds telefonnummer dieses Elements zurück;
// ansonsten den Wert -1, wenn suchwert nicht gefunden wurde.
//
// Das Array muss aufsteigend sortiert sein und
// suchwert darf im Array höchstens einmal vorkommen.
//
// links und rechts sind die Grenzen des Suchbereichs,
// der beide Grenzen mit umfasst.
//
// Der initiale Aufruf erfolgt mit dem kompletten Array,
// also links=0 und rechts=arrlen-1.
func binSuche(personArr *[arrlen]person, links, rechts int,
	suchwert string) int {
	var ergebnis int
	// prüfen, ob Suchbereich zulässig ist
	if rechts >= links {
		// Mitte des Suchbereichs berechnen
		// (Index-Position im gesamten Array)
		mitte := links + (rechts-links)/2
		if personArr[mitte].name == suchwert {
			// suchwert gefunden
			ergebnis = personArr[mitte].telefonnummer
		} else {
			if personArr[mitte].name > suchwert {
				// Suche im linken Teilbereich fortsetzen
				ergebnis = binSuche(personArr, links, mitte-1,
					suchwert)
			} else {
				// Suche im rechten Teilbereich fortsetzen
				ergebnis = binSuche(personArr, mitte+1, rechts,
					suchwert)
			}
		}
	} else {

		ergebnis = -1 // suchwert nicht gefunden
	}
	return ergebnis
}
