package main

func main() {
	elements := (make[string]string)
	elements["H"] = "Hydrogen"
	elements["He"] = "Helium"
	elements["Li"] = "Lithium"

	h := elements["H"]
	fmt.Println(h)

	n, found := elements["N"]
	
}
