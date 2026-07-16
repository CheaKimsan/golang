package main

import "fmt"

type student struct {
	firstName string
	lastName  string
	class     []string
	age       int
	height    int
	weight    int
}

func main() {
	s := student{
		firstName: "Kim",
		lastName:  "San",
		age:       22,
	}

	prettyPrintStudents(&s)
}

func prettyPrintStudents(s *student) {
	fmt.Printf("First Name : %s\nLast Name : %s\nAge : %d\n", s.firstName, s.lastName, s.age)
}
