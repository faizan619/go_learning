package main

// If Else Condition 
import "fmt"

func main(){
	fmt.Println("If Else Condition in GO Lang")
	x := 10
	// if x == 0{
	// 	fmt.Println("The Value of X is  : ",x)
	// }

	if x > 20{
		fmt.Println("X Value si Greater than 20")
	} else if x > 5{
		fmt.Println("X value is greater than 5")
	} else{
		fmt.Println("X Value is Smaller than 5")
	}

}

// Arrays in Go Lang
// import "fmt"

// func main(){
// 	fmt.Println("Arrays in Go Lang")

// 	var names[5] string
// 	names[0] = "faizan"
// 	names[2] = "salman"
// 	names[4] = "maroof"
// 	names[1] = "shifa"
// 	fmt.Println("Names of Friends ",names)



// 	var numbers = [5]int{1,2,3,4,5}
// 	fmt.Println("Numbers are : ",numbers)

// 	fmt.Println("length of An Array : ",len(numbers))

// 	fmt.Println("Value of Name at 2nd Index : ",names[2])




// 	// var prices[5]int
// 	var prices[5]string
// 	prices[0] = "example"
// 	fmt.Println("Prices are : ",prices)			// what is shown in empty array with int and string
// 	fmt.Printf("Prices are %q",prices)
// }

// ################################################################################################
// Arrays in Go Lang

// ################################################################################################
// Error Handling in Go
// import "fmt"

// // func divide(a,b float64) (float64, error) {
// // 	if(b == 0){
// // 		return 0, fmt.Errorf("Dominator must not be Zero")
// // 	}
// // 	return a / b, nil
// // }
// func divide(a,b float64) (float64, string) {
// 	if(b == 0){
// 		return 0, "Dominator must not be Zero"
// 	}
// 	return a / b, ""
// }

// func main(){
// 	fmt.Println("Error Handling in Go");
// 	// data := divide(10,5)

// 	// data, err := divide(10,5)
// 	// if(err != nil){
// 	// 	fmt.Println("Error Handling");
// 	// }
// 	// fmt.Println("data : ",data);

// 	data, _ := divide(10,5)
// 	fmt.Println("data : ",data);

// }

// ################################################################################################
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



// ################################################################################################
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