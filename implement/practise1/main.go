// ############################### Part 4 - Done #############################
package main 

import "fmt"

func main(){
	fmt.Println("Arrays in Go Lang")
	var names[5] string 
	names[0] = "faizan"
	names[1] = "shifa"
	names[2] = "salman"
	names[3] = "zubiya"
	fmt.Println(names)

	var number = [5]int{22,44,66}
	fmt.Println("numbers : ",number)
	fmt.Println("Len Of Numbers : ",len(number))
	fmt.Println("Len Of names : ",len(names))

	var marks = []int{99,98,97,99,97,98}
	fmt.Println("Marks : ",marks)
	fmt.Println("Marks Len : ",len(marks))
	fmt.Println("Marks Capacity : ",cap(marks))
	marks = append(marks,44,55,33,44,55)
	fmt.Println("Append marks : ",marks)

}
// ############################### Part 3 - Done #############################

// package main 

// import "fmt"

// func addition(num1 int, num2 int) int{
// 	return num1+num2
// }

// func subtraction(num1 int, num2 int){
// 	res := num1-num2
// 	fmt.Println("Subtraction : ",res)
// }	

// func multiply(num1 int,num2 int) int{
// 	return num1*num2;
// }

// func divide(num1 int,num2 int) (float64, error) {
// 	if(num2 == 0){
// 		fmt.Println("Cannot Divide By Zero");
// 	} else{
// 		num3 := num1/num2;
// 		fmt.Println("Division : ",num3)
// 	}
// }

// func main(){
// 	fmt.Println("Function in GO Lang")
// 	number1 := 50
// 	number2 := 0
// 	fmt.Printf("Add, Sub, Mul and Div of number %d and %d\n",number1,number2)
// 	fmt.Println("Additon :",addition(number1,number2))
// 	// fmt.Println(subtraction(number1,number2))		// this is wrong method
// 	subtraction(number1,number2)
// 	fmt.Println("Multiple : ",multiply(number1,number2))
// 	divide(number1,number2)
	
// }

// ############################### Part 2 - Done #############################

// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main(){
// 	fmt.Printf("Enter User First Name : ")
// 	var name string;
// 	fmt.Scan(&name);
// 	fmt.Scanln();
// 	fmt.Println("User Name is : ",name)

// 	fmt.Printf("Enter Short User Details :")
// 	reader := bufio.NewReader(os.Stdin)
// 	desc, _ := reader.ReadString('\n')

// 	fmt.Println(name," short details is : ",desc)
// }

// ############################### Part 1 - Done #############################
// package main

// import (
// 	"fmt"

// 	"golang.org/x/example/hello/reverse"
// 	"rsc.io/quote"
// )

// func Greeting(name string, age int)string{
// 	message := fmt.Sprintf("Hello Mr/Ms :  %s. with age %d", name, age)

// 	return message;  
// }

// func main(){
// 	fmt.Println("_____________________________________________________")
// 	name := "shifa mulla"
// 	var age int= 22
// 	fmt.Printf("The Name of the User is : %s and age is %d  \n",name,age)
// 	fmt.Println(quote.Go())
// 	fmt.Println("Reverse : ",reverse.String(name))
// 	fmt.Println(Greeting(name,age))
// 	fmt.Println("Greet : ",hello("faizan alam"))
// 	fmt.Println("=====================================================")
// 	var pi float64 = 3.14
// 	fmt.Println("Value of Pi :",pi)
// 	fmt.Printf("Value of Pi : %.1f\n",pi)
// 	fmt.Printf("Value of Pi : %.2f\n",pi)

// }

// func hello(name string) string{
// 	message := fmt.Sprintf("Welcome to Go lang Programming Mr/Ms : %s\n",name)
// 	return message
// }
