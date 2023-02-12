package main

import "fmt"

func main() {
	name := "Hello world"
	fmt.Printf("String: %s\n", name)

}
func printChars(s string) {
	fmt.Printf("Characters: ")
	for i := 0; i < len(s); i++ {
		fmt.Printf("%c ", s[i])
	}
}
