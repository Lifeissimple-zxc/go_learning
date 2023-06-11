package main

import (
	"fmt"
	"io"
	"net/http"
	"os"
)

// Custom writer interface
type logWriter struct{}

func main() {
	// Making an http request
	resp, err := http.Get("http://google.com")
	// Handle error
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Naive way of printing our response from the server
	// fmt.Println(resp)

	// Proper logging of the response body to the command line
	// We set a size of our slice because Read does not resize our slice
	// bs := make([]byte, 99999) // give us an empty byte slice with 99999 elements

	// resp.Body.Read(bs) // This parses our body
	// // Log our HTML response
	// fmt.Println(string(bs))

	// The correct way of logging response to the console
	lw := logWriter{}
	io.Copy(lw, resp.Body) // Works 99999X faster!
}

// Making our logWriter match the Writer interface definition
func (logWriter) Write(bs []byte) (n int, err error) {
	fmt.Println(string(bs))
	bsLen := len(bs)
	fmt.Println("Just wrote this many bytes:", bsLen)

	return bsLen, nil
}
