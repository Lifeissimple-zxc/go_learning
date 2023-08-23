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

	fmt.Println("Story has been parsed")
	fmt.Println("###################")

	router := story.StoryRouter{St: st}

	fmt.Println("Router instantiated")

	// The below line is needed bc CSS and other static content
	// is server separately from HTML
	http.Handle("/static/", http.StripPrefix("/static/", http.FileServer(http.Dir("static"))))

	log.Fatal(http.ListenAndServe(":5070", router))

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
