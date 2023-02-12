package main

import "fmt"

func binarySearch(arr []int, l int, r int, x int) int {
	if r >= l {
		mid := l + (r-l)/2

		//Neu arr[mid] = x tra ve va ket thuc
		if arr[mid] == x {
			return mid
		}

		if arr[mid] > x {
			return binarySearch(arr, l, mid-1, x)
		}

		if arr[mid] < x {
			return binarySearch(arr, mid+1, r, x)
		}

	}
	//Khong tim thay
	return -1
}

func main() {
	arr := []int{2, 3, 4, 10, 40}
	n := len(arr)
	x := 10
	result := binarySearch(arr, 0, n-1, x)
	if result == -1 {
		fmt.Println("Khong tim thay ", x)
	} else {
		fmt.Println("tim thay ", result)
	}
}
