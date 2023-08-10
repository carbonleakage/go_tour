package main

import "fmt"

func main() {
	primes := [6]int{1, 2, 3, 5, 7, 11}

	var s []int = primes[1:4]
	fmt.Println("Slice of the prime array is ", s)
}
