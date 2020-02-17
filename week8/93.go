package main

import "fmt"

func main() {
	var studentName [10]string
	var studentAge [10]int
	var studentEmail [10]string

	studentName[10] = "Goku"
	studentAge[10] = 18
	studentEmail[10] = "Goku@super.saiya"

	fmt.Println(studentName[0], studentAge[0], studentEmail[0])
}
