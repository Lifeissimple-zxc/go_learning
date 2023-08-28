package main

import (
	"cyoa/story"
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"os"
)

func main() {

	file, err := readFile(story.STORY_PATH)
	if err != nil {
		fmt.Printf("Err opening file %s. Details: %v\n", story.STORY_PATH, err)
		os.Exit(1)
	}

	var storyData map[string]json.RawMessage
	err = json.Unmarshal(file, &storyData)
	if err != nil {
		fmt.Printf("Err parsing to json. Details: %v\n", err)
		os.Exit(1)
	}

	htmlBytes, err := readFile("templates/arc.html")
	if err != nil {
		fmt.Printf("Err opening html file. Details: %v\n", err)
		os.Exit(1)
	}

	var st story.Story
	st.Init(storyData, htmlBytes)

	fmt.Println("Story has been parsed")
	fmt.Println("###################")

	sr := story.StoryRouter{St: st}

	fmt.Println("Router instantiated")

	mux := http.NewServeMux()
	sr.MapToMUX(mux)
	// This is needed for static content to render
	fs := http.FileServer(http.Dir("./static"))
	mux.Handle("/static/", http.StripPrefix("/static", fs))

	log.Fatal(http.ListenAndServe(":3000", mux))
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
