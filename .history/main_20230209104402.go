package main

import (
	"fmt"
	"movie/testPackage"
)

func main() {
	fmt.Println("Hello word")
	fmt.Println(testPackage.Greet())
	fmt.Println(testPackage.Area(4, 5))
	fmt.Println(testPackage.Calculate(5000.0, 10.0, 1.0))
}
