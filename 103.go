package main

import "fmt"

type I interface {
	F()
}

func desc(i I) {
	fmt.Println("%v, %T \n", i, i)
}

type T1 struct {
	text string
}

func (t T1) F() {
	fmt.Println(t.text)
}

type T2 struct {
	number int
}

func (t T2) F() {
	fmt.Println(t.number)
}
