package links

import (
	"github.com/yuin/goldmark"
	"github.com/yuin/goldmark/ast"
	"github.com/yuin/goldmark/parser"
	"github.com/yuin/goldmark/text"
	"github.com/yuin/goldmark/util"
)

type (
	linksExtension struct{}
)

const (
	// Used to signal to the rendering step that an image is used in a block context.
	// Dont's change this; the prefix must match the internalAttrPrefix in the root goldmark package.
	AttrIsBlock = "_h__isBlock"
)

func New() goldmark.Extender {
	return &linksExtension{}
}

func (e *linksExtension) Extend(m goldmark.Markdown) {
	m.Parser().AddOptions(
		parser.WithASTTransformers(
			util.Prioritized(&Transformer{}, 300),
		),
	)
}

type Transformer struct{}

// Transform transforms the provided Markdown AST.
func (t *Transformer) Transform(doc *ast.Document, reader text.Reader, pctx parser.Context) {
	ast.Walk(doc, func(node ast.Node, enter bool) (ast.WalkStatus, error) {
		if !enter {
			return ast.WalkContinue, nil
		}

		if n, ok := node.(*ast.AutoLink); ok {
			parent := n.Parent()

			isBlock := parent.ChildCount() == 1
			if isBlock {
				n.SetAttributeString(AttrIsBlock, true)
			}
		}

		if n, ok := node.(*ast.Link); ok {
			parent := n.Parent()

			isBlock := parent.ChildCount() == 1
			if isBlock {
				n.SetAttributeString(AttrIsBlock, true)
			}
		}

		return ast.WalkContinue, nil

	})

}
