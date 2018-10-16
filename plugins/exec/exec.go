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

	imgdata := "![](data:image/png;base64," + base64.StdEncoding.EncodeToString(data) + ")"
	io.WriteString(w, imgdata)
	io.WriteString(w, "\n")

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
	go func(reader io.Reader) {
		stdin.Write(in)
		stdin.Close()
		data, err = ioutil.ReadAll(reader)
	}(reader)

	if err := cmd.Start(); err != nil {
		return nil, err
	}
	cmd.Wait()
	return data, err
}
