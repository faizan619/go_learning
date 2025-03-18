package main

import "fmt"

func addition(num1 int, num2 int) {
	num3 := num1 + num2
	fmt.Printf("Addition -> %d + %d = %d\n", num1, num2, num3)
}

func subtraction(num1 int, num2 int) {
	num3 := num1 - num2
	fmt.Printf("Subtraction -> %d - %d = %d\n", num1, num2, num3)
}

func multiplication(num1 int, num2 int) {
	num3 := num1 * num2
	fmt.Printf("Multiplication -> %d * %d = %d\n", num1, num2, num3)
}

func division(num1 int, num2 int) {
	if num2 != 0 {
		num3 := num1 / num2
		fmt.Printf("Division -> %d / %d = %d\n", num1, num2, num3)
	} else {
		fmt.Println("Cannot Divide By Zero")
	}
}

func main() {
	fmt.Println("Command Line Calculator")
	var num1 int
	var num2 int
	var res int
	for {
		fmt.Println("Welcome to Cli Calculator")
		fmt.Printf("Enter Number 1 : ")
		fmt.Scan(&num1)
		fmt.Printf("Enter Number 2 : ")
		fmt.Scan(&num2)
		for {
			fmt.Println("============================================")
			fmt.Println("1 - Addition")
			fmt.Println("2 - Subtraction")
			fmt.Println("3 - Multiplication")
			fmt.Println("4 - Division")
			fmt.Println("5 - Quite")
			fmt.Printf("What Operation You Want to Perform : ")
			fmt.Scan(&res)
			if res == 5{
				fmt.Println("")
				break
			}
			switch res {
			case 1:
				addition(num1, num2)
			case 2:
				subtraction(num1, num2)
			case 3:
				multiplication(num1, num2)
			case 4:
				division(num1, num2)
			default:
				fmt.Println("Please Select the Given Number Only!")
			}
		}
		fmt.Printf("Do You Want to Continue (1- Yes, any number to exit) : ")
		fmt.Scan(&res)
		if res == 1{
			continue
		} else{
			break
		} 
	}
}
