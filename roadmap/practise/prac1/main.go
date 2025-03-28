package main

import (
	"fmt"
	"strconv"
)

func main() {

	num := "54"
	fmt.Println("Num Val : ",num)
	fmt.Printf("Num Type : %T\n",num)

	str2int, _ := strconv.Atoi(num)
	fmt.Println("Num Val : ",str2int)
	fmt.Printf("Num Type : %T\n",str2int)
	
	num_str := "45.34"
	fmt.Println("Num str  :",num_str)
	fmt.Printf("Num str : %T\n",num_str)
	int2float, _ := strconv.ParseFloat(num_str,64)
	fmt.Println(int2float)
	fmt.Printf("val here : %T\n ",int2float)




	// fmt.Println("Practise 1 in Go Lang")

	// arr1 := []int{1,2,3,4,5,6,7,8}
	// fmt.Println("Array 1 : ",arr1)
	// fmt.Println("Len Array :", len(arr1) )
	// fmt.Println("Cap Array :", cap(arr1) )

	// arr2 := [][]int {{1,2,3,4},{2,3,4,5}}
	// fmt.Println(arr2)
	// fmt.Println(len(arr2))
	// fmt.Println(cap(arr2))
}