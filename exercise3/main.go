package main

import (
	"fmt"
	"io"
	"net/http"
)
func countHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodGet {
		fmt.Fprint(w, "Send a POST request with text to count words")
		return
	}
	if r.Method == http.MethodPost {
		body, _ := io.ReadAll(r.Body)
		fmt.Fprintf(w, "%d", len(body))
		return
	}
	http.Error(w, "Method Not Allowed", http.StatusMethodNotAllowed)
}

func main() {
	http.HandleFunc("/count", countHandler)
	fmt.Println("server running")
	http.ListenAndServe(":8080", nil)
}
