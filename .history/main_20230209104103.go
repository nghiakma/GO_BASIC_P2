package main

import (
	"fmt"
	"movie/testPackage"
)

func main() {
	fmt.Println("Hello word")
	fmt.Println(testPackage.Greet())
	fmt.Println(testPackage.Area(4, 5))
}
