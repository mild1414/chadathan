package main

import (
	"fmt"
	"time"
)

func println(tick, boom <-chan time.Time) {
	for {
		select {
		case <-tick:
			fmt.Println("tick.")
		case <-boom:
			fmt.Println("BOOM!")
			return
		default:
			fmt.Println("	.")
			time.Sleep(50 * time.Millisecond)
		}
	}
}
