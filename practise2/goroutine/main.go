package main

import (
	"fmt"
	"time"
)

func sayHello(){
	fmt.Println("Hello World")
	time.Sleep(2000 * time.Millisecond)	// simulating some work
	fmt.Println("SayHello Function Ended")
}

func sayHi(){
	fmt.Println("Hi Faizan Alam!")
}

func main() {
	fmt.Println("Goroutine in Go lang")
	go sayHello()
	sayHi()

	// Wait for the moment to allow the goroutine to finish
	// time.Sleep(1000 * time.Millisecond)
	time.Sleep(3000 * time.Millisecond)
}