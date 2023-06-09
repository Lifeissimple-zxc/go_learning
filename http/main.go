package main

import (
	"fmt"
	"net/http"
	"os"
)

func main() {
	// Making an http request
	resp, err := http.Get("http://google.com")
	// Handle error
	if err != nil {
		fmt.Println("Error:", err)
		os.Exit(1)
	}

	// Naive way of printing our response from the server
	fmt.Println(resp)
}
