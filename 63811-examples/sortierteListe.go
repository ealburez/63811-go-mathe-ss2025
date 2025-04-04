package main

import "fmt"

// Verbund-Typ für ein Listen-Element
type listenElement struct {
	wert     int
	nächstes *listenElement
}

// Verbund-Typ für eine Liste
type sortierteListe struct {
	kopf *listenElement
}

// gibt alle Werte der Liste der Reihe nach aus
func listeAusgeben(liste *sortierteListe) {
	element := liste.kopf
	fmt.Printf("-> ")
	for element != nil {
		fmt.Printf("%d -> ", element.wert)
		element = element.nächstes
	}
	fmt.Printf("nil\n")
}

// liefert die Länge der Liste (d.h. die Anzahl ihrer Elemente) zurück
func längeDerListe(liste *sortierteListe) int {
	länge := 0
	element := liste.kopf
	for element != nil {
		länge++
		element = element.nächstes
	}
	return länge
}

// fügt ein neues Element mit dem Wert neuerWert so in eine
// aufsteigend sortierte Liste ein, dass die Liste danach
// weiterhin aufsteigend sortiert ist
func elementEinfügen(liste *sortierteListe, neuerWert int) {

	// wenn die Liste noch leer ist ODER
	// neuerWert kleiner-gleich dem Wert im Kopf der Liste ist,
	// wird das neue Element zum Kopf der Liste
	if liste.kopf == nil || neuerWert <= liste.kopf.wert {
		if liste.kopf == nil {
			fmt.Printf("Wert %v in leere Liste einfügen\n", neuerWert)
		} else {
			fmt.Printf("Wert %v am Anfang vor %v einfügen\n",
				neuerWert, liste.kopf.wert)
		}

		// neues Element vor dem bisherigen Listen-Kopf einfügen
		neues := &listenElement{wert: neuerWert} // Element erzeugen
		neues.nächstes = liste.kopf
		liste.kopf = neues

		return
	}

	aktuelles := liste.kopf

	// Liste durchlaufen, um passende Einfüge-Position zu finden:
	// solange das aktuelle Element einen Nachfolger hat UND
	// der Wert dieses Nachfolgers kleiner als neuerWert ist,
	// weiter zum nächsten Element
	for aktuelles.nächstes != nil &&
		aktuelles.nächstes.wert < neuerWert {
		aktuelles = aktuelles.nächstes
	}

	if aktuelles.nächstes == nil {
		fmt.Printf("Wert %v am Ende hinter %v einfügen\n",
			neuerWert, aktuelles.wert)
	} else {
		fmt.Printf("Wert %v zwischen %v und %v einfügen\n",
			neuerWert, aktuelles.wert, aktuelles.nächstes.wert)
	}

	// neues Element zwischen dem aktuellen Element und
	// seinem Nachfolger einfügen
	neues := &listenElement{wert: neuerWert} // Element erzeugen
	neues.nächstes = aktuelles.nächstes
	aktuelles.nächstes = neues

}
