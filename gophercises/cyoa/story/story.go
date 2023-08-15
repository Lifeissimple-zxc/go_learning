package story

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
)

const STORY_PATH = "gopher.json"

type StoryOption struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type StoryArc struct {
	Title   string        `json:"title"`
	Story   []string      `json:"story"`
	Options []StoryOption `json:"options"`
}

// StoryArc's init instantiates a StoryArc from json.RawMessage
func (sa *StoryArc) Init(data *json.RawMessage) error {
	if err := json.Unmarshal(*data, sa); err != nil {
		return err
	}
	return nil
}

// RenderHTML uses StoryArc data to generate
// an HTML page using HTMLBytes
func (sa StoryArc) RenderHTML(HTML []byte, name string, wr io.Writer) error {
	html, err := template.New(name).Parse(string(HTML))
	if err != nil {
		return err
	}
	if err = html.Execute(wr, sa); err != nil {
		return err
	} else {
		return nil
	}
}

type Story struct {
	Arcs map[string]StoryArc `json:"-"`
}

// Story's init converts raw json to a Story struct
func (s *Story) Init(rawData map[string]json.RawMessage) error {
	arcs := make(map[string]StoryArc)
	for arcKey, arcVal := range rawData {
		var arc StoryArc
		if err := arc.Init(&arcVal); err != nil {
			fmt.Printf("Err parsing %s data: %v\n", arcKey, err)
			// Exiting because missing arcs == broken story
			return err
		}
		arcs[arcKey] = arc
	}
	s.Arcs = arcs
	return nil

}
