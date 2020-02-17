package main

import "fmt"

func subtract(number *int) {
	*number = *number - 1
}

func main() {
	x := 10
	subtract(&x)
	fmt.Println(x)
}
