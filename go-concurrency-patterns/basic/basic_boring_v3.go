package main

import (
	"fmt"
	"math/rand"
	"time"
)

func main() {
	go boringV3("boring!")
	fmt.Println("I'm listening.")
	time.Sleep(2 * time.Second)
	fmt.Println("You're boring; I'm leaving.")
}

func boringV3(msg string) {
	for i := 0; ; i++ {
		fmt.Println(msg, i)
		time.Sleep(time.Duration(rand.Intn(1e3)) * time.Millisecond)
	}
}