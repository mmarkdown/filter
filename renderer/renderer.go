package renderer

import (
	"github.com/gomarkdown/markdown"
)

// R holds the plugins that will be used to render the document.
type R struct {
	plugins []markdown.Renderer
}

// New returns a pointer to a new R.
func New() *R { return &R{} }

// RegisterPlugin registers a plugin in r.
func (r *R) RegisterPlugin(mr markdown.Renderer) {
	r.plugins = append(r.plugins, mr)
}

// Render renders the document through all contained renderers. It
// will basicly run: parse | render | parse | render.
func (r *R) Render(data []byte) []byte {
	if len(r.plugins) == 0 {
		return nil
	}
	for _, plugin := range r.plugins {
		doc := markdown.Parse(data, newParser())
		data = markdown.Render(doc, plugin)
	}
	return data
}
