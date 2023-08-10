package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

func main() {
	i := Vertex{10, 20}
	fmt.Println("My first go struct is ", i)
}
