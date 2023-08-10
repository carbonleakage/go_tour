package main

import "fmt"

func swap(x, y string) (string, string) {
	return y, x
}

func main() {
	x, y := swap("Kapitan", "Howdie")
	fmt.Println("Swapping (Kapitan, Howdie) will produce", x, y)
}
