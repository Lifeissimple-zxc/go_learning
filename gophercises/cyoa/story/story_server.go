package story

import (
	"net/http"
	"strings"
)

type StoryRouter struct {
	St Story
}

func (sr StoryRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// sr is not actually needed?
	// Here we handle different paths of incoming requests
	cleanPath := strings.ToLower(r.URL.Path)
	// Defaulting index path to intro
	if cleanPath == "/" {
		cleanPath = "intro"
	}
	if arc, ok := sr.St.Arcs[cleanPath]; !ok {
		// Quick fail for unexpected arcs
		http.Error(w, "Page not found", http.StatusNotFound)
	} else {
		// Render using template saved within story
		// This should use template saved within self!
		arc.RenderHTML(w)
	}
}
