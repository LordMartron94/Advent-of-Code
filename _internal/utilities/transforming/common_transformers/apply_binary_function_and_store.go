package common_transformers

import (
	"fmt"
	"strconv"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
	shared2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming/shared"
)

// ApplyBinaryOperationToChildren applies a binary operation to the children of a given node and appends the result to the provided slice.
func ApplyBinaryOperationToChildren(operation func(left, right int) int, result *[]int) shared2.TransformCallback {
	return func(node *shared.ParseTree) {
		if len(node.Children) < 2 {
			fmt.Println("Invalid binary operation node. Too few children for binary operation. Expected 2, got", len(node.Children))
			return
		} else if len(node.Children) > 2 {
			fmt.Println("Invalid binary operation node. Too many children for binary operation. Expected 2, got", len(node.Children))
			return
		}

		leftNum, err := strconv.Atoi(string(node.Children[0].Token.Value))
		if err != nil {
			fmt.Printf("error converting token value to desired type: %v\n", err)
			return
		}

		rightNum, err := strconv.Atoi(string(node.Children[1].Token.Value))
		if err != nil {
			fmt.Printf("error converting token value to desired type: %v\n", err)
			return
		}

		resultValue := operation(leftNum, rightNum)

		//fmt.Println("Applying binary operation to children:", leftNum, "DISTANCE FUNC ->", rightNum, "=>", resultValue)

		*result = append(*result, resultValue)
	}
}
