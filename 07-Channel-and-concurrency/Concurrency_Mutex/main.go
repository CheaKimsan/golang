package main

import (
	"fmt"
	"sync"
)

func main() {
	var counter int
	var lock sync.Mutex

	for range 1000 {
		go func() {
			lock.Lock()
			counter++
			lock.Unlock()
		}()
	}

	lock.Lock()
	fmt.Println("Expected Counter :", 1000)
	fmt.Println("Actual counter :", counter)
	lock.Unlock()
}
