package main

import "fmt"

type student struct {
	name  string
	age   int
	email string
}

func (std student) introduce() {
	fmt.Println("")
}
