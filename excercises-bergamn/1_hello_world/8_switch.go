package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	rand.Seed(time.Now().UnixNano())
	randomInt := rand.Intn(11)
	fmt.Println("Random number: ", randomInt)
	switch randomInt {

	case 1:
		fmt.Println("One is the loneliest number that you'll ever do")
	case 2:
		fmt.Println("Two can be as bad as one. It's the loneliest number since the number one")
	default:
		fmt.Println("Sorry, no mention of the number in the song")
	}

}
