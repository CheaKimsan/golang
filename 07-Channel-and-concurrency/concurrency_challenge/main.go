package main

import (
	"fmt"
	"time"
)

func main() {
	var counter int

	for range 1000 {
		go func() {
			counter++
		}()
	}

	time.Sleep((5 * time.Second))

	fmt.Println("Expected Counter :", 1000)
	fmt.Println("Actual counter :", counter)
}
