package main

import "fmt"

func main() {
	n, e := fmt.Println("Hello", "World", 123, 456, "\n")
	fmt.Print("number of bytes written :", n, "\n")
}
