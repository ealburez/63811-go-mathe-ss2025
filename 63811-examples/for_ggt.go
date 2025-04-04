package main

import "fmt"

// ggT von n und m mittels Euklidischen Algorithmus berechnen
func ggT(n, m int) int {
	for n != m {
		if n > m {
			n -= m
		} else {
			m -= n
		}
	}
	return n
}

func main() {
	var n, m int
	fmt.Println("Bitte zwei natürliche Zahlen größer 0 eingeben:")
	fmt.Scan(&n)
	fmt.Scan(&m)
	if n > 0 && m > 0 {
		fmt.Printf("%v ist der ggT von %v und %v.\n", ggT(n, m), n, m)
	} else {
		fmt.Printf("Die Eingabe %v und %v ist unzulässig.\n", n, m)
	}
}
