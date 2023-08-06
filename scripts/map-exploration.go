package main

import "fmt"

func main() {
	m := make(map[string]string)
	fmt.Println(m)
	fmt.Printf("%q\n", m["hi"])
	m["hi"] = "bye"
	m["howdie"] = "Kapitan"
	fmt.Println(m)
}
