package main

import(
	"fmt"
)

func main(){

	var name string
	var age int

	fmt.Print("What's your name: ")
	fmt.Scan(&name)
	fmt.Print("\nWhat's your age: ")
	fmt.Scan(&age)

	fmt.Printf("Your name is %v and your age is %v \n",name ,age)
}