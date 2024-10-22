package main

import (
	"fmt"
	"sync"
)

// Sum of the numbers in the given range
func count(start, end int, wg *sync.WaitGroup, ch chan int) {
	defer wg.Done()
	sum := 0
	for i := start; i < end; i++ {
		sum += i
	}
	//	sending the result to a channel
	ch <- sum
}

func main() {
	n := 10 // no of goroutines
	totalNums := 100
	step := totalNums / n

	//	channel to communicate between the goroutines
	ch := make(chan int, n)
	var wg sync.WaitGroup

	for i := 0; i < n; i++ {
		wg.Add(1)
		//calculate the start and end for each goroutine
		start := i*step + 1
		end := (i + 1) * step
		go count(start, end, &wg, ch)
	}

	// Close the channel when all goroutines are done
	go func() {
		wg.Wait()
		close(ch)
	}()

	totalSum := 0
	for sum := range ch {
		totalSum += sum
	}

	fmt.Println("Total Sum:", totalSum)

}
