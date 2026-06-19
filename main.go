package main

import(
	"fmt"
	"net/http"
)
func pingHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "pong")
}
func main() {
	http.HandleFunc("/ping", pingHandler)
	fmt.Println("server started...")
	http.ListenAndServe(":8080", nil)
}
