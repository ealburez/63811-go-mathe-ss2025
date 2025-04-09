package main

import(
	"fmt"
)

func main(){
	const inch_cr float64 = 2.54
	const mm_cr float64 = 0.001
	const km_cr float64 = 1000
	const seamiles_cr float64 = 1852
	const lightyear_cr float64 = 9_460_730_472_580_800
	var input_in_meter float64

	fmt.Println("Enter the length in meter:")
	fmt.Scan(&input_in_meter)
	fmt.Printf("%.3e mm\n", input_in_meter / mm_cr)
	fmt.Printf("%.3f km\n", input_in_meter / km_cr)
	fmt.Printf("%.3e in\n", input_in_meter / inch_cr)
	fmt.Printf("%.3f Sea Miles\n", input_in_meter / seamiles_cr)
	fmt.Printf("%.3e Light Years\n", input_in_meter / lightyear_cr)



}