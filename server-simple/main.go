package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		if r.Method == "POST" {
			fmt.Println("Get POST request")
			w.Write([]byte("hello world dari post"))
			return
		}
		fmt.Println("Get non-POST request")
		// w.Write([]byte("hello world"))
		jsonB, _ := json.Marshal(map[string]string{
			"name":  "john snow",
			"house": "stark",
		})
		w.Write(jsonB)
	})
	fmt.Println("Server running on localhost:8080")
	http.ListenAndServe(":8080", nil)
}
