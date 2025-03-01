package main

// Maps in Go Lang
import "fmt"

func main() {
	fmt.Println("Maps in Go Lang")

	studentGrades := make(map[string]int)
	studentGrades["Faizan"] = 98
	studentGrades["maroof"] = 92
	studentGrades["Salman"] = 95
	studentGrades["Shifa"] = 99

	fmt.Println("Marks of Salman : ", studentGrades["Salman"])
	studentGrades["Salman"] = 100
	fmt.Println("Marks of Salman : ", studentGrades["Salman"])
	delete(studentGrades, "Salman")
	fmt.Println("Marks of Salman : ", studentGrades["Salman"])

	grades, exists := studentGrades["Zubiya"]
	fmt.Println("Grades of Zubiya : ", grades)
	fmt.Println("Does Zubiya Exists", exists)

	for index, value := range studentGrades {
		fmt.Printf("key is %s and Marks is %d\n", index, value)
	}

	childGrades := map[string]int{
		"Faizan": 99,
		"Salman": 95,
		"Maroof": 55,
		"Shifa":  89,
	}

	for index, value := range childGrades {
		fmt.Printf("----------key is %s and Marks is %d\n", index, value)
	}
}

// Loops in Go Lang
// import "fmt"

// func main(){
// 	fmt.Println("Loops in Go lang")

// 	// // Exmaple 1
// 	// for i:=0;i<10;i++{
// 	// 	fmt.Println("Loops I : ",i)
// 	// }

// 	// // Example 2
// 	// counter := 1;
// 	// for{
// 	// 	fmt.Println("Looping ",counter)
// 	// 	counter++
// 	// 	if(counter > 10){
// 	// 		break
// 	// continue
// 	// 	}
// 	// }

// 	// for{
// 	// 	fmt.Println("Infinite Loop")
// 	// }

// 	// Exmaple 3
// 	numbers := []int{5,6,4,3,7,8,2,1,9}
// 	for index, value := range numbers {
// 		fmt.Printf("Index : %d, Value : %d\n",index,value)
// 	}

// 	data := "Faizan alam"
// 	for index, value := range data{
// 		// fmt.Println(index, "-> Value : ",value)		// it giving something wierd
// 		fmt.Printf("Index : %d, Value : %c\n",index,value)
// 	}

// }

// Switch Case in Go Lang
// import "fmt"

// func main(){
// 	fmt.Println("Switch Case in Go Lang")

// 	// example 1
// 	day :=3
// 	switch day {
// 	case 1:
// 		fmt.Println("Day is Monday")
// 	case 2:
// 		fmt.Println("Day is Monday")
// 	case 3:
// 		fmt.Println("Day is Monday")
// 	case 4:
// 		fmt.Println("Day is Monday")
// 	default:
// 		fmt.Println("Days are WeekEnds")
// 	}

// 	month := "January"
// 	switch month {
// 	case "January":
// 		fmt.Println("New year Celebration")
// 	case "February":
// 		fmt.Println("28 Days Celebreation")
// 	case "March":
// 		fmt.Println("It is March")
// 	default:
// 		fmt.Println("It is not jan, feb, march Month")
// 	}

// 	number := 7
// 	switch number {
// 	case 3,5,7,9:
// 		fmt.Println("This are Odd Numbers")
// 	case 2,4,6,8:
// 		fmt.Println("This are Even Numbers")
// 	}

// 	var temp = 34
// 	switch {
// 	case temp < 30:
// 		fmt.Println("Temperature is less than 30")
// 	case temp >30 && temp < 45:
// 		fmt.Println("Temperature is Managable")
// 	case temp>45:
// 		fmt.Println("It is Sooo Hot")

// 	}
// }

// If Else Condition
// import "fmt"

// func main(){
// 	fmt.Println("If Else Condition in GO Lang")
// 	x := 10
// 	// if x == 0{
// 	// 	fmt.Println("The Value of X is  : ",x)
// 	// }

// 	if x > 20{
// 		fmt.Println("X Value si Greater than 20")
// 	} else if x > 5{
// 		fmt.Println("X value is greater than 5")
// 	} else{
// 		fmt.Println("X Value is Smaller than 5")
// 	}

// }

// Slices in Go Lang

// import "fmt"

// func main(){
// 	fmt.Println("Slices in Go Lang")

