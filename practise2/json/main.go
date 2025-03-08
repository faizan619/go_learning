package main

import (
	"encoding/json"
	"fmt"
)

type Person struct {
	Name    string `json:"name"`
	Age     int    `json:"age"`
	// IsAdult bool   `json:"isAdult"`
	IsAdult bool   // aise bhi chalega
}

func main() {
	fmt.Println("Json in Go Lang")
	person := Person{Name: "Faizan",Age :22,IsAdult:true}
	fmt.Println("Person Data : ",person)

	// convert json into json encoding (marshalling)
	jsonData,err := json.Marshal(person)
	if err  != nil{
		fmt.Println("Error marshalling : ",err)
		return
	}
	fmt.Println("Slice of Data : ",string(jsonData))

	var personData Person
	err1 := json.Unmarshal(jsonData,&personData)
	if err1 != nil{
		fmt.Println("Error Unmarshalling : ",err)
		return
	}
	fmt.Println("Person Data : ",personData)

}
