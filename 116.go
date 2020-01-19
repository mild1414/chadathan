package main

import (
	"sync"
	"time"
)

func say(tet string, sleep time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
}
