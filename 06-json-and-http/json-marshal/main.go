package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type Student struct {
	Name     string   `json:"full_name"`
	Age      int      `json:"age"`
	Subjects []string `json:"subjects"`
}

type User struct {
	ID          int      `json: "user_id"`
	Name        string   `json: "name,omitempty"`
	Age         int      `json:"age"`
	Pwd         string   `json : "-"`
	Permissions []string `json : "roles"`
}

type Person struct {
	Name   string `json: "username"`
	Age    int    `json: "age"`
	Active bool   `json: "is_active"`
}

func main() {

	// Mashal Json
	u := User{
		ID:          001,
		Name:        "User A",
		Age:         21,
		Pwd:         "admin123",
		Permissions: []string{"admin", "group-member"},
	}

	b, err := json.Marshal(u)
	if err != nil {
		fmt.Println("error mashalling JSON : ", err)
		panic(err)
	}
	fmt.Println(string(b))

	jsonData := `{
		"full_name": "Chea Kimsan",
		"age": 22,
		"subjects": ["MIS", "Computer Science"]
	}`

	// Unmashal json
	var stu Student

	error := json.Unmarshal([]byte(jsonData), &stu)
	if error != nil {
		panic(err)
	}

	fmt.Println("Name:", stu.Name)
	fmt.Println("Age:", stu.Age)
	fmt.Println("Subjects:", stu.Subjects)

	// Encoded Json
	p := Person{
		Name:   "Chea Kimsan",
		Age:    21,
		Active: true,
	}
	f, err := os.Create("output.json")
	if err != nil {
		fmt.Println("Error creating file : ", err)
		panic(err)
	}

	defer f.Close()

	encoder := json.NewEncoder(f)

	err = encoder.Encode(&p)
	if err != nil {
		fmt.Println("error encoding person : ", err)
		panic(err)
	}

	fmt.Println(p)

}
