package main

import "fmt"

type Employee struct {
	firstName string
	lastName  string
	age       int
	salary    int
}

func main() {
	emp1 := Employee{
		firstName: "Sam",
		age:       25,
		salary:    500,
		lastName:  "Anderson",
	}
	emp2 := Employee{"Thomas", "Paul", 29, 800}
	fmt.Println("Employee 1", emp1)
	fmt.Println("Employee 2", emp2)

	add := Address{
		city: "vinh",
		state: "illll",
		fmt.Println("Address 1", add.city);
		fmt.Println("Address 2", add.state);
	}
}
