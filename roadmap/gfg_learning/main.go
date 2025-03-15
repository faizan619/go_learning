package main

// Structure
import (
	"fmt"
)

type Teacher struct{
	name string
	age int
	subject string
	exp int
	student Student
}

type Student struct{
	name string
	subject string
	class int
}

func main(){
	fmt.Println("Nested Structure Occured!")
	Teacher1 := Teacher{
		name : "Suman Jha",
		age :34,
		subject : "PHP",
		exp : 14,
		student: Student{"rishu","PHP",13},
	}
	fmt.Println(Teacher1)
	Teacher2 := Teacher{
		student: Student{"vikash","php",15},
	}
	fmt.Println(Teacher2)
}

// type Person struct {
// 	Name       string `json:"name"`
// 	Age        int    `json:"age"`
// 	Gender     string `json:"gender"`
// 	Profession string `json:"profession"`
// }

// func main() {
// 	fmt.Println("Geeks For Geeks Go Lang Learning")
// 	person1 := Person{"faizan", 22, "Male", "Software Developer"}
// 	fmt.Println("Person1 Details : ", person1)
// 	person2 := Person{Name:"salman",Gender:"Male"}
// 	fmt.Println(person2)
// 	fmt.Println(testing.Testing())
// }
