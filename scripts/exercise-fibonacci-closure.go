package main

import "fmt"

// 0 1 1 2 3 5 8 13
func fibonacci() func() int {
	current := 0
	next := 1
	var return_val int
	return func() int {
		switch current {
		case 0:
			{
				current = 1
				return 0
			}
		default:
			{
				return_val = current
				current, next = next, current+next
				return return_val
			}
		}

	}
}

func main() {
	f := fibonacci()
	for i := 0; i < 10; i++ {
		fmt.Println(f())
	}
}
