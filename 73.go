package main

import "fmt"

func main() {
	var a int
ReadInput:
	fmt.Println("type number :")
	fmt.scan(&a)
	if a < 20 {
		goto ReadInput
	}
	fmt.Println(a)
}
