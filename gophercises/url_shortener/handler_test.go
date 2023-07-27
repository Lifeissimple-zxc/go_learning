package urlshort

import (
	"reflect"
	"testing"
)

// Tests correct parsing of a valid yaml
func TestParseYamlOk(t *testing.T) {
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
	res, _ := parseYaml([]byte(validYaml))
	// Using refelct bc direct comparsion is not possible
	if !reflect.DeepEqual(exp, res) {
		t.Errorf("TestParseYamlOk test failed. expected: %v, got: %v", exp, res)
	}

}

// Tests parsing bad yaml, should result in an error
func TestParseYamlErr(t *testing.T) {
	badYaml := `ur: gog path: https://google.com`
	// We skip error because res comparsion against nil is enough
	res, _ := parseYaml([]byte(badYaml))
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
