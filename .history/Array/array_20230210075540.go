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

}
