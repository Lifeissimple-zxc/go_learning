package main

import (
	"cyoa/story"
	"encoding/json"
	"fmt"
	"io"
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

	// parse to Story struct
	var story story.Story
	story.Init(storyData)

	fmt.Println("Story has been parsed")
	fmt.Println("###################")
	// fmt.Printf("Story data: %#v", story.Arcs["intro"])

	// Generate HTML
	htmlBytes, err := readFile("templates/arc.html")
	if err != nil {
		fmt.Printf("Err opening html file. Details: %v\n", err)
		os.Exit(1)
	}

	err = story.Arcs["intro"].RenderHTML(htmlBytes, "intro", os.Stdout)
	if err != nil {
		fmt.Printf("Err rendering html file. Details: %v\n", err)
		os.Exit(1)
	}

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
