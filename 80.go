package main

func say(greet string) func(string) string {
	return func(name string) string {
		return greet + name
	}
}
