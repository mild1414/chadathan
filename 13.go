package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.Cont("Hello world", "o"))
	fmt.Println(strings.Count("Hello World", ""))
}
