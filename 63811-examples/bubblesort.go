package main

// sortiert das Verbund-Array aufsteigend nach dem Feld name
func bubblesort(a *[10]person) {
	// äußere Schleife: noch zu sortierende Elemente
	for n := len(a); n > 1; n-- {
		// innere Schleife: Sortieren durch Vertauschen
		for i := 0; i < n-1; i++ {
			if a[i].name > a[i+1].name {
				// vertausche benachbarte Elemente
				a[i], a[i+1] = a[i+1], a[i]
			}
		}
	}
}
