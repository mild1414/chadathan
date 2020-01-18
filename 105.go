package main

import "fmt"

func hello(t interface{}) {
	fmt.Printf("Hello %T \n", t)
}
