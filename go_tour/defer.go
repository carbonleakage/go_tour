package main

import "fmt"

func main() {
	defer fmt.Print("Kapitan\n")
	fmt.Print("Howdie, this func runs after calling ")
}
