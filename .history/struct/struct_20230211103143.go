package main

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

}
