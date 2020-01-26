package main

import "fmt"

func sum(c chan int, number ...int) {
	sum := 0
	for _, v := range number {
		sum = sum + v
	}
	c <- sum
}

func printer(c1, c2 chan int) {
	select {
	case num1 := <-c1:
		fmt.Println("channel-1 :", num1)
	case num2 := <-c2:
		fmt.Println("channal-2 :", num2)
	}
}

func main() {}
