package main

import (
	"fmt"
	"strconv"
)

func main(){
	fmt.Println("Welcome To Data Conversion")

	var num int = 42
	fmt.Println("Number is  :", num)
	fmt.Printf("Types of Num is : %T\n",num)

	var data float64 = float64(num)
	data = data +1.34
	fmt.Println("Data is : ",data)
	fmt.Printf("Type of Data is %T\n",data)


	num = 123
	str := strconv.Itoa(num)
	fmt.Println("Data is : ",str)
	fmt.Printf("Type of Data is %T\n",str)

	number := "1234"
	fmt.Println("Data is : ",number)
	fmt.Printf("Type of Data is %T\n",number)

	number_string := "1234"
	number_int, _ := strconv.Atoi(number_string)
	fmt.Println("Data is : ",number_int)
	fmt.Printf("Type of Data is %T\n",number_int)

	num_string := "3.14"
	number_float, _ := strconv.ParseFloat(num_string,64)
	fmt.Println("Data is : ",number_float)
	fmt.Printf("Type of Data is %T\n",number_float)



}