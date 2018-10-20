package rot13

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

func Name() string { return "rot13" }

func (r *Plugin) RenderNode(w io.Writer, node ast.Node, entering bool) ast.WalkStatus {
	text, ok := node.(*ast.Text)
	if !ok {
		return r.Renderer.RenderNode(w, node, entering)
	}
	txt := markdown.EscapeText(text.Literal)
	for _, c := range txt {
		w.Write([]byte{rot13(c)})
	}
	return ast.GoToNext
}

func rot13(b byte) byte {
	var a, z byte
	switch {
	case 'a' <= b && b <= 'z':
		a, z = 'a', 'z'
	case 'A' <= b && b <= 'Z':
		a, z = 'A', 'Z'
	default:
		return b
	}
	return (b-a+13)%(z-a+1) + a
}
