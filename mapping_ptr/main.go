package main

import "fmt"

func main() {
	myMap := map[int]string{
		0: "GDI",
		1: "Center",
	}

	fmt.Println(myMap)

	somethingMap(myMap)

	fmt.Println(myMap)

}
func somethingMap(m map[int]string) {
	delete(m, 0)
}
