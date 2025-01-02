package main

import (
	"fmt"
	"net/http"
)

func main() {
	http.HandleFunc("/", simple)
	http.HandleFunc("/api", api)
	http.ListenAndServe(":8080", nil)
}

func simple(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "hello")
}

func api(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "api")
}
