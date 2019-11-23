package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Replace("Hello world", "l", 2))
	fmt.Println(strings.Replace("Hello World", "l", -1))
}
