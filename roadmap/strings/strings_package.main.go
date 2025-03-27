// Strings Package
package main

import (
	"fmt"
	"strings"
)

func main(){
	fmt.Println("Strings Package in Go Lang")
	a1 := "faizan11"
	a2 := "faizan1"

	// Clone Function
	a1A := strings.Clone(a1)
	fmt.Println("Clone : ",a1A)

	// Compare Function
	// The result will be 0 if a == b, -1 if a < b, and +1 if a > b.
	fmt.Println("Compare : ",strings.Compare(a1,a1A))
	fmt.Println("Compare : ",strings.Compare(a1,a2))

	// Contains
	fmt.Println("Contain : ",strings.Contains(a1,"zan1"))

	// Count
	fmt.Println("Count : ",strings.Count(a1,""))
	fmt.Println("Count : ",strings.Count(a1,"a"))

	// Equal Folds
	fmt.Println("Equal Folds : ",strings.EqualFold(a1,a2))
	fmt.Println("Equal Folds : ",strings.EqualFold(a1,a1A))

	a3 := "      this         is        faizan"
	fmt.Println("Fields : ",strings.Fields(a3))

	// Index
	fmt.Println(strings.Index(a1,"an"))

	// Repeat 
	fmt.Println("Repeat : ba"+strings.Repeat("na",2))

	// Split
	fmt.Printf("Split : %q\n ",strings.Split("a,b,c",","))

	// Title,Upper,Lower
	fmt.Println("Title : ",strings.Title(a1))
	fmt.Println("Upper : ",strings.ToUpper(a1))
	fmt.Println("Lower : ",strings.ToLower(a1))

	// Trim [Space,left,right]
	a4 := "            hello world            "
	fmt.Println("Trim : ",strings.Trim(a4,"or"))
	a41 := "            hello world            "
	fmt.Println("Trim Space : ",strings.TrimSpace(a41))
	a42 := "            hello world            "
	fmt.Println("Trim Right : ",strings.TrimRight(a42," "))
	a43 := "            hello world            "
	fmt.Println("Trim Left : ",strings.TrimLeft(a43," "))

}
