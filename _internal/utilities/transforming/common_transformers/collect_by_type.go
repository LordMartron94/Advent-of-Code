package common_transformers

import (
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
	shared2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming/shared"
)

func CollectNodesByType[T comparable](nodeType string, target *[]*shared.ParseTree[T]) shared2.TransformCallback[T] {
	return func(node *shared.ParseTree[T]) {
		if node.Symbol == nodeType {
			*target = append(*target, node)
		}
	}
}
