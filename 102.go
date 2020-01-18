package main

type I interface {
	F()
}

type T struct {
	text string
}

func (t T) F() {}
