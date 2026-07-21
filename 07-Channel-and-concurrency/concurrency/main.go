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
func main() {
	sentEmail("Hello, Kimsan..!")
}
