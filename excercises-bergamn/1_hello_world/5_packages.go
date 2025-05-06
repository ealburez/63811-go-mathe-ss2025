// From: https://github.com/bregman-arie/go-exercises/blob/main/exercises/packages/exercise.md
package main

import(
	"fmt"
	"time"
	"math"
	"math/rand"
)

func main(){

	var randomNumber int
	randomNumber= rand.Intn(101)

	fmt.Printf("The time now is: %v\n",time.Time(time.Now()))
	fmt.Printf("A random number is: %v \n", randomNumber)
	fmt.Printf("Sqr of a random number: %v \n", math.Sqrt(float64(randomNumber)))

}