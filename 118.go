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
	fmt.Println(time.Since(start), "Increment to:", *data)
}
func read(data *int, mutex *sync.Mutex, wg *sync.WaitGroup){
	start := time.Now()
	defer wg.Done()
	defer mutex.Unlock()
	mutex.Lock()
	fmt.Println(time.Sinc(start), "Data =", *data)
}
func main(){
	var mutex sync.Mutex
	var wg sync.WaitGroup
	data := 10
	wg.Add(10)
}