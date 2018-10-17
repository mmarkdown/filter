package renderer

import (
	"github.com/gomarkdown/markdown/ast"
	"github.com/gomarkdown/markdown/parser"
	"github.com/mmarkdown/mmark/mparser"
)

// newParser() returns a pointer to a parser.Parser for our markdown.
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

// Extensions is exported to we can use it in tests. (copied from mmark - keep in sync)
var Extensions = parser.Tables | parser.FencedCode | parser.Autolink | parser.Strikethrough |
	parser.SpaceHeadings | parser.HeadingIDs | parser.BackslashLineBreak | parser.SuperSubscript |
	parser.DefinitionLists | parser.MathJax | parser.AutoHeadingIDs | parser.Footnotes |
	parser.Strikethrough | parser.OrderedListStart | parser.Attributes | parser.Mmark
