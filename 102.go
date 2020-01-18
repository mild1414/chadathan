package main

import "fmt"

type I interface {
	F()
}

type T struct {
	text string
}

func (t T) F() {
	fmt.Println(t.text)
}
