package main

import "fmt"

func main() {
	i, j := 1, 2

	p := &i
	fmt.Println("i has value, read by the pointer", *p)
	*p = 20
	fmt.Println("i has value, read by the pointer", *p)
	fmt.Println("i has value, read by the variable", i)
	p = &j
	fmt.Println("i has value, read by the pointer", *p)
	*p = 200
	fmt.Println("i has value, read by the pointer", *p)
	fmt.Println("i has value, read by the variable", j)

}
