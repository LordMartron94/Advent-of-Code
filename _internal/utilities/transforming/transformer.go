package transforming

import (
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
	shared2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming/shared"
)

type TransformFindFunc func(node *shared.ParseTree) (shared2.TransformCallback, int)

// Transformer takes a parsetree and applies a specified callback transformation to it
type Transformer struct {
	callbackFinder   TransformFindFunc
	callbacksByOrder map[int][]shared2.TransformCallback
	callbackNodes    map[int][]*shared.ParseTree
}

// NewTransformer creates a new Transformer with the given callbackFinders
func NewTransformer(callbackFinder TransformFindFunc) *Transformer {
	return &Transformer{
		callbackFinder:   callbackFinder,
		callbacksByOrder: make(map[int][]shared2.TransformCallback),
		callbackNodes:    make(map[int][]*shared.ParseTree),
	}
}

// Transform applies the transformations to the given parsetree recursively
func (t *Transformer) Transform(tree *shared.ParseTree) {
	t.transformRecursive(tree)

	// execute callbacks in order of appearance
	for callbackOrder, callbacks := range t.callbacksByOrder {
		for callbackIndex, callback := range callbacks {
			node := t.callbackNodes[callbackOrder][callbackIndex]

			if callback != nil {
				callback(node)
			}
		}
	}
}

func (t *Transformer) transformRecursive(tree *shared.ParseTree) {
	// Produce callbacks for the current node
	callback, order := t.callbackFinder(tree)
	t.callbacksByOrder[order] = append(t.callbacksByOrder[order], callback)
	t.callbackNodes[order] = append(t.callbackNodes[order], tree)

	// Recursively process children
	for _, node := range tree.Children {
		t.transformRecursive(node)
	}
}
