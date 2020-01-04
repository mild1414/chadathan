package main

import "fmt"

func handlePanic() {
	r := recover()
	if r == "to much" {
		fmt.Println("your number out of range")
	}
}
