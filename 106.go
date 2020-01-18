package main

type I interface{}

func main() {
	var i I
	i = "Hello"
	s, ok := i.(string)
}
