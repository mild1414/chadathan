package main

import (
	"sync"
	"time"
)

func increment(data *int, rwmutex *sync.RWMutex, wg *sync.WaitGroup) {
	start := time.Now()
	defer wg.Done()
	defer rwmutex.Unlock()
}
