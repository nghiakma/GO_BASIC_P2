//gom cac phan tu thuoc cung mot kieu dc luu tru lien ke nhau trong bo nho

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
}
