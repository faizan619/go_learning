package main

import (
	"fmt"

	"golang.org/x/example/hello/reverse"
	"rsc.io/quote"
)

func Greeting(name string, age int)string{
	message := fmt.Sprintf("Hello Mr/Ms :  %s. with age %d", name, age)

	return message;  
}

func main(){
	fmt.Println("Hello world")
	name := "faizan alam"
	age := 22
	fmt.Printf("The Name of the User is : %s and age is %d  \n",name,age)
	fmt.Println(quote.Go())
	fmt.Println("Reverse : ",reverse.String(name))
	fmt.Println(Greeting(name,age))
}