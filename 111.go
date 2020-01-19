package main

func Average(numbers ...float64) float64 {
	var total float64
	for _, v := range numbers {
		total = total + v
	}
}
