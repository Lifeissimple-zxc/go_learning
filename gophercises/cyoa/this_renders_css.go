package main

import (
	"cyoa/story"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"log"
	"net/http"
	"os"
)

var tmpl *template.Template

func main() {

	mux := http.NewServeMux()
	tmpl = template.Must(template.ParseFiles("templates/arc.html"))

	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))
	mux.HandleFunc("/intro", showStory)

	log.Fatal(http.ListenAndServe(":9091", mux))

}

func showStory(w http.ResponseWriter, r *http.Request) {
	// Generate HTML
	htmlBytes, err := readFile("templates/arc.html")
	if err != nil {
		fmt.Printf("Err opening html file. Details: %v\n", err)
		os.Exit(1)
	}

	file, err := readFile(story.STORY_PATH)
	if err != nil {
		fmt.Printf("Err opening file %s. Details: %v\n", story.STORY_PATH, err)
		os.Exit(1)
	}

	// parse to storyData
	var storyData map[string]json.RawMessage
	err = json.Unmarshal(file, &storyData)
	if err != nil {
		fmt.Printf("Err parsing to json. Details: %v\n", err)
		os.Exit(1)
	}

	var st story.Story
	st.Init(storyData, htmlBytes)

	// Execute template
	tmpl.Execute(w, st.Arcs["intro"])
}

// readFile reads a file to a bytes slice
func readFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return io.ReadAll(f)
}
