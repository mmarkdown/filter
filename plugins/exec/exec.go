package exec

import (
	"bufio"
	"bytes"
	"encoding/base64"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"sync"

	"github.com/gomarkdown/markdown/ast"
	"github.com/mmarkdown/mmark/markdown"
	"github.com/mmarkdown/mmark/mast"
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

func Name() string { return "exec" }

func (r *Plugin) RenderNode(w io.Writer, node ast.Node, entering bool) ast.WalkStatus {
	codeblock, ok := node.(*ast.CodeBlock)
	if !ok {
		return r.Renderer.RenderNode(w, node, entering)
	}
	if !bytes.HasPrefix(codeblock.Info, []byte("exec:")) {
		return r.Renderer.RenderNode(w, node, entering)
	}
	// it has a prefix
	cmd := codeblock.Info[len("exec:"):]
	if len(cmd) == 0 {
		return r.Renderer.RenderNode(w, node, entering)
	}

	data, err := run(string(cmd), codeblock.Literal)
	if err != nil {
		log.Printf("Failed to run %q: %s\n", string(cmd), err)
	}
	_, isCaptionFigure := codeblock.GetParent().(*ast.CaptionFigure)
	attr := mast.AttributeFromNode(codeblock)

	if isCaptionFigure || attr != nil {
		if attr != nil {
			w.Write(mast.AttributeBytes(attr))
			io.WriteString(w, "\n")
		}
		io.WriteString(w, "!---\n")
	}
	imgdata := "![](data:image/png;base64," + base64.StdEncoding.EncodeToString(data) + ")"
	io.WriteString(w, imgdata)
	io.WriteString(w, "\n")
	if isCaptionFigure || attr != nil {
		io.WriteString(w, "!---\n")
	}

	return ast.GoToNext

}

func run(path string, in []byte) ([]byte, error) {
	cmd := exec.Command(path)
	cmd.Stderr = os.Stderr
	stdin, err := cmd.StdinPipe()
	if err != nil {
		return nil, err
	}
	stdout, err := cmd.StdoutPipe()
	if err != nil {
		return nil, err
	}

	reader := bufio.NewReader(stdout)
	data := []byte{}
	wg := sync.WaitGroup{}
	go func(reader io.Reader) {
		wg.Add(1)
		defer wg.Done()
		stdin.Write(in)
		stdin.Close()
		data, err = ioutil.ReadAll(reader)
	}(reader)

	if err := cmd.Start(); err != nil {
		return nil, err
	}
	cmd.Wait()
	wg.Wait()
	return data, err
}
