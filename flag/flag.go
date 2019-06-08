package main

import (
	"fmt"
	"flag"
	"time"
)

var (
	message = flag.String("message", "Hello!", "what to say")
	delay   = flag.Duration("delay", 2*time.Second, "how long to wait")
)

// To execute it: flag -message 'Hold on...' -delay 5m
func main() {
	flag.Parse()
	fmt.Println(*message)
	time.Sleep(*delay)
}