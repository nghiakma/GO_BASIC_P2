// gom cac phan tu thuoc cung mot kieu dc luu tru lien ke nhau trong bo nho
// mang trong go la mang gia tri chu khong phai mang tham chieu
package main

import "fmt"

func main() {
	var a [3]int
	a[0] = 12
	a[1] = 78
	a[2] = 50
	fmt.Println(a)

	b := [3]int{12, 78, 50}
	fmt.Println(b)

	c := [...]int{12, 78, 50}
	fmt.Println(c)

	d := [...]string{"USA", "China", "India", "Germany"}
	f := d
	f[0] = "Singapore"
	fmt.Println(d)
	fmt.Println(f)

	num := [...]int{5, 6, 7, 8, 8}
	fmt.Println("before passing to function", num)
	changeLocal(num)
	fmt.Println("after passing to function ", num)
	//chieu dai cua mang
	fmt.Println(len(num))

	//lap mang
	for i := 0; i < len(num); i++ {
		fmt.Printf("%d th element of a is %d\n", i, num[i])
	}

	//tinh tong bang for...range
	sum := 0
	for i, v := range num {
		fmt.Printf("%d th element of a is %d\n", i, v)
		sum += v
	}
	fmt.Println("\nsum of all elements of a", sum)
}

func changeLocal(num [5]int) {
	num[0] = 55
	fmt.Println("Inside function", num)
}
