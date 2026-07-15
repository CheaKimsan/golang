package main

import "fmt"

func main() {
	_, l := fullname("Chea", "Kimsan")
	fmt.Println(l)
}

func fullname(firstName string, lastName string) (string, int) {
	fn := firstName + " " + lastName
	return fn, len(fn)
}
