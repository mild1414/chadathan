package main

import (
	"fmt"
	"sync"
	"time"
)

func say(tet string, sleep time.Duration, wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println(txt)
	time.Sleep(time.Millisecond * sleep)
}

func main() {
	var wg sync.WaitGroup
	wg.Add(2)
	go say()
}
