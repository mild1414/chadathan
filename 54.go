package main

import "fmt"

func main() {
	a := []int{1, 2, 3}
	b := a
	fmt.Println(a, b)
	a[0] = 10
	fmt.Println(a, b)
	fmt.Println("________________")
	c := []int{1, 2, 3}
	d := make([]int, len(c))
}
