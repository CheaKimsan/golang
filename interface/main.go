package main

import "fmt"

type Sound interface {
	Sound() string
}

type Animal struct {
	cat string
	dog string
}

func (a Animal) Sound() string {
	return a.dog
}

func main() {
	a := Animal{
		cat: "Meow",
		dog: "Woof",
	}

	var s Sound = a

	fmt.Println(s.Sound())
}
package 