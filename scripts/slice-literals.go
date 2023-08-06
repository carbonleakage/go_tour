package main

import "fmt"

type FamilyMember struct {
	name string
	age  int
}

func main() {
	q := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println("The int slice is ", q)

	t := []bool{true, false, true, true, true, false, false, false}
	fmt.Println("The boolean slice is ", t)

	family := []FamilyMember{
		{"Peter", 40},
		{"Lois", 35},
		{"Meg", 35},
		{"Chris", 35},
		{"Stewie", 35},
	}
	fmt.Println("Family Guy family members are ", family)
}
