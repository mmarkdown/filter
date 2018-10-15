package main

import (
	"github.com/gomarkdown/markdown"
	"github.com/mmarkdown/filter/plugins/emph"
	"github.com/mmarkdown/filter/plugins/noop"
)

type Plugin interface {
	New() *markdown.Renderer
}

var Loaded = []markdown.Renderer{
	noop.New(),
	emph.New(),
}
