package transforming

import (
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
	shared2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming/shared"
)

type TransformFindFunc[T comparable] func(node *shared.ParseTree[T]) (shared2.TransformCallback[T], int)

// Transformer takes a parsetree and applies a specified callback transformation to it
type Transformer[T comparable] struct {
	callbackFinder   TransformFindFunc[T]
	callbacksByOrder map[int][]shared2.TransformCallback[T]
	callbackNodes    map[int][]*shared.ParseTree[T]
}

// NewTransformer creates a new Transformer with the given callbackFinders
func NewTransformer[T comparable](callbackFinder TransformFindFunc[T]) *Transformer[T] {
	return &Transformer[T]{
		callbackFinder:   callbackFinder,
		callbacksByOrder: make(map[int][]shared2.TransformCallback[T]),
		callbackNodes:    make(map[int][]*shared.ParseTree[T]),
	}
}

// Transform applies the transformations to the given parsetree recursively
func (t *Transformer[T]) Transform(tree *shared.ParseTree[T]) {
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

func (t *Transformer[T]) transformRecursive(tree *shared.ParseTree[T]) {
	// Produce callbacks for the current node
	callback, order := t.callbackFinder(tree)
	t.callbacksByOrder[order] = append(t.callbacksByOrder[order], callback)
	t.callbackNodes[order] = append(t.callbackNodes[order], tree)

	// Recursively process children
	for _, node := range tree.Children {
		t.transformRecursive(node)
	}
}
