package common_transformers

import (
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
	shared2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming/shared"
)

func CombineCallbacks(callbacks ...shared2.TransformCallback) shared2.TransformCallback {
	return func(node *shared.ParseTree) {
		for _, callback := range callbacks {
			callback(node)
		}
	}
}
