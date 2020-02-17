package main

import "fmt"

func main() {
	a := []string{"A", "B", "C", "D", "E"}
	deleteIndex := 3
	a = append(a[:deleteIndex], a[deleteIndex+1:]...)
	fmt.Println(a)
}
