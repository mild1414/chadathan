package main

type myError struct {
	error string
}

func (e myError) Error() string {
	return e.error
}
func say(word string) error {}
