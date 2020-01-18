package main

import "fmt"

type I interface{}

func desc(i I) {
	fmt.Printf("%v ,%T \n", i, i)
}

func main() {}
