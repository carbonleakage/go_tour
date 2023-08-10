package main

import (
	"fmt"
	"math"
)

func sqrtStop10Iterations(x float64) float64 {
	s := 1.0
	for i := 0; i < 10; i++ {
		s -= (s*s - x) / (2 * s)
		fmt.Println("Sqrt approximation is ", s)
	}
	return s
}

func sqrt(x float64, s float64) (float64, int) {
	tol := 1e-6
	i := 0
	for diff := math.Abs(x - s*s); diff > tol; i++ {
		s -= (s*s - x) / (2 * s)
		diff = math.Abs(x - s*s)
	}
	return s, i
}

func main() {
	fmt.Println("Square root of 100 is ", sqrtStop10Iterations(100))
	s, count := sqrt(100, 1.0)
	fmt.Printf("Square root of 100 (with starting guess 1.0) is %g, and it took %d iterations\n", s, count)
	s, count = sqrt(100, 50.0)
	fmt.Printf("Square root of 100 (with starting guess 5.0) is %g, and it took %d iterations\n", s, count)
}
