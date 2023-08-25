package main

import (
	"cyoa/story"
	"encoding/json"
	"fmt"
	"html/template"
	"io"
	"net/http"
	"os"
)

func main() {

	// open file
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

	// Generate HTML
	htmlBytes, err := readFile("templates/arc.html")
	if err != nil {
		fmt.Printf("Err opening html file. Details: %v\n", err)
		os.Exit(1)
	}
	// parse to Story struct
	var st story.Story
	st.Init(storyData, htmlBytes)

	mux := http.NewServeMux()
	tmpl := template.Must(template.ParseFiles("templates/arc.html"))
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))
	mux.HandleFunc("/todo", todo)

}

func showStory(name string, st *story.Story, tmpl *template.Template)

// readFile reads a file to a bytes slice
func readFile(path string) ([]byte, error) {
	f, err := os.Open(path)
	if err != nil {
		return nil, err
	}
	defer f.Close()

	return io.ReadAll(f)
}
