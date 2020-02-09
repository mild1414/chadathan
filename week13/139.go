package main

type Employee struct {
	firstName, lastName string
	age, salary         int
}

func main() {
	emp1 := Employee{
		firstName: "Sam",
		age:       25,
		salary:    500,
	}
}
