package main

import (
	"fmt"
	"io/ioutil"
)

func main() {
	bs, err := ioutil.ReadFile("test.txt")
	if err != nill {
		return
	}
	str := string(bs)
	fmt.Println(str)
}
