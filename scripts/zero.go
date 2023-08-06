package main

import "fmt"

func main() {
	var i int
	var f float64
	var b bool
	var s string
	// For string using %v returns the string literal.
	// From docs: %q - a single-quoted character literal safely escaped with Go syntax.
	fmt.Printf("%v, %v, %v, %q", i, f, b, s)

}
