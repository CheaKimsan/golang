package main

import "fmt"

func main() {
	nameAge := map[string]int{
		"a": 24,
		"b": 30,
		"c": 60,
	}

	nameAge["foo"] = 25
	nameAge["bar"] = 30
	nameAge["foo bar"] = 45

	fmt.Println(nameAge)
}
