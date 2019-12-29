package main

func makeEven() func() int {
	even := 0
	even = even + 2
	return even
}
