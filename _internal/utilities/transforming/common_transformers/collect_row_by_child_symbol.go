package common_transformers

import (
	"slices"

	shared3 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
	shared2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming/shared"
)

func CollectRowByChildSymbols[T comparable](symbols []string, storage *[][]shared3.Token[T]) shared2.TransformCallback[T] {
	return func(node *shared.ParseTree[T]) {
		row := make([]shared3.Token[T], 0)

		for _, child := range node.Children {
			if slices.Contains(symbols, child.Symbol) {
				row = append(row, *child.Token)
			}
		}

		*storage = append(*storage, row)
	}
}
