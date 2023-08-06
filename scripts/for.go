package main

import "fmt"

func calculateSum(x int) (sum int) {
	sum = 0
	for i := 0; i < x; i++ {
		sum += i
	}
	return
}

func main() {
	fmt.Println("Sum of integers 0 to 10 is ", calculateSum((10)))
	fmt.Println("Sum of integers 0 to 10 is ", calculateSum((10e6)))
}
