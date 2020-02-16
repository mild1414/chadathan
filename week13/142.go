package main

import "fmt"

func main() {
	m := make(map[string]int)
	m["one"] = 12
	m["two"] = 05

	fmt.Println("map:", m)

	value := m["one"]
	fmt.Println("value:", value)

	len := len(m)
	fmt.Println("len:", value)

	delete(m, "two")

	_, can_valeu := m["two"]
	fmt.Println("can_valeu")

	mapint := make(map[int]int)
	mapint[1] = 12
	mapint[2] = 05
	for i := range mapint {
		fmt.Println("i:", i)
	}
}
