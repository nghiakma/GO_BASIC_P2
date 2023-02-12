package main

import "fmt"

//slice : cat mang or chuoi
func main() {
	a := [...]int{76, 77, 78, 79, 80}
	var b []int = a[1:4]
	fmt.Println(b)
}
