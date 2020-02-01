package main

import "os"

func main() {
	dir, err := os.Open(".")
	if err != nill {
		return
	}
	defer dir.Close()
}
