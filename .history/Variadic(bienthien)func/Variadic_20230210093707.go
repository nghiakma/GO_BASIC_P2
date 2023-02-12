package main

import "fmt"

//ham chi chap nhat mot so doi so co dinh, ham bien thien chap so doi so thay doi
func main() {

}

func find(numb int, numbs ...int) {
	fmt.Printf("type of numb is %T\n", numbs)
	found := false
	for i, v := range numbs {
		if v == numb {
			fmt.Println(numb, "found at indenx", i, "in", numbs)
		}
	}

}
