// ############################### Part 3 - Done #############################
package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main(){
	fmt.Println("Web Request in Go Lang")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/2");
	if err != nil{
		fmt.Println("Failed to Get the Data")
		return
	}
	defer res.Body.Close()
	// fmt.Println(res.Body)
	data, err1 := ioutil.ReadAll(res.Body)
	if err1 != nil{
		fmt.Println("Failed to Get the Data")
		return
	}
	fmt.Println("Data : ",string(data))

}
// ############################### Part 2 - Done #############################
// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// func main(){
// 	file,err := os.Open("testing.txt")
// 	if err != nil{
// 		fmt.Println("Error While Reading the File")
// 		return
// 	}
// 	defer file.Close()

// 	buffer := make([]byte,1024)

// 	for{
// 		n,err := file.Read(buffer)
// 		if err == io.EOF{
// 			break
// 		}
// 		if err != nil{
// 			fmt.Println("Error While Readin")
// 			return
// 		}
// 		fmt.Println(string(buffer[:n]))
// 	}
// }
// ############################### Part 1 - Done #############################
// package main

// import (
// 	"fmt"
// 	"io"
// 	"os"
// )

// func main(){
// 	fmt.Println("File Handling in Go Lang")
// 	data, err := os.Create("testing.txt")

// 	if err != nil{
// 		fmt.Println("Error Occured While Creating a File : ",err)
// 		return
// 	}
// 	defer data.Close()

// 	content := "package testing \n import \"fmt\" \n func testing(){\n fmt.Println(\"Hello World\") \n }"
// 	_, errors := io.WriteString(data,content)
// 	if errors != nil{
// 		fmt.Println("Error While Writing the File : ",errors)
// 		return
// 	}
// 	fmt.Println("Content Entered")
// }