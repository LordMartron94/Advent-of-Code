package common_transformers

import (
	shared3 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
	shared2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming/shared"
)

func AppendChildrenToSlice[T comparable](target *[][]shared3.Token[T]) shared2.TransformCallback[T] {
	return func(node *shared.ParseTree[T]) {
		current := make([]shared3.Token[T], 0)

		for _, child := range node.Children {
			current = append(current, *child.Token)
		}

		*target = append(*target, current)
	}
}
