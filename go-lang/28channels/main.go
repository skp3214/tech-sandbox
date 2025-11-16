package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Channel in golang")

	myCh := make(chan int, 1)
	wg := &sync.WaitGroup{}

	// fmt.Println(<-myCh)
	// myCh <- 5

	wg.Add(2)
	// RECEIVE ONLY
	go func(ch <-chan int, wg *sync.WaitGroup) {
		val, isChannelOpen := <-ch
		fmt.Println(isChannelOpen)
		fmt.Println(val)
		// fmt.Println(<-ch)

		wg.Done()
	}(myCh, wg)

	// SEND ONLY
	go func(ch chan<- int, wg *sync.WaitGroup) {
		ch <- 5
		ch <- 6
		defer close(ch)
		wg.Done()
	}(myCh, wg)

	wg.Wait()

}
