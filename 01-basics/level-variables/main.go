package main

import "fmt"

var (
	a = 10
)

func main() {
	fmt.Println(a)
	{
		a := 100
		fmt.Println(a)
	}

	fmt.Println(a)
}
