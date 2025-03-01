package main

import "fmt"

// Example 1
type Person struct{
	FirstName string
	LastName string
	Age int
}

// example 2
type Contact struct{
	Email string
	Phone int
	link string
}

type Address struct{
	House int
	Area string
	State string
}

type Employee struct{
	Person_Details Person
	Person_Contact Contact
	Person_Address Address
}


func main(){

	// // Example 1
	// fmt.Println("Struct in Go Lang")
	// var faizan Person
	// fmt.Println("Faizan Person : ",faizan)
	// faizan.FirstName = "Faizan"
	// faizan.LastName = "ALam"
	// faizan.Age = 22
	// fmt.Println(faizan)

	// person1 := Person{
	// 	FirstName: "Salman",
	// 	LastName: "Sayyed",
	// 	Age: 21,
	// }
	// fmt.Println(person1)

	// var person2 = new(Person)
	// person2.FirstName = "Maroof"
	// person2.LastName = "Baig"
	// person2.Age = 23
	// fmt.Println(person2)

	// fmt.Println("Age of Person1 : ",person1.Age)


	// Example 2
	// employee1 := Employee
	var employee1 Employee
	employee1.Person_Details = Person{
		FirstName: "Rishu",
		LastName: "Singh",
		Age: 20,
	}
	employee1.Person_Contact = Contact{
		Email: "rishu@gmail.com",
		Phone: 9983774454,
	}
	employee1.Person_Address = Address{
		House: 401107,
		Area: "Vasai Road",
		State: "Maharashtra",
	}

	fmt.Println("Employee 1 : ",employee1 )
	fmt.Println("Employee 1 Address : ",employee1.Person_Address )
	fmt.Println("Employee 1 Address Area: ",employee1.Person_Address.Area )

	employee1.Person_Contact.link = "example.com"
	fmt.Println("Employee 1 : ",employee1)


}