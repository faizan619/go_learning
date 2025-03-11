package main

import (
	"fmt"
	"sync"
)

func worker(i int, wg *sync.WaitGroup){
	defer wg.Done()
	fmt.Printf("Worker %d Started\n",i)

	fmt.Printf("Worker %d Ended\n",i)
}

func main() {
	fmt.Println("Explore Goroutine Started!")

	var wg sync.WaitGroup

	for i:=1;i<=5;i++{
		wg.Add(1)	// increment the Waitgroup counter
		go worker(i, &wg)
	}

	wg.Wait()
	fmt.Println("Worker Work Finished")
}