package common_transformers

import (
	"fmt"

	shared3 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
	shared2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming/shared"
)

// GetPairsFromSpecificChildren collects children of symbol 1 and 2 and stores it in your given list..
func GetPairsFromSpecificChildren[T comparable](symbol1 string, symbol2 string, storage *[][]shared3.Token[T]) shared2.TransformCallback[T] {
	return func(node *shared.ParseTree[T]) {
		if len(node.Children) < 2 {
			fmt.Println(fmt.Sprintf("Invalid node. Expected at least 2 children, got %d", len(node.Children)))
			return
		}

		pair := make([]shared3.Token[T], 2)

		collectedT1 := false
		collectedT2 := false

		for _, child := range node.Children {
			if child.Symbol == symbol1 {
				if collectedT1 {
					fmt.Println("Duplicate type 1 token found.")
					return
				}

				pair[0] = *child.Token
				collectedT1 = true
			}
			if child.Symbol == symbol2 {
				if collectedT2 {
					fmt.Println("Duplicate type 2 token found.")
					return
				}

				pair[1] = *child.Token
				collectedT2 = true
			}
		}

		*storage = append(*storage, pair)
	}
}
