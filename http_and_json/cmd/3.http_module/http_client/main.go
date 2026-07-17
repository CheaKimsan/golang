package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type StatusCode struct {
	Status   int    `json:"status"`
	Username string `json:"username,omitempty"`
	Pwd      string `json:"pwd"`
}

func main() {
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		response := StatusCode{
			Status:   http.StatusOK,
			Username: "Chea Kimsan",
			Pwd:      "123456",
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)

		if err := json.NewEncoder(w).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
		}
	})

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
