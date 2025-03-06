package main 

import (
	"fmt"
	"strings"
)

func main(){
	// str := "apple,orange,banana,grapes"
	// parts := strings.Split(str,",")
	str1 := "apple.orange.banana.grapes"
	parts := strings.Split(str1,".")
	fmt.Println("Split : ",parts)


	str2 := "one two three four five six"
	count := strings.Count(str2, "two")
	fmt.Println("Count : ",count)

	str3 := "        Hello      World       "
	trim := strings.TrimSpace(str3)
	fmt.Println("Trimmed : ",trim)


	str4A := "Faizan"
	str4B := "Alam"
	result := strings.Join([]string{str4A,str4B,"Sajjad"}," ")
	fmt.Println("Result : ",result)
}