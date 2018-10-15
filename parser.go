package main

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
	"github.com/mmarkdown/mmark/mparser"
)

func newParser() *parser.Parser {
	p := parser.NewWithExtensions(Extensions)

	p.Opts = parser.ParserOptions{
		ParserHook: func(data []byte) (ast.Node, []byte, int) {
			node, data, consumed := mparser.Hook(data)
			return node, data, consumed
		},
		Flags: parser.FlagsNone,
	}
	return p
}
