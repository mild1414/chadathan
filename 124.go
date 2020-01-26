package main

func sum(c chan int, number ...int) {
	sum := 0
	for _, v := range number {
		sum = sum + v
	}
	c <- sum
}
