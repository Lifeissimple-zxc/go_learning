package urlshort

import (
	"reflect"
	"testing"
)

// Tests correct parsing of a valid yaml
func TestParseYAMLOk(t *testing.T) {
	validYaml := `
- path: /urlshort
  url: https://github.com/gophercises/urlshort
- path: /urlshort-final
  url: https://github.com/gophercises/urlshort/tree/solution
`
	exp := make([]pathUrl, 2)
	exp[0] = pathUrl{
		Path: "/urlshort",
		URL:  "https://github.com/gophercises/urlshort",
	}
	exp[1] = pathUrl{
		Path: "/urlshort-final",
		URL:  "https://github.com/gophercises/urlshort/tree/solution",
	}
	// Ignoring the error coz it's a test
	res, _ := parseYAML([]byte(validYaml))
	// Using refelct bc direct comparsion is not possible
	if !reflect.DeepEqual(exp, res) {
		t.Errorf("TestParseYamlOk test failed. expected: %v, got: %v", exp, res)
	}

}

// Tests parsing bad yaml, should result in an error
func TestParseYAMLErr(t *testing.T) {
	badYaml := `ur: gog path: https://google.com`
	// We skip error because res comparsion against nil is enough
	res, _ := parseYAML([]byte(badYaml))
	if res != nil {
		t.Errorf("TestParseYamlErr test failed. Res should be nil. Got %v instead", res)
	}

}

// Tests creation if map from a slice of a slice of pathUrl structs
func TestCreateUrlPathsMapOk(t *testing.T) {
	pathUrls := make([]pathUrl, 1)
	pathUrls[0] = pathUrl{Path: "gog", URL: "https://google.com"}
	exp := map[string]string{
		"gog": "https://google.com",
	}
	res := createUrlPathsMap(pathUrls)

	// Using reflect library here because cannot compare maps using == in GO
	if !reflect.DeepEqual(exp, res) {
		t.Errorf("TestCreateUrlPathsMapOk test failed. expected: %v, got: %v", exp, res)
	}
}

func TestParseJSONOk(t *testing.T) {
	validJSON := `
[
	{
		"path": "/goog",
		"url": "https://google.com/search?q=hello+world"
	},
	{
		"path": "/anime",
		"url": "https://myanimelist.net/"
	}
]
`
	exp := make([]pathUrl, 2)
	exp[0] = pathUrl{
		Path: "/goog",
		URL:  "https://google.com/search?q=hello+world",
	}
	exp[1] = pathUrl{
		Path: "/anime",
		URL:  "https://myanimelist.net/",
	}
	// Ignoring the error coz it's a test
	res, _ := parseYAML([]byte(validJSON))
	// Using refelct bc direct comparsion is not possible
	if !reflect.DeepEqual(exp, res) {
		t.Errorf("TestParseJSONOk test failed. expected: %v, got: %v", exp, res)
	}
}

func TestParseJSONErr(t *testing.T) {
	badJSON := `
[
	{
		"pth": "/goog",
		"ul": "https://google.com/search?q=hello+world"
	}
]
`
	exp := make([]pathUrl, 1)
	exp[0] = pathUrl{}
	// We skip error because res comparsion against nil is enough
	res, _ := parseJSON([]byte(badJSON))
	if !reflect.DeepEqual(exp, res) {
		t.Errorf("TestParseJSONErr test failed. Res has to be an empty pathUrl struct. Got %v  data instead", res)
	}
}
