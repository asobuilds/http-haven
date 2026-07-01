package main

import(
	"fmt"
	"net/http"
)
func agentHandler(w http.ResponseWriter, r *http.Request) {
	name := r.Header.Get("User-Agent")
	if name == "" {
		name = "Unknown"
	}
	fmt.Fprintf(w, "You are visiting us using: [Aso]")
}
func main() {
	http.HandleFunc("/agent", agentHandler)
	fmt.Println("server is running...")
	http.ListenAndServe(":8080", nil)
}