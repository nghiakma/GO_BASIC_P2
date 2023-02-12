package main

import "fmt"

func main() {
	b := 255
	var a *int = &b
	fmt.Printf("Type of a is %T\n", a)
	fmt.Println("address of b is", a)

	//Passing pointer to a function
	a1 := 58
	fmt.Println("value of a before function call is", a)
	b1 := &a

}
func change(val *int) {
	*val = 55
}
