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

	//Appending to a slice : them vao mot lat cat
	cars := []string{"Ferrari", "Honda", "Ford"}
	fmt.Println("cars:", cars, "has old length", len(cars), "and capacity", cap(cars)) //capacity of cars is 3
	//add phan tu vao no thi capacity(suc chua) + 2
	cars = append(cars, "Toyota")
	fmt.Println("cars:", cars, "has new length", len(cars), "and capacity", cap(cars)) //capacity of cars is doubled to 6

	slice := make([]int, 3, 5)

	fmt.Println(len(slice))

	//zero values of slice type is nil
	var names []string
	if names == nil {
		fmt.Println("slice is nil going to append")
		names = append(names, "John", "Sebastian", "Vinay")
		fmt.Println("names contents:", names)
	}
	//chen mang vao mang
	veggies := []string{"potatoes", "tomatoes", "brinjal"}
	fruits := []string{"oranges", "apples"}
	food := append(veggies, fruits...)
	fmt.Println("food:", food)
	countries := []string{"USA", "Singapore", "Germany", "India", "Australia"}
	neededCountries := countries[:len(countries)-2]
}
