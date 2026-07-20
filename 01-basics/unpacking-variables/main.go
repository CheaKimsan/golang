package main

import "fmt"

func main() {
	// Group Similar Variable
	var (
		myName     string = "chea kimsan"
		university string = "Setec"
	)

	fmt.Println(myName)
	fmt.Println(university)

	// multiple variable declaration
	age := 25
	year := 2026

	fmt.Println(age, year)

	age, gdi := 2027, "GDI Center"
	fmt.Println(age, gdi)

	gdi = "GDI Center"
	fmt.Println(gdi)
}
