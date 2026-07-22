package main

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"
	"time"
)

type UserRes struct {
	Status   int    `json:"status"`
	Username string `json:"username,omitempty"`
	Pwd      string `json:"pwd"`
}

func main() {
	http.HandleFunc("/user", func(w http.ResponseWriter, r *http.Request) {
		response := UserRes{
			Status:   http.StatusOK,
			Username: "Chea Kimsan",
			Pwd:      "123456",
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		json.NewEncoder(w).Encode(response)
	})

	http.HandleFunc("/delay", func(w http.ResponseWriter, r *http.Request) {
		time.Sleep(5 * time.Second)
		response := UserRes{
			Status:   http.StatusOK,
			Username: "Chea Kimsan",
			Pwd:      "123456",
		}

		var buff bytes.Buffer
		if err := json.NewEncoder(&buff).Encode(response); err != nil {
			http.Error(w, err.Error(), http.StatusInternalServerError)
			return
		}

		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusOK)
		w.Write(buff.Bytes())

	})

	fmt.Println("Server running on :8080")
	http.ListenAndServe(":8080", nil)
}
