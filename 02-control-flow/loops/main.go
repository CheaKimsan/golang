package main

import "fmt"

func main() {
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(i)
	// }

	// a := []int{1, 2, 3, 4, 5, 6, 7}

	// for index, value := range a {
	// 	fmt.Printf("Index : %d value : %d\n", index, value)
	// }

	ages := map[string]int{
		"Chea Kimsan": 21,
		"Meng Hong":   23,
	}

	for name, age := range ages {
		fmt.Printf("%s is %d years olds. \n", name, age)

	}
}
