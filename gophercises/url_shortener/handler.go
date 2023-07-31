package urlshort

import (
	"encoding/json"
	"net/http"

	"gopkg.in/yaml.v2"
)

// Encapsulating some code from YAMLHandler for better testing

// Custom type for parsing yaml & json inputs, uses tags
type pathUrl struct {
	Path string `yaml:"path" json:"path"`
	URL  string `yaml:"url" json:"url"`
}

// Parses yaml into a slice of pathUrl structs
func parseYAML(yml []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl
	err := yaml.Unmarshal(yml, &pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, nil
}

// Prases json into a slice of pathUrl structs
func parseJSON(jsn []byte) ([]pathUrl, error) {
	var pathUrls []pathUrl
	err := json.Unmarshal(jsn, &pathUrls)
	if err != nil {
		return nil, err
	}
	return pathUrls, nil
}

// Converts slice of pathUrl to a map for quick lookups by the handler
func createUrlPathsMap(paths []pathUrl) map[string]string {
	// Container for our data
	pathsToUrls := make(map[string]string)
	// Iteratively write to container
	for _, pu := range paths {
		pathsToUrls[pu.Path] = pu.URL
	}
	return pathsToUrls
}

// MapHandler will return an http.HandlerFunc (which also
// implements http.Handler) that will attempt to map any
// paths (keys in the map) to their corresponding URL (values
// that each key in the map points to, in string format).
// If the path is not provided in the map, then the fallback
// http.Handler will be called instead.
func MapHandler(pathsToUrls map[string]string, fallback http.Handler) http.HandlerFunc {
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
	pathUrls, err := parseYAML(yml)
	if err != nil {
		return nil, err
	}
	// Convert yaml array to a map
	// Create a map-container of our data
	pathsToUrls := createUrlPathsMap(pathUrls)

	// return a map hander using the map
	return MapHandler(pathsToUrls, fallback), nil
}

// JSONHandler parses the provided JSON and returns
// an http.HandlerFunc (which also implements http.Handler)
// that will attempt to map any paths to their corresponding
// URL. If the path is not provided in the JSON, then the
// fallback http.Handler will be called instead.
//
// JSON is expected to be in the format:
//
//	[
//
//		{
//			"path": "/some-path",
//			"url": "https://www.some-url.com/demo"
//		},
//		{
//			"path2": "/some-other-path",
//			"url2": "https://www.some-other-url.com/demo"
//		}
//		{...}
//
//	]
//
// The only errors that can be returned all related to having
// invalid JSON data.
//
// See MapHandler to create a similar http.HandlerFunc via
// a mapping of paths to urls.
func JSONHandler(jsn []byte, fallback http.Handler) (http.HandlerFunc, error) {
	// Parse the JSON
	pathUrls, err := parseJSON(jsn)
	if err != nil {
		return nil, err
	}
	// Convert JSON array to a map
	// Create a map-container of our data
	pathsToUrls := createUrlPathsMap(pathUrls)

	// return a map hander using the map
	return MapHandler(pathsToUrls, fallback), nil
}
