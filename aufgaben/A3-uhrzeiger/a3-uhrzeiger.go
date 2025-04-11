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
	hour_degree := 360.0 * (float64(hour_12h) / 12) + 30.0 * (float64(minute)/60.0)

	fmt.Printf("Time: %02d:%02d\n", hour, minute)
	fmt.Printf("Angulo del horario: %3.2f°\n", hour_degree)
	fmt.Printf("Angulo del minutero: %3.2f°\n", minute_degree)

}