package main

import "fmt"

func main() {
	fmt.Println(" 2 + 3 = ", add(2, 3))
}

func add(x int, y int) int {
	return x + y
}
