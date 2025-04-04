package main

import "fmt"

// liefert eine Kopie des übergebenen Arrays zurück,
// in der alle negativen Elemente auf 0 gesetzt sind
func negToZeroValue(a [10]int) [10]int {
	for i := range a {
		if a[i] < 0 {
			a[i] = 0
		}
	}
	return a
}

func main() {
	nums := [10]int{8, -5, 6, -3, -7, 1, -4, 3, -9, 2}

	fmt.Println(nums) // [8 -5 6 -3 -7 1 -4 3 -9 2]
	posNums := negToZeroValue(nums)
	fmt.Println(nums)    // [8 -5 6 -3 -7 1 -4 3 -9 2]
	fmt.Println(posNums) // [8 0 6 0 0 1 0 3 0 2]
}
