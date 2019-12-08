package main

import "fmt"

func main() {
	alphabets := [4]string{"A", "B", "C", "D"}
	fmt.Println(alphabets)

	x := alphabets[0:2]
	fmt.Println(x)

	y := x[2:4]
}
