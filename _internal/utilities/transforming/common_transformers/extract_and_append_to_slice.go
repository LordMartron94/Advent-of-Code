package common_transformers

import (
	"fmt"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
	shared2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming/shared"
)

func AppendTokenValueToSlice[T any](slice *[]T, converter func(string) (T, error)) shared2.TransformCallback {
	return func(node *shared.ParseTree) {
		value, err := converter(string(node.Token.Value))
		if err != nil {
			fmt.Printf("error converting token value to desired type: %v\n", err)
			return
		}
		*slice = append(*slice, value)
	}
}

func AppendTokenValueToSliceSorted[T any](slice *[]T, converter func(string) (T, error), sortFunc func([]T)) shared2.TransformCallback {
	return func(node *shared.ParseTree) {
		value, err := converter(string(node.Token.Value))
		if err != nil {
			fmt.Printf("error converting token value to desired type: %v\n", err)
			return
		}
		*slice = append(*slice, value)
		sortFunc(*slice)
	}
}
