package main

import "fmt"

func main() {
	name := "Hello world"
	fmt.Printf("String: %s\n", name)
	printChars(name)

}
func printChars(s string) {
	fmt.Printf("Characters: ")
	for i := 0; i < len(s); i++ {
		//%c la dinh dang in ra character
		fmt.Printf("%c ", s[i])
	}
}
