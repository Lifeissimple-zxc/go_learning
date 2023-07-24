package urlshort

import (
	"fmt"
	"net/http"

	"gopkg.in/yaml.v2"
)

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
	fmt.Println("pathsToUrls:", pathsToUrls)
	// Func that matches http.HandlerFunc signature: https://pkg.go.dev/net/http#HandleFunc
	return func(w http.ResponseWriter, r *http.Request) {
		// Logic: if we match a path --> redirect,
		// else --> fallback
		path := r.URL.Path
		// Oneliner for variable init + control flow
		if dest, ok := pathsToUrls[path]; ok {
			http.Redirect(w, r, dest, http.StatusFound)
			// Exit the func at this point
			return
		}
		fallback.ServeHTTP(w, r)
	}
}

// YAMLHandler will parse the provided YAML and then return
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the YAML, then the
// fallback http.Handler will be called instead.
//
// YAML is expected to be in the format:
//
//   - path: /some-path
//     url: https://www.some-url.com/demo
//
// The only errors that can be returned all related to having
// invalid YAML data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func YAMLHandler(yml []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// Parse the yaml
	var pathUrls []pathUrl
	err := yaml.Unmarshal(yml, &pathUrls)
	if err != nil {
		return nil, err
	}
	// Convert yaml array to a map
	// Create a map-container of our data
	pathsToUrls := make(map[string]string)
	for _, pu := range pathUrls {
		// Write data from our array of structs to the map
		pathsToUrls[pu.Path] = pu.URL
	}
	// return a map hander using the map

	return MapHandler(pathsToUrls, fallback), nil
}

// Custom type for parsing yaml, uses tags
type pathUrl struct {
	Path string `yaml:"path"`
	URL  string `yaml:"url"`
}
