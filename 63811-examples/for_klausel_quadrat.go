package main

import "fmt"

func main() {
	for i := 1; i <= 10; i++ {
		fmt.Printf("%2[1]v * %2[1]v ist %3v\n", i, i*i)
	}
}
