package main

import (
	"fmt"
	"net/url"
)

func main() {
	fmt.Println("Url In Go Lang")
	Link := "https://example.com:8080/path/to/resource?key1=value&key2=value2"
	// fmt.Println("Link : ",Link)
	// fmt.Printf("The Type of Link is : %T\n ",Link)

	parseLink,err := url.Parse(Link)
	if err !=nil{
		fmt.Println("Error While Parsing the Link")
		return
	}
	// fmt.Println("Parsed Link : ",parseLink)
	// fmt.Printf("Type of link is : %T\n",parseLink )
	fmt.Println("Scheme : ",parseLink.Scheme)
	fmt.Println("Host : ",parseLink.Host)
	fmt.Println("Path : ",parseLink.Path)
	fmt.Println("RawQuery : ",parseLink.RawQuery)

	// Modifhying Url Components
	parseLink.Path = "/newPath/changed"
	parseLink.RawQuery = "username=faizan"

	// Constructing a Url string from a URL Object
	newurl := parseLink.String()

	fmt.Println("New Url : ",newurl)

}