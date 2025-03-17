package main

import "fmt"

type Person struct{
	name string 
	age int
}

func main(){
	// Struct with pointers
	person1 := Person{"faizan",22}
	fmt.Println(person1)

	person1Pointer := &person1
	fmt.Println("Person1Pointer : ",person1Pointer)
	fmt.Println("*Person1Pointer : ",*person1Pointer)
	fmt.Println("Person1Pointer Name : ",person1Pointer.name)
	fmt.Println("Person1Pointer Age : ",person1Pointer.age)

	// Creating pointer using basic method


	// Creating pointer using new keyword
	person2 := new(Person)
	person2.name = "salman"
	person2.age = 21
	fmt.Println(person2)
	fmt.Println(*person2)
}



// func ptr(a *int){
// 	*a = 30
// }

// func main() {

// 	fmt.Println("Pointers in Go Lang")

// 	// Pointer value passing in functions
// 	var age = 22
// 	fmt.Println("Value stored in age before calling function : ",age)
// 	fmt.Println("& : ",&age)
// 	var page *int = &age
// 	fmt.Println("* : ",*page)
	
// 	ptr(page)
// 	fmt.Println("Value stored in age after calling function : ",age)


	// basic pointers concept

// 	x := 0xFF
// 	// y := 0x9C

// 	fmt.Printf("Type of Variable of x : %T\n",x)
// 	fmt.Printf("Hexadecimial of x : %X\n",x)
// 	fmt.Printf("Value of x : %d\n",x)

// 	fmt.Println("====================================")

// 	a1 := "faizan"

// 	// var a2 *string = &a1;
// 	var a2 *string

// 	fmt.Println("Before Referencing : ",a2)
// 	a2 = &a1;


//    // displaying the result
//    fmt.Println("Value stored in a1 = ", a1)
//    fmt.Println("Address of a1 = ", &a1)
//    fmt.Println("Value stored in variable a2 = ", a2)
//    fmt.Println("Using Deferencing Variable a2 = ",*a2)

// //    changing the value of assigning
// 	*a2 = "salman"
// 	fmt.Println("Value stored in a1(*a2) after changing : ",a1)


// }