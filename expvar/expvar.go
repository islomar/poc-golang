package main

import (
	"expvar"
	"time"
	"net/http"
	"log"
)

// Monitoring: export variables via an HTTP handler registered at /debug/vars (http://localhost:8080/debug/vars)
func main() {
	count := expvar.NewInt("count")
	go func() {
		for {
			count.Add(1)
			time.Sleep(time.Second)
		}
	}()
	log.Fatal(http.ListenAndServe("localhost:8080", nil))
}