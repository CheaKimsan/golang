package main

import (
	"fmt"
	"time"
)

func hi() {
	fmt.Println("Hello world goroutine")
}
func main() {
	go hi()
	// go routine works when has the schedule
	time.Sleep(time.Millisecond * 500)

}
