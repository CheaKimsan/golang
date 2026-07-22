package main

import (
	"fmt"
	"sync"
)

func hi(wg *sync.WaitGroup) {
	defer wg.Done()
	fmt.Println("Hello world goroutine")
}

func main() {
	var wg sync.WaitGroup

	wg.Add(1)
	go hi(&wg)

	wg.Wait()
}
