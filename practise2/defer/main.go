package main

import "fmt"

func main() {
	fmt.Println("Defer in Go lang")

	fmt.Println("Start of the Program")
	data := 55
	defer fmt.Println("Date :",data)
	defer fmt.Println("Middle of the Program")
	fmt.Println("End of the Program")
}