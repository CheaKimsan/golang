package main

import "fmt"

func double(x int) int {
	return x * 2
}
func tripple(x int) int {
	return x * 3
}
func apply(f func(int) int, x int) int {
	return f(x)
}

func main() {
	fiveTime := func(x int) int {
		return 5 * x
	}

	result := apply(fiveTime, 5)
	fmt.Println(result)
}
