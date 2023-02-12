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
	fmt.Println()
	n := 5
	for i := 0; i < n; i++ {
		for j := 0; j <= i; j++ {
			fmt.Print("*")
		}
		fmt.Println()
	}

	//labels: dung vong lap ben ngoai

outer:
	for i := 0; i < 3; i++ {
		for j := 1; j < 4; j++ {
			fmt.Printf("i = %d , j = %d\n", i, j)
			if i == j {
				break outer
			}
		}
	}
	i := 0
	for i <= 10 { //semicolons are ommitted and only condition is present
		fmt.Printf("%d ", i)
		i += 2
	}
}
