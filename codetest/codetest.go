package main

import (
	"fmt"
	"os"
)

func getDrives() (r []string) {
	fmt.Println("input : ")
	var get string
	fmt.Scan(&get)
	f, err := os.Open(string(get) + ":\\")
}
