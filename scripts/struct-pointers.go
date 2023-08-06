package main

import "fmt"

type Vertex struct {
	X     int
	Y     int
	label string
}

func main() {
	point_x := Vertex{1, 2, "point x"}
	p := &point_x
	fmt.Println("Vertex with label is ", *p)
	p.X = 200
	fmt.Println("Vertex with X modified by pointer", point_x)
	fmt.Printf("Struct pointer has %T type and %v value", p, p)
}
