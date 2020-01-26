package main

func send1(c <-chan string) {
	c <- "Hello"
}
