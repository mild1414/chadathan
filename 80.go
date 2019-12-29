package main

import "fmt"

func say(greet string) func(string) string {
	return func(name string) string {
		return greet + name
	}
}

func main() {
	x := say("Hello")

	fmt.Println(x("Goku"))
}
