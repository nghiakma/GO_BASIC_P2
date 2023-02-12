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
	//slice chua tat ca gia tri trong mang
	dsrr := darr[:]
	fmt.Println(dsrr)

	//length slice and capacity of slice
	fruitarray := [...]string{"apple", "orange", "grape", "mango", "water melon", "pine apple", "chikoo"}
	fruitslice := fruitarray[1:3]
	fmt.Printf("length of slice %d capacity %d", len(fruitslice), cap(fruitslice))

	//creating a slice using make // parameter capacity is optional
	//and defaults to the length
	i := make([]int, 5, 5)
	i[0] = 1
	fmt.Println(i)
}
