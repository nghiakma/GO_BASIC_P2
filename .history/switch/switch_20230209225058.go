package main

import (
	"fmt"
)

// mot cau lenh co dieu kien de danh gia mot bieu thuc va so sanh no voi cac
// danh sach cac ket qua phu hop
func main() {
	finger := 0
	fmt.Printf("Fingerprint: %d\n", finger)
	switch finger {
	case 1:
		fmt.Println("Thumb")
	case 2:
		fmt.Println("Index")
	case 3:
		fmt.Println("Middle")
	case 4:
		fmt.Println("Ring")
	case 5:
		fmt.Println("Pinky")
	default:
		fmt.Println("incorrect finger number")
	}

	letter := "i"
	fmt.Printf("Letter %s is a ", letter)
	switch letter {
	case "a", "e", "i", "o", "u":
		fmt.Println("vowel")
	default:
		fmt.Println("not a vowel")
	}

	//fallthrough: dieu khien xuat phat tu cau lenh switch ngay sau khi
	//mot truong hop dc thuc thi
	//in ra 75 nho hon 100 va 200

}
func number() int {
	num := 15 * 5
	return num
}
