package main

import (
	"fmt"
)

func main() {

	type student struct {
		firstName string
		lastName  string
		age       int
		subject   []string
	}

	// var stu student

	// stu = student{"Kim", "San", 21, []string{"Java", "Python"}}
	// fmt.Printf("%+v", stu)

	// stu1 := student{
	// 	firstName: "Meng",
	// 	lastName:  "Hong",
	// 	age:       20,
	// 	subject:   []string{"java", "mobile"},
	// }
	// fmt.Printf("%+v", stu1.firstName)

	// anonymous struct
	guardian := struct {
		firstName string
		lastName  string
	}{
		firstName: "Cristino",
		lastName:  "Ronaldo",
	}

	fmt.Println(guardian.firstName, guardian.lastName)
}
