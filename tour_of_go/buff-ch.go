package main

import "fmt"

func main() {
	ch := make(chan int, 2)
	ch <- 1
	ch <- 2

	// With the below uncommented we get a deadlock :cryalot:
	// ch <- 3
	// Same happens if we try using an unbuffered channel w/o goros

	fmt.Println(<-ch)
	fmt.Println(<-ch)
}
