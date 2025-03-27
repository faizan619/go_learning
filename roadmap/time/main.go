package main

import (
	"fmt"
	"time"
)

func main() {
	fmt.Println("Time Package in Go Lang")

	time1 := time.Now()
	fmt.Println("Current Date : ",time.Now().Weekday(),"-",time.Now().Day(),"-",time.Now().Month(),"-",time.Now().Year())
	fmt.Println("Tomorrow Date : ",time.Now().Weekday()+1,"-",time.Now().Day()+1,"-",time.Now().Month(),"-",time.Now().Year())
	currMin := time1.Minute()
	fmt.Println("Current Minute : ",currMin)
	currMinAdd5 := time1.Add(time.Minute * 5)
	fmt.Println("Add 5 minute in Current time : ",currMinAdd5.Minute())

	// fmt.Println(5*time.Hour + 30*time.Minute + 10000*time.Millisecond)

	// h,_ := time.ParseDuration("5h30m8s")
	// fmt.Println("I have got ",h.Hours(),"Hours for completing my project")
	// fmt.Println("I have got ",h.Minutes(),"Minutes for completing my project")
	// fmt.Println("I have got ",h.Seconds(),"Seconds for completing my project")
	// fmt.Println("I have got ",h.Milliseconds(),"Milliseconds for completing my project")
	
	// var count int
	// counter1 := time.Now();
	// timer := time.Tick(10 * time.Millisecond)
	// for next := range timer {
	// 	fmt.Println(next)
	// 	count++
	// 	if count == 100 {
	// 		break
	// 	}
	// }
	// counter2 := time.Since(counter1);
	// fmt.Println("THe Loop Took ",counter2," to run the program")
	
	
	
	// var count int
	// counter1 := time.Now();
	// timer := time.Tick(100 * time.Millisecond)
	// for next := range timer {
	// 	fmt.Println(next)
	// 	count++
	// 	if count == 100 {
	// 		break
	// 	}
	// }
	// counter2 := time.Now();
	// fmt.Println("THe Loop Took ",counter2.Sub(counter1)," to run the program")



	// fmt.Println("Start Timing ...")
	// time.Sleep(3000 * time.Millisecond)
	// fmt.Println("Stop Timing ...")

	// start := time.Now();
	// fmt.Println("Current Time : ",start)
}

// const (
// 	Layout      = "01/02 03:04:05PM '06 -0700" // The reference time, in numerical order.
// 	ANSIC       = "Mon Jan _2 15:04:05 2006"
// 	UnixDate    = "Mon Jan _2 15:04:05 MST 2006"
// 	RubyDate    = "Mon Jan 02 15:04:05 -0700 2006"
// 	RFC822      = "02 Jan 06 15:04 MST"
// 	RFC822Z     = "02 Jan 06 15:04 -0700" // RFC822 with numeric zone
// 	RFC850      = "Monday, 02-Jan-06 15:04:05 MST"
// 	RFC1123     = "Mon, 02 Jan 2006 15:04:05 MST"
// 	RFC1123Z    = "Mon, 02 Jan 2006 15:04:05 -0700" // RFC1123 with numeric zone
// 	RFC3339     = "2006-01-02T15:04:05Z07:00"
// 	RFC3339Nano = "2006-01-02T15:04:05.999999999Z07:00"
// 	Kitchen     = "3:04PM"
// 	// Handy time stamps.
// 	Stamp      = "Jan _2 15:04:05"
// 	StampMilli = "Jan _2 15:04:05.000"
// 	StampMicro = "Jan _2 15:04:05.000000"
// 	StampNano  = "Jan _2 15:04:05.000000000"
// 	DateTime   = "2006-01-02 15:04:05"
// 	DateOnly   = "2006-01-02"
// 	TimeOnly   = "15:04:05"
// )
