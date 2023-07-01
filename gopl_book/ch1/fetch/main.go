// Fetch prints content found at a given URL
package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"
)

func main() {
	// Check for arg len
	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "Usage: programm.go url1 url2 ... urln")
		os.Exit(1)
	}

	for _, url := range os.Args[1:] {
		// 1.8  Make sure we add https:// to all the urls
		urlBase := "https://"
		if !(strings.Contains(url, urlBase)) {
			url = urlBase + url
		}
		resp, err := http.Get(url) // Makes a get request
		if err != nil {
			// Inform on error & exit
			fmt.Fprintf(os.Stderr, "Got an error for url %s. Details: %v", url, err)
			os.Exit(1)
		}
		// 1.9 Print resp status code
		fmt.Printf("Got %d status code for %s\n", resp.StatusCode, url)
		// Parse body
		// b, err := ioutil.ReadAll(resp.Body)
		// 1.7 Read response body using io.Copy() instead of ReadAll to optimize memory allocation
		_, err = io.Copy(os.Stdout, resp.Body)
		if err != nil {
			// Inform on error & exit
			fmt.Fprintf(os.Stderr, "Error reading resp body for %s. Details: %v", url, err)
			os.Exit(1)
		}
		resp.Body.Close() // Needed to avoid resource leaks, can be used with defer
		// Test comment
	}
}
