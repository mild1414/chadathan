package main

import (
	"fmt"
	"strings"
)

func main() {
	fmt.Println(strings.HasSuffix("Hello World", "world"))
	fmt.Println(strings.HasSuffix("Hello World", "World"))
}
