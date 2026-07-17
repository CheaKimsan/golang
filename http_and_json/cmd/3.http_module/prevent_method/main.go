package main

import (
	"fmt"
	"net/http"
)

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method != http.MethodGet {
		http.Error(w, "Only Get Method is Allowed", http.StatusMethodNotAllowed)
	}
}
func main() {
	http.HandleFunc("/hello", helloHandler)

	fmt.Println("Try going to 8080 port")

	err := http.ListenAndServe(":8080", nil)

	fmt.Println(err)
}
