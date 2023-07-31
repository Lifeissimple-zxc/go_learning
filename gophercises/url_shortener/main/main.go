package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"
	"strings"

	urlshort "url_shortener"
)

func main() {

	// CLI args parsing
	fPtr := flag.String("file", "paths.yaml", "Path to a yaml of [{path: url: }, {..}] form")
	flag.Parse()

	// We read mapping file before anything
	file, err := readYAML(*fPtr)
	if err != nil {
		fmt.Printf("Error when opening file: %v\n", err)
		os.Exit(1)
	}

	mux := defaultMux()

	// Build the MapHandler using the mux as the fallback
	pathsToUrls := map[string]string{
		"/urlshort-godoc": "https://godoc.org/github.com/gophercises/urlshort",
		"/yaml-godoc":     "https://godoc.org/gopkg.in/yaml.v2",
	}
	mapHandler := urlshort.MapHandler(pathsToUrls, mux)

	// Dynamically decide on handler depending on the type of file we get
	var myHandler http.HandlerFunc
	// Declaring e here to avoid writing if e != nil two times
	var e error
	if strings.Contains(*fPtr, "json") {
		myHandler, e = urlshort.JSONHandler([]byte(file), mapHandler)
	} else {
		myHandler, e = urlshort.YAMLHandler([]byte(file), mapHandler)
	}
	if e != nil {
		panic(e)
	}

	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", myHandler)
}

func defaultMux() *http.ServeMux {
	mux := http.NewServeMux()
	mux.HandleFunc("/", hello)
	return mux
}

func hello(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Hello, world!")
}

func readYAML(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return io.ReadAll(f)
}
