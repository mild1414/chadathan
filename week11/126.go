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
	go generateInt(10, 20, ch, &wg)
	go generateInt(50, 200, ch, &wg)
	go func(){
		wg.Wait()
		close(ch)
	}
}
