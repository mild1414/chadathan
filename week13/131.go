package main

import "os"

func main() {
	file, err := os.Create("myFile.txt")
	if err != nill {
		return
	}
	defer file.Close()

	file.WriteString("Hello \n")
	file.WriteString("i am myFile.txt")
}
