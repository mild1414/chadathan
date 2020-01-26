package main

import "fmt"

func send1(c <-chan string) {
	c <- "Hello"
}

func receive1(c <-chan string) {
	fmt.Println(<-c)
}

func send2(c chan<- string)
