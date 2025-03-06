package main

import (
	"fmt"
	"io/ioutil"
	"net/http"
)

func main() {
	fmt.Println("Web Request in Go Lang")
	res,err := http.Get("https://jsonplaceholder.typicode.com/todos/1")

	if err != nil {
		fmt.Println("Error Getting Get Response : ",err)
		return 
	}

	defer res.Body.Close()
	fmt.Printf("Type of Response is : %T\n",res)

	data, err := ioutil.ReadAll(res.Body)
	if err != nil{
		fmt.Println("Error while Reading the File")
		return 
	}

	fmt.Println("Response : ",string(data))
}