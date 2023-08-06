package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("When is the next Saturday?")
	today := time.Now().Weekday()

	switch time.Saturday {
	case today:
		fmt.Println("Today is Saturday!")
	case today + 1:
		fmt.Println("Tomorrow!")
	case today + 2:
		fmt.Println("In two days!")
	default:
		fmt.Println("Too far away!")
	}

}
