package main

import (
	"fmt"
	"time"
)

func sentEmail(message string) {
	func() {
		time.Sleep(time.Millisecond * 250)
		fmt.Printf("Email Recieved: '%s'\n", message)
	}()
	fmt.Printf("Email Sent: '%s'\n", message)
}

func count(thing string) {
	for i := 1; i <= 5; i++ {
		fmt.Println(i, thing)
		time.Sleep(time.Millisecond * 500)
	}
}
func main() {
	count("Moew")
}
