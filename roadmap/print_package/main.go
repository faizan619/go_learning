// file handling
package main

import (
	"fmt"
	"os"
)

func main() {
	fmt.Println("Reading File in Go lang")
	file, err := os.Open("example1.txt")
	if err != nil {
		fmt.Println("Error While Reading the File : ", err)
		return
	}
	// fmt.Println("Data : ", string(file))

	data1 := "deltasoft garden"
	_, err1 := file.WriteString(data1)
	if err1 != nil {
		fmt.Println("Error While Reading the File : ", err1)
		return
	}
	fmt.Println("Done With Writing the Content")
}

// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// func main() {
// 	fmt.Println("File Handling in Go Lang")

// 	file, err := os.Create("example1.txt")
// 	if err != nil{
// 		fmt.Println("Error while creating the file : ",err)
// 		return
// 	}

// 	defer file.Close()

// 	content := "This is Faizan Alam of DeltaSoft Systems";
// 	_, err1 := io.WriteString(file,content+"\n")
// 	if err1 != nil{
// 		fmt.Println("Error while Writing the file : ",err)
// 		return
// 	}
// 	fmt.Println("Content is Written Successfully")
// }

// package main

// import "fmt"

// func main() {

// 	// const name,id = "faizan alam",21
// 	// err := fmt.Errorf("User %q (id %d) not found",name,id)
// 	// fmt.Println(err)

// 	fmt.Println(" # This is normal Println # ")

// 	name := "faizan alam"
// 	age := 22
// 	isAdult := true
// 	marks := 65.3465
// 	// fmt.Printf("The name of User is : %[2]s and age is %[1]d \n",age,name)

// 	// fmt.Printf("The Name is : %s and age is %d \n",name,age)

// 	// user := fmt.Sprintf("The name is %s and age is %d \n",name,age)
// 	// fmt.Println(user)

// 	fmt.Printf("The name Type : %T \n",name);
// 	fmt.Printf("The name Value : %v \n",name);
// 	fmt.Printf("Is Faizan Adult : %t \n",isAdult);
// 	fmt.Printf("The age is : %d \n",age)
// 	fmt.Printf("The Marks Type : %T \n",marks)
// 	fmt.Printf("The Marks Value : %f \n",marks)
// 	fmt.Printf("The Marks Value[2] : %.2f \n",marks)

// }
