package main

import "fmt"

// setzt in dem per Zeiger referenzierten Array
// alle negativen Elemente auf 0
func negToZeroPointer(pa *[10]int) {
	for i := range *pa {
		if (*pa)[i] < 0 {
			(*pa)[i] = 0
		}
	}
}

func main() {
	nums := [10]int{8, -5, 6, -3, -7, 1, -4, 3, -9, 2}

	fmt.Println(nums) // [8 -5 6 -3 -7 1 -4 3 -9 2]
	negToZeroPointer(&nums)
	fmt.Println(nums) // [8 0 6 0 0 1 0 3 0 2]
}
