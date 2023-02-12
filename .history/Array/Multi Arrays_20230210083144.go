package main

import "fmt"

//go build: no se tap mot tep thuc thi duoi dang ten du an cua ban

func printery(a [3][2]string) {
	for i := 0; i < len(a); i++ {
		for j := 0; j < len(a[i]); j++ {
			fmt.Printf("%s ", a[i][j])
		}
		fmt.Printf("\n")
	}
}
