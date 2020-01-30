package main

import "time"

func println(tick, boom <-chan time.Time) {
	for {
		select {
		case <-tick:
		}
	}
}
