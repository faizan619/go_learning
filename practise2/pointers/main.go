package main

import "fmt"

func modifyValueByReference(a *int){
	*a = *a + 20

}


func main(){
	fmt.Println("Pointers in Go")

	// var num int
	// num = 2

	// var ptr *int
	// ptr = &num

	// or 
	num := 2
	ptr := &num

	fmt.Println("Num has Value  : ",num)
	fmt.Println("Pointers :",ptr)
	fmt.Println("Data Contain Through Pointers :",*ptr)

	var pointer *string
	if pointer == nil {
		fmt.Println("Pointer is not defined!")
	}


	value := 10
	modifyValueByReference(&value)
	fmt.Println("Value Containers : ",value)
}