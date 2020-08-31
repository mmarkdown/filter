package main

import (
	"github.com/gomarkdown/markdown"
	"github.com/mmarkdown/filter/plugins/emph"
	"github.com/mmarkdown/filter/plugins/exec"
	"github.com/mmarkdown/filter/plugins/noop"
	"github.com/mmarkdown/filter/plugins/protocol"
	"github.com/mmarkdown/filter/plugins/rot13"
)

type Plugin interface {
	New() markdown.Renderer
	Name() string
}

var Plugins = map[string]markdown.Renderer{
	emph.Name():     emph.New(),
	exec.Name():     exec.New(),
	noop.Name():     noop.New(),
	protocol.Name(): protocol.New(),
	rot13.Name():    rot13.New(),
}
