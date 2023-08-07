package main

import "fmt"

func main() {
	var b byte = 100
	fmt.Printf("%T, %v\n", b, b)
	fmt.Printf("%T, %v\n", uint8(65), uint8(65))
	if b >= uint8(65) {
		fmt.Println("Value is greater than 65")
	}
}
