package main

import "fmt"

type I interface{}

func main() {
	var i I
	i = "Hello"
	s, ok := i.(string)
	fmt.Println(s, ok)

	f, ok := i.(float64)
	fmt.Println(f, ok)

	b := i.(bool)
	fmt.Println(b)
}
