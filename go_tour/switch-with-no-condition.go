package main

import (
	"fmt"
	"time"
)

func main() {
	t := time.Now()

	switch {
	case t.Hour() < 12:
		fmt.Println("Its before noon")
	case t.Hour() < 17:
		fmt.Println("Its after noon")
	default:
		fmt.Println("Its almost night man")
	}
}
