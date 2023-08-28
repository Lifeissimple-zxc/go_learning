package story

import (
	"fmt"
	"net/http"
	"strings"
)

type StoryRouter struct {
	St Story
}

func (sr *StoryRouter) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	// Here we handle different paths of incoming requests
	// path handling can be a separate function! TODO
	cleanPath := strings.ToLower(r.URL.Path)
	// Defaulting index path to intro
	if cleanPath == "/" {
		cleanPath = "/intro"
	}
	fmt.Println("Got an inbond for", cleanPath)

	arc, ok := sr.St.Arcs[cleanPath[1:]]
	// 404 here causes css file to be MIME, fix this
	if !ok {
		// Quick fail for unexpected arcs
		fmt.Printf("%s is not supported", r.URL.Path)
		return
	}
	// Render using template saved within story
	// This should use template saved within self!
	arc.RenderHTML(w)
}

// MapToMUX registers handling of routes with mux
func (sr *StoryRouter) MapToMUX(mux *http.ServeMux) {
	for arc := range sr.St.Arcs {
		mux.Handle(fmt.Sprintf("/%s", arc), sr)
	}
	mux.Handle("/", sr)
}
