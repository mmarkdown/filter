//+build ignore

package main

import (
	"bytes"
	"fmt"
	"go/format"
	"io/ioutil"
	"log"
	"os"
)

func main() {
	dirs, err := ioutil.ReadDir("plugins")
	if err != nil {
		log.Fatal(err)
	}
	b := &bytes.Buffer{}
	fmt.Fprintf(b, `
type Plugin interface {
	New() markdown.Renderer
	Name() string
}

var Plugins = map[string]markdown.Renderer{
`)
	prog := &bytes.Buffer{}
	fmt.Fprintf(prog, "package main\n\nimport (\n\"github.com/gomarkdown/markdown\"\n")

	for _, d := range dirs {
		if !d.IsDir() { // each plugin sits in a subdir
			continue
		}
		fmt.Fprintf(prog, "\""+"github.com/mmarkdown/filter/plugins/"+d.Name()+"\"\n")
		fmt.Fprintf(b, d.Name()+".Name():\t"+d.Name()+".New(),\n")
	}
	fmt.Fprintf(prog, ")\n\n")
	fmt.Fprintf(b, "}\n")
	prog.Write(b.Bytes())

	res, err := format.Source(prog.Bytes())
	if err != nil {
		prog.WriteTo(os.Stderr)
		log.Fatal(err)
	}

	f, err := os.Create("zplugins.go")
	if err != nil {
		log.Fatal(err)
	}
	defer f.Close()
	f.Write(res)
}
