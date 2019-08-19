package emph

import (
	"io"

	"github.com/gomarkdown/markdown/ast"
	"github.com/mmarkdown/mmark/render/markdown"
)

type Plugin struct {
	*markdown.Renderer
	marker string
}

// New implements the plugin interface.
func New() *Plugin {
	p := &Plugin{
		Renderer: markdown.NewRenderer(markdown.RendererOptions{}),
		marker:   "XX",
	}
	return p
}

func Name() string { return "emph" }

func (r *Plugin) Emph(w io.Writer, node *ast.Emph, entering bool) { io.WriteString(w, r.marker) }

func (r *Plugin) RenderNode(w io.Writer, node ast.Node, entering bool) ast.WalkStatus {
	if emph, ok := node.(*ast.Emph); ok {
		r.Emph(w, emph, entering)
		return ast.GoToNext
	}
	return r.Renderer.RenderNode(w, node, entering)
}
