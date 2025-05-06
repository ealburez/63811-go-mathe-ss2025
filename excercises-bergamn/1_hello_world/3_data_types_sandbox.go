// From: https://github.com/bregman-arie/go-exercises/blob/main/exercises/data_types/exercise.md
package main

import "fmt"

func main() {
    var food = "Pizza"
    var slices = 4
    var pineappleOnPizza = true

	fmt.Printf("%T %T %T", food,slices,pineappleOnPizza)
}