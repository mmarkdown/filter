package protocol

import (
	"bufio"
	"bytes"
	"io"
	"io/ioutil"
	"log"
	"os"
	"os/exec"
	"sync"

	"github.com/gomarkdown/markdown/ast"
	"github.com/mmarkdown/mmark/render/markdown"
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

func Name() string { return "protocol" }

const proto = "protocol"

func (r *Plugin) RenderNode(w io.Writer, node ast.Node, entering bool) ast.WalkStatus {
	codeblock, ok := node.(*ast.CodeBlock)
	if !ok {
		return r.Renderer.RenderNode(w, node, entering)
	}
	if !bytes.Equal(codeblock.Info, []byte(proto)) {
		return r.Renderer.RenderNode(w, node, entering)
	}

	data, err := run(proto, codeblock.Literal)
	if err != nil {
		log.Printf("Failed to run %q: %s\n", proto, err)
		return r.Renderer.RenderNode(w, node, entering)
	}
	codeblock.Literal = data
	codeblock.Info = nil
	return r.Renderer.RenderNode(w, node, entering)
}

func run(path string, in []byte) ([]byte, error) {
	// instead of giving 'in' on stdin, we need to supply it on the command line.
	cmd := exec.Command(path, string(in))
	cmd.Stderr = os.Stderr
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
		data, err = ioutil.ReadAll(reader)
	}(reader)

	if err := cmd.Start(); err != nil {
		return nil, err
	}
	cmd.Wait()
	wg.Wait()
	return data, err
}
