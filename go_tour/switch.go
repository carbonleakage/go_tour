package main

import (
	"fmt"
	"runtime"
)

func main() {
	fmt.Print("This script is running on ")
	switch os := runtime.GOOS; os {
	case "darwin":
		fmt.Println("Mac OS X")
	case "linux":
		fmt.Println("Linux")
	default:
		fmt.Printf("%s\n", os)

	}
}
