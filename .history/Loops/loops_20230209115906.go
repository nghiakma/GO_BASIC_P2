package main

import (
	"fmt"
)

func main() {
	for i := 1; i < 10; i++ {
		if i%2 == 0 {
			continue
		}
		fmt.Printf("%d", i)
	}

	n := 5
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}
}
