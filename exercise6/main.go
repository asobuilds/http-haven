// package main

// import(
// 	"fmt"
// 	"net/http"
// )
// func dashboardHandler(w http.ResponseWriter, r *http.Request) {
// 	key := r.Header.Get("X-API-Key")
// 	if key != "secret123" {
// 		http.Error(w, "invalid key", http.StatusUnauthorized)
// 		return
// 	}
// 	fmt.Fprintf(w, "Welcome to your dashboard")
		
// }
// func main() {
// 	http.HandleFunc("/dashboard", dashboardHandler)
// 	fmt.Println("server is 	running...")
// 	http.ListenAndServe(":8080", nil)
// }












package main

import(
	"fmt"
	"net/http"
)
func dashboardHandler(w http.ResponseWriter, r *http.Request) {
	key := r.Header.Get("X-API-Key")
	if key != "secret123" {
		http.Error(w, "invalid key", http.StatusUnauthorized)
		return
	}
	fmt.Fprint(w, "Welcome to your dashboard secret123")
}
func main() {
	http.HandleFunc("/dashboard", dashboardHandler)
	fmt.Println("server is running...")
	http.ListenAndServe(":8080", nil)
}