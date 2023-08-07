package main

import (
	"fmt"
	"math"
)

type ErrNegativeSqrt float64

func (e ErrNegativeSqrt) Error() string {
	return fmt.Sprintf("cannot Sqrt negative number: %f", float64(e))
}

func sqrtStop10Iterations(x float64) float64 {
	s := 1.0
	for i := 0; i < 10; i++ {
		s -= (s*s - x) / (2 * s)
		fmt.Println("Sqrt approximation is ", s)
	}
	return s
}

func sqrt(x float64, s float64) (float64, int, error) {
	if x < 0 {
		return 0, 0, ErrNegativeSqrt(x)
	}

	tol := 1e-6
	i := 0
	for diff := math.Abs(x - s*s); diff > tol; i++ {
		s -= (s*s - x) / (2 * s)
		diff = math.Abs(x - s*s)
	}
	return s, i, nil
}

func presentResults(x float64, guess float64) {
	s, count, err := sqrt(x, guess)
	if err == nil {
		fmt.Printf("Square root of %f (with starting guess %f) is %f, and it took %d iterations\n", x, guess, s, count)
	} else {
		fmt.Println(err)
	}
}

func main() {
	fmt.Println("Square root of 100 is ", sqrtStop10Iterations(100))
	presentResults(10, 1.0)
	presentResults(100, 50.0)
	presentResults(-100, 1.0)
}
