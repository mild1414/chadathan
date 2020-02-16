package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["one"] = 12
	m["two"] = 05

	fmt.Println("map:", m)

	value := m["one"]
	fmt.Println("value")
}
