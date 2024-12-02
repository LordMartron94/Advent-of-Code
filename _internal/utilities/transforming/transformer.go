package transforming

import (
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
	shared2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming/shared"
)

type TransformFindFunc func(node *shared.ParseTree) shared2.TransformCallback

// Transformer takes a parsetree and applies a specified callback transformation to it
type Transformer struct {
	callbackFinder TransformFindFunc
}

// NewTransformer creates a new Transformer with the given callbackFinders
func NewTransformer(callbackFinder TransformFindFunc) *Transformer {
	return &Transformer{
		callbackFinder: callbackFinder,
	}
}

// Transform applies the transformations to the given parsetree recursively
func (t *Transformer) Transform(tree *shared.ParseTree) {
	for _, node := range tree.Children {
		callback := t.callbackFinder(node)
		if callback != nil {
			callback(node)
		}

		if node.Children != nil {
			t.Transform(node)
		}
	}
}
