//FROML: https://github.com/bregman-arie/go-exercises/blob/main/exercises/variables/exercise.md

package main

import(
	"fmt"
)

func main(){
	var userName, userAge string
	fmt.Print("Enter your name: ")
	fmt.Scan(&userName)
	fmt.Print("/n Enter your age: ")
	fmt.Scan(&userAge)
	fmt.Printf("Your name is %v and your age is %v \n", userName, userAge)
}