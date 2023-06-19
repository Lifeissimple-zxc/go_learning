package main

import (
	"fmt"
	"net/http"
	"time"
)

// Const with our websites
var LINKS = []string{
	"http://google.com",
	"http://amazon.com",
	"http://facebook.com",
	"http://stackoverflow.com",
	"http://golang.org",
}

func main() {

	// Create a channel to communicate with GoRoutines
	c := make(chan string)

	for _, link := range LINKS {
		go checkLink(link, c) // Pass our channel here also
	}

	// Blocking calls to print website status
	/*
		Read data from our channel.
		Blocking call, by default waits on some info to come here.
		Without waiting on the channel, we only get 1 link processed.
	*/
	// fmt.Println(<-c)
	// // Add more blocking calls to print more status logs
	// fmt.Println(<-c)

	// Loop approach to consume messages from the channel
	// While pattern to check for status indefinitely!
	/*
		Wait for channel c to retun a value, then use it as input checkLink
	*/
	for l := range c {
		// fmt.Println(<-c) // blocking call to wait on a message
		// go checkLink(l, c) // redo our request to 'poll' for status
		// Using function literal (ala lambda from .py)
		// func() {} defines the literal, the following () invokes it
		go func(link string) {
			// Adding a pausing call to our programm
			time.Sleep(time.Second * 5)
			// We need to make l an input to the function literal!
			checkLink(link, c)
		}(l)
	}

	fmt.Println("Sleeping between polling attempts!")
}

func checkLink(link string, c chan string) {
	_, err := http.Get(link)
	// c <- sends data to our channel
	if err != nil {
		// Non-nil error means there is an issue
		fmt.Println(link, "might be down :(")
	} else {
		fmt.Println(link, "Tutto Bene")
	}
	// Check for a website is done so we communicate this to our channel.
	c <- link
}
