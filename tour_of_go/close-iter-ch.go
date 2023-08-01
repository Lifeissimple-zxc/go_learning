package main

import "fmt"

func fib(n int, ch chan int) {
	x, y := 0, 1
	for i := 0; i < n; i++ {
		ch <- x
		x, y = y, x+y
	}
	close(ch)

}

func main() {
	ch := make(chan int, 10) // buffered channel (queue) with a limit of 10
	go fib(cap(ch), ch)      // cap here returns the capacity of the channel which we use as input to fib
	// Since fub() closes the channel, we can safely iterate here using range
	for i := range ch {
		fmt.Println(i)
	}
}
