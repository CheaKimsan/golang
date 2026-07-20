package main

import "fmt"

func main() {
	// var a []int
	// fmt.Println(a == nil)

	// a := []int{1, 2, 3, 4, 5}

	// a = []int{5, 2: 10, 50}
	// fmt.Println(len(a))

	// fmt.Println(cap(a))

	// append

	// a := []int{10, 20, 30}
	// b := []int{40, 50, 60}

	// fmt.Println(append(a, b...))

	// a := []int{1, 2, 3, 4, 5}
	// slide from index 2 to index 5
	// b := a[1:5]
	// append is to push new number into array
	// fmt.Println(append(b, 10))

	a := []int{1, 2, 3, 4, 5}
	b := a[2:5]

	n := copy(a, b)

	fmt.Println("Copied:", n)
	fmt.Println("a:", a)
	fmt.Println("b:", b)

}
