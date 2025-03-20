// Practise on 20 March 2025

package main

import "fmt"

func main(){
	fmt.Println("Pointer on Go Lang")
	a1 := "faizan"

	var a2 *string;
	a2 = &a1
	fmt.Println("Value of A2 : ",a2)
	fmt.Println("Address of A2 : ",&a2)
	fmt.Println("Reference of A2 : ",*a2)
}


// package main

// import "fmt"

// func main(){
// 	fmt.Println("Slices in Go Lang")
// 	arr1 := []int{2,4,6,8,10,12,14,16,18,20}
// 	arrSlice := arr1[2:5]
// 	fmt.Println("Arr1 : ",arr1)
// 	fmt.Println("Arr1 Len : ",len(arr1))
// 	fmt.Println("Arr1 Cap : ",cap(arr1))
// 	fmt.Println("Arr Slice : ",arrSlice)
// 	fmt.Println("Arr Slice Len : ",len(arrSlice))
// 	fmt.Println("Arr Slice Cap : ",cap(arrSlice))

// 	arr2 := make([] int, 0, 5)
// 	arr2 = append(arr2, 5,5,6,7,8,9)
// 	fmt.Println("Arr 2 :",arr2)
// 	fmt.Println("Arr 2 Len :",len(arr2))
// 	fmt.Println("Arr 2 Cap :",cap(arr2))

// }



// package main

// import "fmt"

// type Person struct {
// 	name  string
// 	age   int
// 	hobby string
// }

// func main() {
// 	fmt.Println("Struct and Slices in Go lang")
// 	person1 := Person{"faizan alam",22,"Coding and Travelling"}
// 	fmt.Println(person1)
// 	person2 := Person{name:"Salman",hobby:"Not mention"}
// 	fmt.Println(person2)


// }

// Practise on 19 March 2025
// package main

// import "fmt"

// func main(){
// 	fmt.Println("Maps in Go Lang")

// 	studentNames := make(map[int]string)
// 	studentNames[22] = "Faizan alam"
// 	studentNames[24] = "Mulla Shifa"
// 	studentNames[21] = "Salman Sayyed"
// 	studentNames[23] = "Maroof Baig"
// 	fmt.Println(studentNames[22])

// 	name,exists := studentNames[22]
// 	fmt.Println("Name of Student : ",name," Does Student Exists : ",exists)

// 	friendsNames := map[int]string{
// 		1:"faizan",
// 		2:"alam",
// 		3:"sajjad",
// 		4:"Alm",
// 	}
// 	fmt.Println(friendsNames)
// }

// Practise on 18 March 2025

// package main

// import "fmt"

// func main() {
// 	fmt.Println("Loop in Go lang")
// 	var num int;
// 	for{
// 		fmt.Printf("Enter Number (press 0 to Exit) : ")
// 		fmt.Scan(&num)
// 		fmt.Println("Number : ",num);
// 		if(num==0){
// 			break
// 		}
// 	}
// }

// package main

// import (
// 	// "bufio"
// 	"fmt"
// 	// "os"
// )

// func Greeting(name string) {
// 	fmt.Println("Welcome to Go Lang Learning Mr/Ms : ", name)
// }

// func EligibleVoting(age int) bool {
// 	// if age > 18 {
// 	// 	return true
// 	// }
// 	// return false
// 	return age >= 18
// }

// func main() {
// 	fmt.Println("Practise in Go Lang")
// 	// var name string
// 	// // name = "Faizan alam"
// 	// var age int
// 	// // age = 22
// 	// fmt.Printf("Enter Your Name : ")
// 	// reader := bufio.NewReader(os.Stdin)
// 	// name, err := reader.ReadString('\t')
// 	// fmt.Printf("Enter Your Age : ")
// 	// fmt.Scan(&age)
// 	// if err != nil {
// 	// 	fmt.Println("Error While Reading Input : ", err)
// 	// 	return
// 	// }
// 	// fmt.Println("User Name : ", name, " and age is  : ", age)
// 	// Greeting(name)

// 	// fmt.Printf("Type of Function : %T \n",EligibleVoting(age))
// 	// fmt.Println("Value of Function :",EligibleVoting(age))

// 	// if EligibleVoting(age) {
// 	// 	fmt.Println("User is Eligible for Voting")
// 	// 	} else {
// 	// 		fmt.Println("User is Not Eligible for Voting")
// 	// 		switch age{
// 	// 		case 15:
// 	// 			fmt.Println("Three Year More Required for Voting Eligiblity")
// 	// 		case 16:
// 	// 			fmt.Println("Two Year More Required for Voting Eligiblity")
// 	// 		case 17:
// 	// 			fmt.Println("One Year More Required For Voting Eligiblity")
// 	// 		default:
// 	// 			fmt.Println("There is Still Many Year Left For Voting Eligiblity")
// 	// 		}
// 	// }

// 	friends := [11]string{"maroof","salman","rishu","giri","asad"}
// 	// defer fmt.Println("First Friend : ",friends[0])
// 	// defer fmt.Println("Second Friend : ",friends[1])
// 	// defer fmt.Println("All Friend : ",friends)
// 	// fmt.Println("=====================================")
// 	// for _,name := range friends{
// 	// 	fmt.Println("The Friends Names are : ",name)
// 	// }
// 	// fmt.Println("=====================================")
// 	// counter := 0
// 	// for {
// 	// 	fmt.Println("For Loop : ",friends[counter])
// 	// 	counter++
// 	// 	if counter >= len(friends){
// 	// 		break
// 	// 	}
// 	// }
// 	fmt.Println("===================== Inserting Start =========================")
// 	fmt.Println()

// 	otherFriend := []string{"shifa","zubiya","Janvi","Sandhiya","Aneeza","Anni"}
// 	fmt.Println(otherFriend)
// 	fmt.Println("Length of Friends : ",len(friends))
// 	fmt.Println("Length of Other Friends : ",len(otherFriend))
// 	for i:=0;i<len(otherFriend);i++{
// 		// fmt.Println("Len Val : ",friends[len(friends)-1])
// 		friends[len(friends)-len(otherFriend)+i] = otherFriend[i]
// 	}
// 	fmt.Println(friends)
// 	// for index, _ := range otherFriend{
// 		// friends[len(friends)+index] = girl
// 		// fmt.Println("List : ",friends[index])
// 	// }
// 	// fmt.Println("Friends List : ",friends)

// 	fmt.Println()
// 	fmt.Println("===================== Inserting End =========================")
// }

// Practise on 17 March 2025
// package main

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	fmt.Println("Practise in Go Lang")
// 	fmt.Printf("Enter you Name : ")
// 	var name string;
// 	fmt.Scanln(&name)
// 	fmt.Println("Welcome to The GO Lang Mr/Ms : ",name)

// 	// var name2 string;
// 	fmt.Printf("Enter Other Name : ")
// 	reader := bufio.NewReader(os.Stdin)
// 	name2,_ := reader.ReadString('\n')
// 	fmt.Println("Name 2 : ",name2)

// }
