package main

import (
	"fmt"
	// "io"
	"io/ioutil"
	// "os"
)

func main(){
	
	// Create file with Data insert
	/*
	fmt.Println("File Handling in Go!");

	file, err := os.Create("example.txt")
	if err != nil {
		fmt.Println("Error Creating file : ",err)
		return
	}

	defer file.Close();		// file is created
	
	initialCount := "Hello, this is the middle Content"
	// _, errors := io.WriteString(file, initialCount+"\n")
	bytes, errors := io.WriteString(file, initialCount+"\n")

	if errors != nil {
		fmt.Println("Error While Writing Files : ",errors);
		return
	}

	fmt.Println("Bytes return is : ",bytes)

	fmt.Println("file Is Created")
	*/

	// Open a File and Read Conotent
	/*
	file, err := os.Open("example.txt")
	if err != nil {
		fmt.Println("Error while Opening the File :",err)
		return
	}
	defer file.Close()

	// create a buffer [a temporary storage] to read the file content
	buffer := make([]byte, 1024)

	// read the file content into the buffer
	for {
		n, err := file.Read(buffer)
		if err == io.EOF{
			break
		}
		if err != nil{
			fmt.Println("Error while reading file :",err)
			return
		}
		fmt.Println(string(buffer[:n]))
	}
	*/

	// Read the Entire File into a slice - shortcut
	content ,err := ioutil.ReadFile("example.txt")

	if err != nil {
		fmt.Println("Error while reading the File : ",err)
		return 
	}
	fmt.Println(string(content))

}