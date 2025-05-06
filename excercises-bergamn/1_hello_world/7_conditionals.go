//From:https://github.com/bregman-arie/go-exercises/blob/main/exercises/conditionals/exercise.md

package main

import (
	"fmt"
	"math/rand"
)
func main(){
	randomInt := rand.Intn(100) + 1
	split :=50
	even:= randomInt%2 == 0
	
	fmt.Println(randomInt)

	if randomInt > split{
		fmt.Print("It's closer to 100")
	} else if randomInt < split {
		fmt.Print("It's closer to 0")
	}else{
		fmt.Print("It's ", split)
	}

	if even{
		fmt.Print(" and it's even")
	}

	fmt.Print(".\n")
}
