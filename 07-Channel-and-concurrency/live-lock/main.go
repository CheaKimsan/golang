package main

import (
	"fmt"
	"sync"
)

var (
	mutexA sync.Mutex
	mutexB sync.Mutex
)

func routineA(name string) {
	mutexA.Lock()
	defer mutexA.Unlock()
	fmt.Println(name, "has mutex A")
	fmt.Println(name, "try to acquired mutex B...")
	if !mutexB.TryLock() {
		fmt.Println(name, "can not acquired mutex B. Backing off.")
		for i := 0; i < 1000; i++ {

		}
		routineA(name)
	} else {
		defer mutexB.Unlock()
		fmt.Println(name, "acquired mutex B")
		fmt.Println(name, "doing some work...")
	}
}

func routineB(name string) {
	mutexB.Lock()
	defer mutexB.Unlock()
	fmt.Println(name, "has mutex B")
	fmt.Println(name, "try to acquired mutex A...")
	if !mutexA.TryLock() {
		fmt.Println(name, "can not acquired mutex A. Backing off.")
		for i := 0; i < 1000; i++ {

		}
		routineB(name)
	} else {
		defer mutexA.Unlock()
		fmt.Println(name, "acquired mutex A")
		fmt.Println(name, "doing some work...")
	}
}

func main() {
	go routineA("Routine A")
	go routineB("Routine B")

	fmt.Println("wating ...")
	select {}
}
