package main

import "fmt"

func main() {
	fmt.Println("Arrays")
	arr1 := [3]int{1, 3, 5}
	fmt.Println(arr1)
	arr2 := arr1
	fmt.Println(arr2)
	var arr3 = [len(arr1)]int{}
	fmt.Println("Before : ",arr3)
	for i:=0;i<len(arr1);i++{
		arr3[i] = arr1[i]
	}
	fmt.Println("After : ",arr3)
	
	// for i:=0;i<len(arr1);i++{
	// 	fmt.Println(" Array ",i,": ",arr1[i])
	// }
}

// import(
// 	"fmt"
// )

// func add(nums ...int)int{
// 	total := 0
// 	for _, val := range nums{
// 		total += val;
// 	}
// 	return total;
// }

// func users(names ...string){
// 	for _, name := range names{
// 		fmt.Printf("%s - ",name)
// 	}
// 	fmt.Println()
// }
// func main(){
// 	fmt.Println("Verdic Functions")
// 	fmt.Println("School Friends")
// 	users("faizan","salman","hammad")
// 	fmt.Println("Office Friends")
// 	users("faizan","vikash","rishu","janvi","abhishek")
// 	// fmt.Println(add(1,2,3))
// 	// fmt.Println(add(1,2,3,4,5,6))
// 	// fmt.Println(add(1,2,3,4,5,6,7,8))
// 	// fmt.Println(add(6,7,8))
// }

// import (
// 	"bytes"
// 	"fmt"
// 	"sort"
// )

// func main() {
// 	fmt.Println("Bytes")
// 	data := []byte("a,b,c,d,e,f,g")
// 	fmt.Println(string(data))
// 	parts := bytes.Split(data, []byte(","))
// 	for _, part := range parts{
// 		fmt.Println(string(part))
// 	}

// 	slice_1 := []byte{'@','#','F','A','I','Z','A','N','#','@'}
// 	fmt.Printf("Slice : %s\n",slice_1)
// 	res1 := bytes.Trim(slice_1,"@#")
// 	fmt.Printf("Res : %s\n",res1)

// 	res2 := bytes.Trim([]byte{'$','%','#','s','a','l','m','a','n','#','%','$'},"$,%,#")
// 	fmt.Printf("res2 : %s\n",res2)
// 	slice_2 := []byte("         hello ,world           ")
// 	fmt.Println("Slice 2 : ",string(slice_2))
// 	res3 := bytes.TrimSpace(slice_2)
// 	fmt.Printf("res3 :  %s\n",res3)

// fmt.Println("Sorting Slices")
// intSlice := []int{1,3,5,7,9,2,4,6,8}
// sort.Slice(intSlice,func(i, j int) bool {
// 	return intSlice[i] > intSlice[j]
// });
// fmt.Println("Decending Order : ",intSlice)
// sort.Ints(intSlice)
// fmt.Println("Assending Order : ",intSlice)
// fmt.Println("Before : ",intSlice)
// fmt.Println("Is the Array is Sorted : ",sort.IntsAreSorted(intSlice))
// sort.Ints(intSlice)
// fmt.Println("After : ",intSlice)
// fmt.Println("Is the Array is Sorted : ",sort.IntsAreSorted(intSlice))
// }

// import "fmt"

// func main() {
// 	fmt.Println("Slices")
// 	var arr = []int{1,2,3,4,5,6,7,8,9}
// 	fmt.Println(arr)
// 	arr1 := arr[2:7]
// 	fmt.Println(arr1)
// 	arr2 := arr[1:]
// 	fmt.Println(arr2)
// }

// Structure
// import (
// 	"fmt"
// )

// type Teacher struct{
// 	name string
// 	age int
// 	subject string
// 	exp int
// 	student Student
// }

// type Student struct{
// 	name string
// 	subject string
// 	class int
// }

// func main(){
// 	fmt.Println("Nested Structure Occured!")
// 	Teacher1 := Teacher{
// 		name : "Suman Jha",
// 		age :34,
// 		subject : "PHP",
// 		exp : 14,
// 		student: Student{"rishu","PHP",13},
// 	}
// 	fmt.Println(Teacher1)
// 	Teacher2 := Teacher{
// 		student: Student{"vikash","php",15},
// 	}
// 	fmt.Println(Teacher2)
// }

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
