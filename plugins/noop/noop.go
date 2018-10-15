package noop

import (
	"io"

	"github.com/gomarkdown/markdown/ast"
	"github.com/mmarkdown/mmark/markdown"
)

type Plugin struct {
	*markdown.Renderer
}

// New implements the plugin interface.
func New() *Plugin {
	return &Plugin{
		markdown.NewRenderer(markdown.RendererOptions{}),
	}
}

func (r *Plugin) RenderNode(w io.Writer, node ast.Node, entering bool) ast.WalkStatus {
	return r.Renderer.RenderNode(w, node, entering)
}
