package main

import "fmt"

func writeLine(a ...interface{}) {
	for _, v := range a {
		fmt.Println(v)
	}
}

func main() {
	writelLin(1, 3.14, "Hello", true)
}
