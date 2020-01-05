package main

type student struct {
	name  string
	age   int
	email string
}

func main() {
	var a student
	a.name = "Goku"
	a.age = 18
	a.email = "Goku@super.saiya"

	b := student{"Gohan", 3, "Gohan@super.saiya"}

	c := student
}
