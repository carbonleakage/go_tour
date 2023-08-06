package main

import "fmt"

func doublingSum(x int) (sum int, i int) {
	sum = 1
	i = 0
	for sum < x {
		sum += sum
		i += 1
	}
	return
}

func main() {
	finalSum, count := doublingSum(10)
	fmt.Printf("Starting with 1 and adding the value to itself, after %d times, we get %d\n", count, finalSum)
	finalSum, count = doublingSum(1e6)
	fmt.Printf("Starting with 1 and adding the value to itself, after %d times, we get %d\n", count, finalSum)
}
