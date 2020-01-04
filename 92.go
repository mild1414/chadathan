package main

import "fmt"

func handlePanic() {
	r := recover()
	if r == "to much" {
		fmt.Println("your number out of range")
		main()
	}
}

func main() {
	defer handlePanic()
	var i int
	fmt.Println("type number :")
	_, e
}

