package main

import (
	"fmt"
	"math"
)

func input_data (var_name string) float64 {
	var var_value float64
	
	fmt.Println("Enter side a of", var_name)
	_,err := fmt.Scan(&var_value)

	if err != nil{
		fmt.Println("Error: ", err)
	}

	return var_value
}

func main(){

	a := input_data("a")
	b := input_data("b")
	c := input_data("c")
	d := math.Sqrt(a*a + b*b + c*c) //Raumdiagonale

	format:=" %.2f\n"

	fmt.Printf("Volumen:"+ format, a * b * c)
	fmt.Printf("Kantensumme:" + format, 4*(a+b+c))
	fmt.Printf("Oberflache:" + format, 2*(a*b + a*c + b*c))
	fmt.Printf("Umkugelradius:" + format, d/2)
	fmt.Printf("Raumdiagonale:" + format, d)
}