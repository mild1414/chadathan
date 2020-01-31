package main

import "sync"

func generateInt(min, max int, ch chan int, wg *sync.WaitGroup) {
	defer wg.Done()
	for i := min; i <= max; i++ {
		ch <- i
	}
}
func main() {
	ch := make(chan int)
	var wg sync.WaitGroup
	wg.Add(2)
}
