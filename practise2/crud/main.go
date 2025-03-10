package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"strings"

	// "io/ioutil"
	"net/http"
)

type Todo struct {
	UserId    int    `json:"userId"`
	Id        int    `json:"id"`
	Title     string `json:"title"`
	Completed bool   `json:"completed"`
}

func performGetRequest() {
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

func performPostRequest() {
	todo := Todo{
		UserId:    23,
		Title:     "Faizan MAster",
		Completed: true,
	}

	// Convert the Todo Struct to Json
	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error Marshaling : ", err)
		return
	}

	// Convert the Json Data to String
	jsonString := string(jsonData)

	// Convert String Data to Reader
	jsonReader := strings.NewReader(jsonString)

	myUrl := "https://jsonplaceholder.typicode.com/todos"
	res, err1 := http.Post(myUrl, "application/json", jsonReader)
	if err1 != nil {
		fmt.Println("Error Sending Post : ", err1)
		return
	}
	defer res.Body.Close()

	data, err2 := ioutil.ReadAll(res.Body)
	if err2 != nil {
		fmt.Println("Error Sending Post : ", err2)
		return
	}
	fmt.Println("Response : ", string(data))

	fmt.Println("Response Data : ", res.Status)

}

func performUpdateRequest() {
	todo := Todo{
		UserId:    1,
		Id:        1,
		Title:     "Update Todo",
		Completed: false,
	}

	jsonData, err := json.Marshal(todo)
	if err != nil {
		fmt.Println("Error Marshaling JSON : ", err)
		return
	}

	jsonString := string(jsonData)

	// Convert String Data to Reader
	jsonReader := strings.NewReader(jsonString)

	const myUrl = "https://jsonplaceholder.typicode.com/todos/1"

	// create PUT Request
	req, err1 := http.NewRequest(http.MethodPut, myUrl, jsonReader)
	if err1 != nil {
		fmt.Println("Error Marshaling JSON : ", err1)
		return
	}
	req.Header.Set("Content-type", "application/json")

	// Sent the Request
	client := http.Client{}
	res, err2 := client.Do(req)
	if err2 != nil {
		fmt.Println("Error  : ", err2)
		return
	}
	defer res.Body.Close()

	data, err3 := ioutil.ReadAll(res.Body)
	if err3 != nil {
		fmt.Println("Error Sending Post : ", err3)
		return
	}
	fmt.Println("Response : ", string(data))

	fmt.Println("Response Data : ", res.Status)

}

func performDeleteRequest(){
	// const myUrl = "https://jsonplaceholder.typicode.com/todos"
	const myUrl = "https://jsonplaceholder.typicode.com/todos/1"

	// Create Delete Request
	req,err := http.NewRequest(http.MethodDelete,myUrl,nil)
	if err != nil {
		fmt.Println("Error Creating Delete Request : ", err)
		return
	}

	// Sent the request
	client := http.Client{}
	resp, err1 := client.Do(req)
	if err1 != nil {
		fmt.Println("Error Marshaling JSON : ", err1)
		return
	}

	defer resp.Body.Close();
	fmt.Println("Response Status : ",resp.Status)
}

func main() {
	fmt.Println("Crud Operation in Crud1")
	// performGetRequest()
	// performPostRequest()
	// performUpdateRequest()
	performDeleteRequest()
}
