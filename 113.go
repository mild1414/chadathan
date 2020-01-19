package main

import (
	"fmt"
)

func say(text string) {
	for i := 0; i < 3; i++ {
		fmt.Println(i, ":", text)
	}
}
