package main

func sum(c chan int, number ...int) {
	sum := 0
	for _, v := range number {
		sum = sum + v
	}
	c <- sum
}

func printer(c1, c2 chan int) {}
