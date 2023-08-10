package main

import "fmt"

type Vertex struct {
	X int
	Y int
}

var m map[string]Vertex

func main() {
	fmt.Println("Unintialised map is ", m)

	m = make(map[string]Vertex)
	fmt.Println("Unintialised map is ", m)
	m["Origin"] = Vertex{0, 0}
	fmt.Println("Initialsed with origin, the map is ", m)
}
