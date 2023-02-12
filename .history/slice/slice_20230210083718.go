package main

import "fmt"

//slice : cat mang or chuoi
func main() {
	a := [...]int{76, 77, 78, 79, 80}
	//a[1:4]= a[start:end-1]
	var b []int = a[1:4]
	fmt.Println(b)
	//sua doi slice
	darr := [...]int{57, 89, 90, 82, 100, 78, 67, 69, 59}
	dslice := darr[2:5]

	fmt.Println("array before", darr)

	for i := range dslice {
		dslice[i]++
	}
	fmt.Println("array after", darr)
}
