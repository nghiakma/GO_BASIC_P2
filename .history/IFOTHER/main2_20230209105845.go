package main

import "fmt"

func main() {
	num := 10
	if num%2 == 0 {
		fmt.Println("The number ", num, "is even")
	} else {
		fmt.Println("The number ", num, "is odd")
	}

	num1 := 99
	if num1 <= 50 {
		fmt.Println(num, "is less than or equal to 50")
	} else if num1 >= 51 && num <= 100 {
		fmt.Println(num, "is between 51 and 100")
	} else {
		fmt.Println(num, "is greater than 100")
	}
}