// 	// declare and initialize a slice
// 	// number := []int {1,2,3,4}

// 	// fmt.Println("Slice : ",number)
// 	// fmt.Println("Length : ",len(number))
// 	// fmt.Println("Capacity : ",cap(number))

// 	// fmt.Println("Slices : ",number)
// 	// fmt.Println("Len Of the Slices : ",len(number))
// 	// number = append(number, 11,22,33,44,55,66,77,88,99)

// 	// fmt.Println("Slices : ",number)
// 	// fmt.Println("Len Of the Slices : ",len(number))
// 	// fmt.Printf("Number has Data Type : %T\n",number)

// 	// names := []string{"faizan","sajjad","salman","maroof","shifa"}
// 	// fmt.Println("Names are : ",names)

// 	// slice - len and capacity
// 	numbers := make([]int, 3, 5)
// 	numbers = append(numbers, 22)
// 	numbers = append(numbers, 98)
// 	// numbers = append(numbers, 9)
// 	fmt.Println("Slice : ",numbers)
// 	fmt.Println("Length : ",len(numbers))
// 	fmt.Println("Capacity : ",cap(numbers))

// 	stock := make([]int, 0)
// 	stock = append(stock, 5)
// 	fmt.Println("Slice : ",stock)
// 	fmt.Println("Length : ",len(stock))
// 	fmt.Println("Capacity : ",cap(stock))

// }

// Arrays in Go Lang
// import "fmt"

// func main(){
// 	fmt.Println("Arrays in Go Lang")

// 	var names[5] string
// 	names[0] = "faizan"
// 	names[2] = "salman"
// 	names[4] = "maroof"
// 	names[1] = "shifa"
// 	fmt.Println("Names of Friends ",names)

// 	var numbers = [5]int{1,2,3,4,5}
// 	fmt.Println("Numbers are : ",numbers)

// 	fmt.Println("length of An Array : ",len(numbers))

// 	fmt.Println("Value of Name at 2nd Index : ",names[2])

// 	// var prices[5]int
// 	var prices[5]string
// 	prices[0] = "example"
// 	fmt.Println("Prices are : ",prices)			// what is shown in empty array with int and string
// 	fmt.Printf("Prices are %q",prices)
// }

// ################################################################################################
// Arrays in Go Lang

// ################################################################################################
// Error Handling in Go
// import "fmt"

// // func divide(a,b float64) (float64, error) {
// // 	if(b == 0){
// // 		return 0, fmt.Errorf("Dominator must not be Zero")
// // 	}
// // 	return a / b, nil
// // }
// func divide(a,b float64) (float64, string) {
// 	if(b == 0){
// 		return 0, "Dominator must not be Zero"
// 	}
// 	return a / b, ""
// }

// func main(){
// 	fmt.Println("Error Handling in Go");
// 	// data := divide(10,5)

// 	// data, err := divide(10,5)
// 	// if(err != nil){
// 	// 	fmt.Println("Error Handling");
// 	// }
// 	// fmt.Println("data : ",data);

// 	data, _ := divide(10,5)
// 	fmt.Println("data : ",data);

// }

// ################################################################################################
// How to Use Functions in GO

// import "fmt"

// func name(name string){
// 	fmt.Println("Welcome Mr/Ms : ",name)
// }

// func add(a,b int) int{
// 	// var c = a+b;
// 	// fmt.Printf("The Addition of %d and %d is %d",a,b,c);
// 	return a + b;
// }

// func main() {
// 	fmt.Println("Welcome to Go programming")
// 	name("faizan alam");

// 	// add(10,30)
// 	var data = add(10,30)
// 	fmt.Println("Add : ",data)
// }

// ################################################################################################
// How to Take INput From Users

// import (
// 	"bufio"
// 	"fmt"
// 	"os"
// )

// func main() {
// 	fmt.Println("Taking Input from User")
// 	fmt.Print("Enter Your Name : ");

// 	// for single letter
// 	// var name string;
// 	// fmt.Scan(&name)
// 	// fmt.Println("Hello Mr/Ms : ",name)
// 	// fmt.Printf("Hello Mr/Ms : %s ",name)

// 	// for multiple letter
// 	reader := bufio.NewReader(os.Stdin)
// 	name, _ := reader.ReadString('\n')
// 	fmt.Println("Hello Mr: ",name)
// }
