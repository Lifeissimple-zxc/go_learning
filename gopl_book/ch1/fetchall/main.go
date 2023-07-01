package main

import (
	"fmt"
	"io"
	"io/ioutil"
	"net/http"
	"os"
	"strings"
	"time"
)

const logFileName string = "response_logs.txt"

func main() {
	// Check for arg len
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "Usage: programm.go url1 url2 ... urln")
		os.Exit(1)
	}
	start := time.Now()
	// Since it is a concurrent implementation, we are relying on channels
	ch := make(chan string)

	for _, url := range os.Args[1:] {
		go fetch(url, ch) // Start a goro
	}

	// Open a file, assuming no issues for the sake of exercise
	// A bit cryptic, in udemy tutorial we wrote bytes
	outFile, _ := os.OpenFile(logFileName, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0644)

	for range os.Args[1:] {
		// fmt.Println(<-ch) // Read from a channel
		// 1.10 Modify the program to send channel data to a file
		// Don't see much uplift from caching, but the websites I tested against should have it
		fmt.Fprintf(outFile, <-ch+"\n")
	}

	// Log runtime
	fmt.Fprintf(outFile, "Exec time: %.2f\n", time.Since(start).Seconds())

	// Close file, no check for errors just coz
	outFile.Close()

}

// Helper via which we actually GET data from URLs
func fetch(url string, ch chan<- string) {
	start := time.Now()
	// 1.8  Make sure we add https:// to all the urls
	urlBase := "https://"
	if !(strings.Contains(url, urlBase)) {
		url = urlBase + url
	}

	resp, err := http.Get(url) // Makes a get request
	if err != nil {
		// Inform on error & exit (it different from fetch due to concurrency)
		// We just convert the err to string an send it to channel
		ch <- fmt.Sprintf("Got an error for url %s. Details: %v\n", url, err)
		return
	}

	nbytes, err := io.Copy(ioutil.Discard, resp.Body)
	resp.Body.Close() // Avoid leaking.
	if err != nil {
		ch <- fmt.Sprintf("Error reading resp body for %s. Details: %v\n", url, err)
	}

	secs := time.Since(start).Seconds()
	// Write url processing outcome to the channel
	ch <- fmt.Sprintf("%.2fs %7d %s", secs, nbytes, url)

}
