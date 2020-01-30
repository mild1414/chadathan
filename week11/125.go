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
		}
	}
}
