package main

import (
	"flag"
	"fmt"
	"io"
	"net/http"
	"os"

	urlshort "url_shortener"
)

func main() {

	// CLI args parsing
	ymlPtr := flag.String("yaml", "paths.yaml", "Path to a yaml of [{path: url: }, {..}] form")
	flag.Parse()

	// We read yaml before anything
	yaml, err := readYAML(*ymlPtr)
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

	// Build the YAMLHandler using the mapHandler as the
	// fallback
	yamlHandler, err := urlshort.YAMLHandler([]byte(yaml), mapHandler)
	if err != nil {
		panic(err)
	}
	fmt.Println("Starting the server on :8080")
	http.ListenAndServe(":8080", yamlHandler)
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
