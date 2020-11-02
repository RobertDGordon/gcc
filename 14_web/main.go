package main

import (
	"fmt"
	"net/http"
)

func index(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>Hello World</h1>")
}

func about(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "<h1>About</h1>")
}

func main() {
	port := ":3000"
	http.HandleFunc("/", index)
	http.HandleFunc("/about", about)
	fmt.Println("Server online, port" + port)
	http.ListenAndServe(port, nil)
}
