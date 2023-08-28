package story

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io"
)

const STORY_PATH = "gopher.json" // TODO move to CLI args?

type StoryOption struct {
	Text string `json:"text"`
	Arc  string `json:"arc"`
}

type StoryArc struct {
	Title   string        `json:"title"`
	IsStart bool          `json:"is_start"`
	Story   []string      `json:"story"`
	Options []StoryOption `json:"options"`
	Tmpl    *template.Template
}

// StoryArc's init instantiates a StoryArc from json.RawMessage
func (sa *StoryArc) Init(
	data *json.RawMessage,
	HTML []byte,
	name string,
) error {
	if err := json.Unmarshal(*data, sa); err != nil {
		return fmt.Errorf("error unmarshalling arc JSON: %v", err)
	}
	html, err := template.New(name).Parse(string(HTML))
	if err != nil {
		return fmt.Errorf("error parsing HTML template: %v", err)
	}
	sa.Tmpl = html
	return nil
}

// RenderHTML uses StoryArc data to generate
// an HTML page using HTMLBytes
func (sa *StoryArc) RenderHTML(wr io.Writer) error {
	if err := sa.Tmpl.Execute(wr, sa); err != nil {
		return err
	} else {
		return nil
	}
}

type Story struct {
	StartArc string
	Arcs     map[string]StoryArc `json:"-"`
}

// Story's init converts raw json to a Story struct
func (s *Story) Init(
	rawData map[string]json.RawMessage,
	HTML []byte,
) error {
	arcs := make(map[string]StoryArc)
	for arcKey, arcVal := range rawData {
		var arc StoryArc
		if err := arc.Init(&arcVal, HTML, arcKey); err != nil {
			// Exiting because missing arcs == broken story
			return fmt.Errorf("err parsing %s data: %v", arcKey, err)
		}
		arcs[arcKey] = arc
		if arc.IsStart {
			s.StartArc = arcKey
		}
	}
	s.Arcs = arcs
	return nil
}
