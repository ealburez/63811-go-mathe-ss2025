// FROM: https://github.com/bregman-arie/go-exercises/blob/main/exercises/arithmetic_operators/exercise.md

package main 

import(
	"fmt"
)

func main(){
	const maxPackages int = 100
	var deliveredPackages, customers, remainingPackages int

	fmt.Printf("I have to deliver a maximum of %v pacckages. \n ", maxPackages)
	fmt.Print("How many packages have been delivered?: ")
	fmt.Scan(&deliveredPackages)
	remainingPackages = maxPackages - deliveredPackages
	fmt.Printf("There are %v remaining packages to deliver. \n ", remainingPackages)
	fmt.Print("How many customers are left?: ")
	fmt.Scan(&customers)
	fmt.Printf("In average each customer will receive %v packages \n", remainingPackages/customers)
}