package main

import (
	"fmt"
	"strings"
)

func main() {
	var b strings.Builder
	b.WriteString("Chadathan")
	b.WriteString(" ")
	b.WriteString("Rungrueang")
	fmt.Println(b.String())
}
