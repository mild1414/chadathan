package main

import "fmt"

func main() {
	fmt.Print("input : ")
	var water float32
	var fire float32
	var totel float32
	n, err := fmt.Scan(&water, &fire)
	water = water * 1
	fire = fire * 3
	fmt.Println("water : %f", water)
	fmt.Println("fire : %f", fire)

	water = water + ((water * 7) / 100)
	fire = fire + ((fire * 7) / 100)
	totel = water + fire
	fmt.Printf("ภาษี %f", water)

}
