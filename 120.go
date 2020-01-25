package main

import (
	"fmt"
	"sync"
	"time"
)

func increment(data *int, rwmutex *sync.RWMutex, wg *sync.WaitGroup) {
	start := time.Now()
	defer wg.Done()
	defer rwmutex.Unlock()
	rwmutex.Lock()
	*data++
	fmt.Println(time.Since(start), "Increment to:", *data)
}
func read(data *int, rwmutex *sync.RWMutex, wg *sync.WaitGroup) {
	start := time.Now()
	defer wg.Done()
	defer rwmutex.RUnlock()
	rwmutex.RLock()
	fmt.Println(time.Since(start), "Data =", *data)
}
func main() {
	var rwmutex sync.RWMutex
	var wg sync.WaitGroup
	data := 10
	wg.Add(10)
	for i := 0
}
