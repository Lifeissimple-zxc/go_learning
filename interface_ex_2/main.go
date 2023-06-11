package main

import (
	"fmt"
	"io"
	"os"
)

// This needs to
// Read filename from CLI args
// Print contents of the file to stdout
func main() {
	// Read filename from CLI args
	filename := os.Args[1]
	// Open a file
	file, err := os.Open(filename)
	// Exit if error for whatever reason
	if err != nil {
		fmt.Println("Failed to open file. Details:", err)
		os.Exit(1)
	}
	// Since Open returns a file, it has matches both Reader and Writer interfaces
	// Therefore, we can call io.Copy with destination being os.Stdout and src being our file
	// Print its contents to stdout
	io.Copy(os.Stdout, file)
	// Close our file
	clErr := file.Close()
	if clErr != nil {
		fmt.Println("Error when closing file. Details:", clErr)
		os.Exit(1)
	}

}
