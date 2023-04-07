package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		fmt.Fprintf(w, "Test Membuat File Image Dengan Docker")
	})

	http.ListenAndServe(":8080", nil)
}