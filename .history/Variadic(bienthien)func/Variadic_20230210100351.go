package main

import "fmt"

//ham chi chap nhat mot so doi so co dinh, ham bien thien chap so doi so thay doi
func main() {
	find(89, 89, 90, 95)
	find(78, 38, 56, 98)

	//Gotcha //sua doi(modifying) slice inside a variadic function
	welcome := []string{"hello", "world"}
	change(welcome...)
	fmt.Println(welcome)
}
func change(s ...string) {
	s[0] = "Go"
}
func find(numb int, numbs ...int) {
	fmt.Printf("type of numb is %T\n", numbs)
	found := false
	for i, v := range numbs {
		if v == numb {
			fmt.Println(numb, "found at index", i, "in", numbs)
			found = true
		}
	}
	if !found {
		fmt.Println(numb, "not found in ", numbs)
	}
	fmt.Printf("\n")

}
