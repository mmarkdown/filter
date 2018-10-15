package main

import (
	"github.com/gomarkdown/markdown"
)

// Renderer renders the document through all contained renderers. It
// will basicly run: parse | render | parse | render.
type Renderer struct {
	plugins []markdown.Renderer
}

func (r *Renderer) RegisterPlugin(mr markdown.Renderer) {
	r.plugins = append(r.plugins, mr)
}

func (r *Renderer) Render(data []byte) []byte {
	if len(r.plugins) == 0 {
		return nil
	}
	for _, plugin := range r.plugins {
		doc := markdown.Parse(data, newParser())
		data = markdown.Render(doc, plugin)
	}
	return data
}
