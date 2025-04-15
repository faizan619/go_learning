package main

import "fmt"

func main(){
	fmt.Println("Learn Go Language in this.")
	fmt.Println("===============================")
	fmt.Print("Faizan ")
	fmt.Print("Alam")
}

package main

import (
	"fmt"
	"reflect"
	"time"
)

var name = "faizan"

// sirname := "alam"	// it gives error

func main() {
	var sent = "my name is faizan \nI am graduated student from mumbai \nI am Single"
	fmt.Println(sent)
	fmt.Println(reflect.TypeOf(sent))
	var quate = `"never trust anyone \nI repeat never"
			got hated one.
			But not twice
			- faizan `
	fmt.Println(quate)
	fmt.Printf("%T\n", quate)
	// var unused = "not using this variable"
	// _ = unused	// now it will not give not used error
	// fmt.Println(name)
	//fmt.Println(sirname)
	// var num float64
	// var str string
	// str1 := "hello world"
	// fmt.Println(str1)
	// fmt.Println(num)
	// fmt.Println(str)
	// fmt.Println("Hello! Welcome to the World")
	fmt.Println("Current Time : ", time.Now())
}
