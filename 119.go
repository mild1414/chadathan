package main

import (
	"sync"
	"time"
)

func increment(data *int, mutex *sync.Mutex, wg *sync.WaitGroup) {
	start := time.Now()
	defer wg.Done()
	//defer mutex.Unlock()
	//mutex.Lock()
	*data++
}
