package main

import(
	"fmt"
	//"math"
)

func main() {
	var hour,minute int
	 
	fmt.Println("Enter the hour followed by the minutes (24 hour format)")
	_,err:=fmt.Scan(&hour, &minute)

	if err!= nil{
		fmt.Println("Error: ", err)
		return
	}

	hour_12h := hour % 12
	

	minute_degree := 360.0 * (float64(minute)/60.0)
	hour_degree := 360.0 * (float64(hour) / 12) + 30.0 * (float64(minute)/60.0)

	fmt.Printf("%d and %d\n", hour_12h, minute)
	fmt.Printf("%.2f and %.2f", hour_degree, minute_degree)

}