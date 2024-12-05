package factory

import (
	"fmt"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
	shared2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
)

type BaseParsingRule[T comparable] struct {
	SymbolString string

	isShell        bool // IsShell indicates that the top element is not a token in itself
	matchFunc      func(tokens []*shared.Token[T], currentIndex int) (bool, string)
	getContentFunc func(tokens []*shared.Token[T], currentIndex int) *shared2.ParseTree[T]
}

func (b *BaseParsingRule[T]) Symbol() string {
	return b.SymbolString
}

func (b *BaseParsingRule[T]) Match(tokens []*shared.Token[T], currentIndex int) (*shared2.ParseTree[T], error, int) {
	matched, errorMessage := b.matchFunc(tokens, currentIndex)

	if !matched {
		return nil, fmt.Errorf(errorMessage), 0
	}

	tree := b.getContentFunc(tokens, currentIndex)

	if b.isShell {
		return tree, nil, len(tree.Children)
	}

	return tree, nil, len(tree.Children) + 1
}
