package main

func compute(fn func(int, int) int) int {
	return fn(3, 4)
}
