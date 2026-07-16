package main

import "fmt"

func main() {
	var a *int
	var i interface{}

	fmt.Println(a == nil)
	fmt.Println(i == nil)
}
