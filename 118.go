package main

import (
	"sync"
	"time"
)

func increment(data *int, mutex *sync.Mutex, wg *sync.WaitGroup) {
	start := time.Now()
	defer wg.Done()
	default mutex.Unlock()
	mutex.Lock()
	*data+++
	fmt.Println()
}
