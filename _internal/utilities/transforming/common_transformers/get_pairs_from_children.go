package common_transformers

import (
	"fmt"

	shared3 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
	shared2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming/shared"
)

// GetPairsFromChildren collects children
func GetPairsFromChildren[T comparable](pairSymbol string, storage *[][]shared3.Token[T]) shared2.TransformCallback[T] {
	return func(node *shared.ParseTree[T]) {
		for _, child := range node.Children {
			if child.Symbol == pairSymbol {
				if len(child.Children) < 2 || len(child.Children) > 2 {
					fmt.Println("Invalid pair node. Expected 2 children, got", len(child.Children))
					return
				}

				pair := make([]shared3.Token[T], 2)

				pair[0] = *child.Children[0].Token
				pair[1] = *child.Children[1].Token

				*storage = append(*storage, pair)
			}
		}
	}
}
