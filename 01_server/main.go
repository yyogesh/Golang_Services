package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", handleRoot)
	fmt.Println("Server is running on port 8080...")
	http.ListenAndServe(":8080", nil)
}

func handleRoot(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to golang server test1")
}
