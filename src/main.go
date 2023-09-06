package main

import (
	"net/http"
	"fmt"
	// "io" // io.WriteString similar to fmt.Fprintf
)

const (
	PORT = "8080"
)

func handler(w http.ResponseWriter, r *http.Request) {
	switch r.Method {
		case http.MethodPost:
			fmt.Fprintf(w, "You made a POST request.")
			fmt.Printf("* Client made a POST request\n")
		default:
			http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
	}
}

func main() {
	http.HandleFunc("/generate", handler)
	fmt.Printf("[HTTP Server started on port %s]\n", PORT)
	http.ListenAndServe(":"+PORT, nil)
}