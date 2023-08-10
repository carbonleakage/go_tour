package main

import (
	"fmt"
	"math"
)

func createPowArray(x int, y int) (powArray []int) {
	for i := 0; i < x; i++ {
		powArray = append(powArray, int(math.Pow(float64(i), float64(y))))
	}
	return
}

func main() {
	xs := createPowArray(10, 2)

	for i, val := range xs {
		fmt.Printf("%d ** 2 =%d\n", i, val)
	}
}
