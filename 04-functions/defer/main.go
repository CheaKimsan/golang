package main

import "fmt"

func main() {
	message := "Hello, "
	greetingfn := func(name string) {
		fmt.Println(message + name)
	}

	defer greetingfn("Alex")
	defer greetingfn("Bob")

	message = "Hi, "
}
