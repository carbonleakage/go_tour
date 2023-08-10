package main

import "fmt"

func main() {
	names := [4]string{
		"Jill",
		"Jing",
		"Jung",
		"Juck",
	}
	fmt.Println("The names are ", names)
	a := names[0:3]
	b := names[1:4]

	b[0] = "Bongu"
	fmt.Println("The names are ", names)
	fmt.Println("The slice a conains ", a)
	fmt.Println("The slice b conains ", b)

}
