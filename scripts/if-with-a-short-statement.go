package main

import (
	"fmt"
	"math"
)

func findSmallestPOW(x, y, z float64) float64 {
	if pow := math.Pow(x, y); pow < z {
		return pow
	}
	return z
}

func main() {
	fmt.Println("The smallest value of 2 to the power 3 and 10 is ", findSmallestPOW(2, 3, 10))
	fmt.Println("The smallest value of 5 to the power 10 and 1e6 is ", findSmallestPOW(5, 10, 1e6))
}
