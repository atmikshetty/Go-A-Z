package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {

	//	Goroutines allow for Concurrency within our application.
	//	Could be a method or function, it runs concurrently with other functions
	//	Example: When you open Zomato or Swiggy, multiple things are loaded at once, the discount tab, nearest restauranets, offers , etc. This is an example of Go routines.

	// Created using the "go" keyword

	var wg sync.WaitGroup // creating a waitgroup
	wg.Add(1)             // adding a goroutine to it

	go printNos(&wg)
	fmt.Println("Executing a Goroutine concurrently!!!")

	time.Sleep(5 * time.Second)
	fmt.Println("Done with the main function!!!")

	// Synchronizing Go Routines
	//	This is to make sure that different go routines do not modify the same data
	//	sync package and WaitGroups are used for this purpose

	//	Communication between Goroutines using Channels
	//ch := make(chan int) // creating a channel

	//go sendNos(ch)

	//	Receiving nos from the Channel
	//for i := range ch {
	//	fmt.Println(i)
	//}
	//fmt.Println("Main function executed!!!")

	//	Using Select with Goroutines
	ch := make(chan int)
	done := make(chan bool)

	go sendNos(ch)

	// receiving numbers
	go func() {
		for i := range ch {
			fmt.Println(i)
		}
		done <- true
	}()

	//	Using select to verify the channel
	select {
	case <-done:
		fmt.Println("All numbers recieved!!!")
	case <-time.After(1 * time.Second):
		fmt.Println("Request Timed out !!!")
	}
}

func sendNos(ch chan int) {
	for i := 1; i <= 5; i++ {
		ch <- i // send nos to the channel
		time.Sleep(1 * time.Second)
	}
	close(ch) // closing the channel
}

func printNos(wg *sync.WaitGroup) {
	defer wg.Done() // tells that the routine is done executing
	for i := 1; i <= 100; i++ {
		fmt.Println(i)
		//	Simulating some work
		time.Sleep(1 * time.Second)
	}
}
