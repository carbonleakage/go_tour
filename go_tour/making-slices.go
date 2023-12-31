package main

import "fmt"

func printSlice(label string, s []int) {
	fmt.Printf("label=%s, len=%d, cap=%d, %v\n", label, len(s), cap(s), s)
}
func main() {
	a := make([]int, 5)
	printSlice("a", a)

	b := make([]int, 0, 5)
	printSlice("b", b)

	c := b[:2]
	printSlice("c", c)

	d := c[2:5]
	printSlice("d", d)
}
