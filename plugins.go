package main

import (
	"github.com/gomarkdown/markdown"
	"github.com/mmarkdown/filter/plugins/emph"
	"github.com/mmarkdown/filter/plugins/noop"
)

type Plugin interface {
	New() markdown.Renderer
	Name() string
}

var Plugins = map[string]markdown.Renderer{
	noop.Name(): noop.New(),
	emph.Name(): emph.New(),
}
