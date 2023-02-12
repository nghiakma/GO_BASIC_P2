package main

import (
	"fmt"
	"log"
	"movie/testPackage"
)

var p, r, t = 5000.0, -10.0, 1.0

func init() {
	println("Main package initialized")
	if p < 0 {
		log.Fatal("Principal is less than zero")
	}
	if r < 0 {
		log.Fatal("Rate of interest is less than zero")
	}
	if t < 0 {
		log.Fatal("Duration is less than zero")
	}
}
func main() {
	fmt.Println("Hello word")
	fmt.Println(testPackage.Greet())
	fmt.Println(testPackage.Area(4, 5))
	fmt.Println(testPackage.Calculate(p, r, t))
	employeeSalary := map[string]int{
		"steve": 12000,
		"jamie": 15000,
	}
	newEmp := "steve"
	value, ok := employeeSalary[newEmp]
	fmt.Println(value)
	fmt.Println(ok)

}
