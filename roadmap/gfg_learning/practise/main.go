package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {
	fmt.Println("Practise in Go Lang")
	fmt.Printf("Enter you Name : ")
	var name string;
	fmt.Scanln(&name)
	fmt.Println("Welcome to The GO Lang Mr/Ms : ",name)

	// var name2 string;
	fmt.Printf("Enter Other Name : ")
	reader := bufio.NewReader(os.Stdin)
	name2,_ := reader.ReadString('\n')
	fmt.Println("Name 2 : ",name2)

}