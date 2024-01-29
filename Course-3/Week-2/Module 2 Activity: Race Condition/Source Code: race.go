package main

import (
	"fmt"
	"sync"
	"time"
)

func main() {
	//initialize the number variable with 0
	number := 0
	// wait is used for waiting for a collection of goroutines to finish
	var wait sync.WaitGroup

	//iteration for 2 goroutines
	for i := 0; i < 2; i++ {
		//added to the waitgroup counter
		wait.Add(1)
		go func() {
			for j := 0; j < 400; j++ {
				/*--------------
				- increasing the value number by 1
				- interleaving is totally random
				- read and set operation can't be predicted, the sequence in different goroutines
				---------------*/
				number++
				time.Sleep(1 * time.Millisecond)
			}
			wait.Done()
		}()
	}
	//waiting for 2 goroutines finish
	wait.Wait()

	/*---------------
	- the result could be random number but less than or equal to 800
	- if there is no race condition between 2 goroutines the value will be precisely 800
	- 2 goroutines increase separately with 400 times
	---------------*/
	fmt.Println("Finally number becomes:", number)
}

// Race Condition occurs in Go when two or more goroutines have shared data and interact with it at the same time.
// This leads to the outcome of the code to depend on non-deterministic ordering.
