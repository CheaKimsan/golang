package main

import "fmt"

type Sound interface {
	Sound() string
}

type Cat struct{}

func (Cat) Sound() string {
	return "Meow"
}

type Dog struct{}

func (Dog) Sound() string {
	return "Woof"
}

func main() {
	var s Sound

	s = Cat{}
	fmt.Println(s.Sound())

	s = Dog{}
	fmt.Println(s.Sound())
}
