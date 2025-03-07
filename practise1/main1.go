// package main1

import (
	"fmt"
	"golang.org/x/example/hello/reverse"
	"rsc.io/quote"
)

func Hello1(name string, age int) string { 
	// message := fmt.Sprintf("Hello Mr/Ms :  %v. with age %d",name,age)
	message := fmt.Sprintf("Hello Mr/Ms :  %s. with age %d", name, age)
	return message
}

func main() {
	fmt.Println("Welcome to GO programming")
	fmt.Println(quote.Go())
	fmt.Println("==========================================")
	fmt.Println(Hello1("faizan", 22))
	fmt.Println("==========================================")
	fmt.Println(reverse.String("faizan"))
}

// CLI Commands
// go mod tidy
// go mod init myLearning
// go get golang.org/x/example/hello/reverse


// Formet Specifer
// Type				Format Specifier	Example Usage
// String				%s					fmt.Sprintf("%s", "Hello")
// Integer				%d					fmt.Sprintf("%d", 123)
// Floating Point		%f					fmt.Sprintf("%f", 3.14)
// Boolean				%t					fmt.Sprintf("%t", true)
// Hexadecimal			%x					fmt.Sprintf("%x", 255)
// Character			%c					fmt.Sprintf("%c", 65)
// Pointer				%p					fmt.Sprintf("%p", &var)
// Scientific Notation	%e					fmt.Sprintf("%e", 12345.6789)
// package main

// import "fmt"

// func main() {
//     name := "John"
//     age := 30
//     salary := 75000.50
//     isActive := true

//     // Using different format specifiers
//     fmt.Printf("Name: %s\n", name)           // String
//     fmt.Printf("Age: %d\n", age)             // Integer
//     fmt.Printf("Salary: %.2f\n", salary)     // Floating-point with 2 decimal precision
//     fmt.Printf("Is Active: %t\n", isActive)  // Boolean
//     fmt.Printf("Age in Hex: %x\n", age)      // Hexadecimal (not commonly used for age)
// }
