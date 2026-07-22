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

func count(thing string, c chan string) {
	for i := 1; i <= 5; i++ {
		// fmt.Println(i, timething)
		c <- thing
		time.Sleep(time.Microsecond * 500)
	}
	close(c)
}
func main() {

	// Concurrency lets your program do multiple things without waiting for each one:

	// var wg sync.WaitGroup
	// wg.Add(1)
	// go func() {
	// 	count("Moew")
	// 	wg.Done()
	// }()
	// wg.Wait()

	// c := make(chan string)
	// go count("sheesh", c)
	// for msg := range c {
	// 	fmt.Println(msg)
	// }
	// fmt.Println("done, channel closed")

	// buffered channel
	// c := make(chan string, 2)
	// c <- "GDI"
	// c <- "SETEC"
	// c <- "CHOU"
	// msg = <-c
	// fmt.Println(msg)

	c1 := make(chan string)
	c2 := make(chan string)

	go func() {
		for {
			c1 <- "Every 500ms"
			time.Sleep(time.Millisecond * 500)
		}
	}()

	go func() {
		for {
			c2 <- "Every two second"
			time.Sleep(time.Second * 5)
		}
	}()

	for {
		select {
		case msg1 := <-c1:
			fmt.Println(msg1)
		case msg2 := <-c2:
			fmt.Println(msg2)
		}
	}

}
