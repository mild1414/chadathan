package main

import "fmt"

type I interface {
	F()
}

func desc(i I) {
	fmt.Println("%v, %T \n", i, i)
}

type T1 struct{}
