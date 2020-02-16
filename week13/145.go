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
	fmt.Printf("water : %f\n", water)
	fmt.Printf("fire : %f\n", fire)

	water = water + ((water * 7) / 100)
	fire = fire + ((fire * 7) / 100)
	totel = water + fire
	fmt.Printf("รวมภาษี %f\n ", totel)
	fmt.Printf("water %f\n", water)
	fmt.Printf("fire %f\n", fire)
	fmt.Println(`number of argument `, n)
	fmt.Println(`error `, err)
}
