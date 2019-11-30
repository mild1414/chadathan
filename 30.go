package main

import "fmt"

func main() {
	fmt.Printf("My name is %s, i am %d year old\n", "Goku", 18)

	fmt.Printf("type = %T \n", 3.14159265359)
	fmt.Printf("pi = %f \n", 3.14159265359)
	fmt.Sprintf("pi = %2.2f \n", 3.14159265359)
	fmt.Printf("pi = %9.f \n", 3.14159265359)
	fmt.Printf("pi = %-9.f \n", 3.14159265359)
	fmt.Println("pi = %09.f \n", 3.14159265359)
	fmt.Println("pi = %09.2f \n", 3.14159265359)
	fmt.Printf("pi = %E \n", 3.14159265359)
}
