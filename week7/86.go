package main

import "fmt"

func say() {
	fmt.Prinln("Hi Goku")
}

func main() {
	defer say()
	fmt.Println("Hello World")
}
