package main

func main() {
	data := 10
	go func() {
		data = 20
	}()
	go func(){
		fmt.Println(data)
	}
}
