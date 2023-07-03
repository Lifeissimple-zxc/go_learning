// This server is a minimal echo server
package main

import (
	"fmt"
	"log"
	"net/http"
	"os"
	"strings"
)

const PORT = "8080"

func main() {

	if len(os.Args) == 1 {
		fmt.Fprintf(os.Stderr, "Usage: programm.go r | l\n")
		os.Exit(1)
	}

	env := strings.ToLower(os.Args[1])

	if !(env == "r" || env == "l") {
		fmt.Fprintf(os.Stderr, "Unexpected env: %s\n", env)
		os.Exit(1)
	}

	var servers = map[string]string{
		"r": "",
		"l": "localhost",
	}

	server := servers[env] + ":" + PORT

	http.HandleFunc("/", handler) // each request calls our handler func defined below
	log.Fatal(http.ListenAndServe(server, nil))
}

// Handler echoes the path component ofm the request URL
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
