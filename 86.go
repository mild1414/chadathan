package main

func say() {
	fmt.Prinln("Hi Goku")
}

func main() {
	defer say()
}
