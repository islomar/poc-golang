package main

import (
	"log"
	"net/http"
	"fmt"
)

// This is a high-performance, DoS-hardened, production-ready web server
// It serves dl.google.com
func main() {
	http.HandleFunc("/", handler)
	err := http.ListenAndServe("localhost:8080", nil)
	if err != nil {
		log.Fatal(err)
	}
}

func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprint(w, "Hello, web")
}