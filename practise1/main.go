package main

// Error Handling in Go
import "fmt"

func divide(a,b float64) float64 {
	return a / b;
}

func main(){
	fmt.Println("Error Handling in Go");
	data := divide(10,5)
	fmt.Println("data : ",data);

}

// How to Use Functions in GO

// import "fmt"

// func name(name string){
// 	fmt.Println("Welcome Mr/Ms : ",name)
// }

// func add(a,b int) int{
// 	// var c = a+b;
// 	// fmt.Printf("The Addition of %d and %d is %d",a,b,c);
// 	return a + b;
// }

// func main() {
// 	fmt.Println("Welcome to Go programming")
// 	name("faizan alam");

// 	// add(10,30)
// 	var data = add(10,30)
// 	fmt.Println("Add : ",data)
// }



// How to Take INput From Users

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	fmt.Println("Taking Input from User")
// 	fmt.Print("Enter Your Name : ");

// 	// for single letter
// 	// var name string;
// 	// fmt.Scan(&name)
// 	// fmt.Println("Hello Mr/Ms : ",name)
// 	// fmt.Printf("Hello Mr/Ms : %s ",name)

// 	// for multiple letter
// 	reader := bufio.NewReader(os.Stdin)
// 	name, _ := reader.ReadString('\n')
// 	fmt.Println("Hello Mr: ",name)
// }