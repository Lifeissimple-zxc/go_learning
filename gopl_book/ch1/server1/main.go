// This server is a minimal echo server
package main

import (
	"fmt"
	"log"
	"net/http"
)

func main() {
	http.HandleFunc("/", handler) // each request calls our handler func defined below
	log.Fatal(http.ListenAndServe("localhost:8000", nil))
}

// Handler echoes the path component ofm the request URL
func handler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "URL.Path = %q\n", r.URL.Path)
}
