package main

func compute(fn func(int, int) int) int {
	return fn(3, 4)
}

func main() {
	sum := func(x, y int) int {
		return x + y
	}
}
