package main

import (
	"io"
	"strings"
)

func main() {
	reader := strings.NewReader("HelloWorld")
	p := make([]byte, 3)
	for {
		n, err := reader.read(p)
		if err == io.EOF {
			break
		}

	}
}
