package main

import (
	"fmt"
	"net/http"
)

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Hello world2")

}

func main() {
	http.HandleFunc("/hello", handler)
	http.ListenAndServe(":8080", nil)

}
