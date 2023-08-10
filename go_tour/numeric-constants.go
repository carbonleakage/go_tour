package main

import "fmt"

const (
	Big   = 1 << 100
	Small = Big >> 99
)

func needInt(x int) int { return 10*x + 1 }
func needFloat(x float64) float64 {
	return 0.1 * x
}

func main() {
	fmt.Println(needInt(Small))
	// fmt.Println(needInt(Big))
	fmt.Println(needFloat(Small))
	fmt.Println(needFloat(Big))
}
