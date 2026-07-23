package main

import (
	"fmt"
	"sync"
	"time"
)

var (
	mutex   sync.Mutex
	counter int
)

func highPriority() {
	mutex.Lock()
	time.Sleep(100 * time.Microsecond)
	defer mutex.Unlock()
	counter += 10
}

func lowpriority() {
	mutex.Lock()
	defer mutex.Unlock()
	counter += 1
}

func main() {
	for i := 0; i < 1000; i++ {
		go highPriority()
		go lowpriority()
	}
	time.Sleep(5 * time.Second)
	fmt.Println("Final counter value", counter)
}
