package main

import "fmt"

type student struct {
	name  string
	age   int
	email string
}

func (std student) introduce() {
	fmt.Println("Hello my name is", std.name)
}

type pupil struct {
	address string
	student
}

func main() {}
