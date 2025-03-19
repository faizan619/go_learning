// Practise on 19 March 2025
package main 

import "fmt"

func main(){
	fmt.Println("Maps in Go Lang")

	studentNames := make(map[int]string)
	studentNames[22] = "Faizan alam"
	studentNames[24] = "Mulla Shifa"
	studentNames[21] = "Salman Sayyed"
	studentNames[23] = "Maroof Baig"
	fmt.Println(studentNames[22])

	name,exists := studentNames[22]
	fmt.Println("Name of Student : ",name," Does Student Exists : ",exists)

	friendsNames := map[int]string{
		1:"faizan",
		2:"alam",
		3:"sajjad",
		4:"Alm",
	}
	fmt.Println(friendsNames)
}

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
