package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time and Date in Go lang")

	currentTime := time.Now();
	fmt.Println("Current Time : ",currentTime)
	fmt.Printf("Type of time package %T\n",currentTime)

	// formatted := currentTime.Format("dd-mm-yyyy")   // its wrong way
	// formatted := currentTime.Format("02-01-2006, Day")   // its correct way
	// formatted := currentTime.Format("02-01-2006, Monday")   // if uh need day also
	// formatted := currentTime.Format("02-01-2006, 15:04:05")   // if uh need 24 hour format
	formatted := currentTime.Format("02/01/2006, 3:04 PM")   // if uh need 12 hour format
	fmt.Println("Formatted date :",formatted)


	layout_str := "2006-01-02"
	dateStr := "2023-11-25"
	formatted_time, _ := time.Parse(layout_str, dateStr)
	fmt.Println("Formatted Time : ",formatted_time)

	new_date := currentTime.Add(24 * time.Hour)
	fmt.Println("New Date time is : ",new_date)
	formatted_new_date := new_date.Format("2006/01/02 Monday")
	fmt.Println("Formatted new Date : ",formatted_new_date)
}