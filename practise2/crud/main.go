package main

import (
	"encoding/json"
	"fmt"

	// "io/ioutil"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func main() {
	fmt.Println("Crud Operation in Crud")
	// res,err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	res, err := http.Get("https://jsonplaceholder.typicode.com/todos/1")
	if err != nil {
		fmt.Println("Error Getting Get Response : ", err)
		return
	}

	defer res.Body.Close()

	if res.StatusCode != http.StatusOK {
		fmt.Println("Error in getting response ; ", res.Status)
		return
	}

	// data, err := ioutil.ReadAll(res.Body)
	// if err != nil {
	// 	fmt.Println("Error while Reading the File")
	// 	return
	// }

	// fmt.Println("Data  : ", string(data))

	var todo Todo
	err = json.NewDecoder(res.Body).Decode(&todo)
	if err != nil {
		fmt.Println("Error Decoding : ", err)
		return
	}
	fmt.Println("Todo : ", todo)

}
