package main

import (
	"fmt"
	"time"
)

func say(text string) {
	for i := 0; i < 3; i++ {
		fmt.Println(time.Now(), ":", i, ":", txt)
	}
}
