package main

import (
	"fmt"
	"math"
)

type Vertex struct {
	X, Y float64
}

func (v Vertex) Abs() float64 {
	return math.Sqrt(v.X*v.X + v.Y*v.Y)
}

func (v Vertex) Scale(f float64) {
	v.X *= f
	v.Y *= f
}

func main() {
	v := Vertex{3, 4}
	fmt.Printf("The magnitude of %v is %f\n", v, v.Abs())
	v.Scale(0.5)
	fmt.Printf("Vertex scaled %f gives %v\n", v, 0.5)
}
