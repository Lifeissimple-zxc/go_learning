package story

import (
	"fmt"
	"net/http"
	"strings"
)

type StoryRouter struct {
	St Story
}

func (sr StoryRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// sr is not actually needed?
	// Here we handle different paths of incoming requests
	// path handling can be a separate function! TODO
	cleanPath := strings.ToLower(r.URL.Path)
	// Defaulting index path to intro
	if cleanPath == "/" {
		cleanPath = "/intro"
	}
	fmt.Println("Got an inbond for", cleanPath)
	fmt.Println("Paths supported", cleanPath)
	for key := range sr.St.Arcs {
		fmt.Println(key)
	}

	arc, ok := sr.St.Arcs[cleanPath[1:]]
	fmt.Printf("%s lookup result: %v, %#v", cleanPath, ok, arc)
	if !ok {
		// Quick fail for unexpected arcs
		http.Error(w, "Page not found", http.StatusNotFound)
	} else {
		// Render using template saved within story
		// This should use template saved within self!
		arc.RenderHTML(w)
	}
}
