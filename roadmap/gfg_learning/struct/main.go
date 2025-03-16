package main

import "fmt"

type employee struct{
	emp_id string
	name string
	age int
	gender string
	salary int
}

func main() {
	fmt.Println("Structure in Go Lang")
	 emp1 := employee{"001","faizan alam",22,"male",54000}
	 fmt.Println(emp1)
	 emp2 := employee{"002","salman sayyed",21,"male",65000}
	 fmt.Println(emp2)
}