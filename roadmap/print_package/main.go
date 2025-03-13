package main

import "fmt"

func main() {

	// const name,id = "faizan alam",21
	// err := fmt.Errorf("User %q (id %d) not found",name,id)
	// fmt.Println(err)

	fmt.Println(" # This is normal Println # ")

	name := "faizan alam"
	age := 22
	isAdult := true
	marks := 65.3465
	// fmt.Printf("The name of User is : %[2]s and age is %[1]d \n",age,name)

	// fmt.Printf("The Name is : %s and age is %d \n",name,age)

	// user := fmt.Sprintf("The name is %s and age is %d \n",name,age)
	// fmt.Println(user)

	fmt.Printf("The name Type : %T \n",name);
	fmt.Printf("The name Value : %v \n",name);
	fmt.Printf("Is Faizan Adult : %t \n",isAdult);
	fmt.Printf("The age is : %d \n",age)
	fmt.Printf("The Marks Type : %T \n",marks)
	fmt.Printf("The Marks Value : %f \n",marks)
	fmt.Printf("The Marks Value[2] : %.2f \n",marks)s
	
}
