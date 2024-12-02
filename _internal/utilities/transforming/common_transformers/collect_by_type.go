package common_transformers

import (
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
	shared2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming/shared"
)

func CollectNodesByType(nodeType string, target *[]*shared.ParseTree) shared2.TransformCallback {
	return func(node *shared.ParseTree) {
		if node.Symbol == nodeType {
			*target = append(*target, node)
		}
	}
}
