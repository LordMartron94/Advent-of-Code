package pipes

import (
	"fmt"

	"github.com/LordMartron94/Advent-of-Code/2024/Day-04/pipeline/common"
	"github.com/LordMartron94/Advent-of-Code/2024/Day-04/task_rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/matrix"
	shared3 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming"
	shared2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming/shared"
)

type TransformDataPipe struct{}

func (t *TransformDataPipe) Process(input common.PipelineContext[task_rules.LexingTokenType]) common.PipelineContext[task_rules.LexingTokenType] {
	horizontalLines := make([][]shared3.Token[task_rules.LexingTokenType], 0)

	callbackFinder := func(node *shared.ParseTree[task_rules.LexingTokenType]) (shared2.TransformCallback[task_rules.LexingTokenType], int) {
		switch node.Symbol {
		case "horizontal_line":
			return func(node *shared.ParseTree[task_rules.LexingTokenType]) {
				currentLine := make([]shared3.Token[task_rules.LexingTokenType], 0)

				for _, token := range node.Children {
					currentLine = append(currentLine, *token.Token)
				}

				horizontalLines = append(horizontalLines, currentLine)
			}, 0
		}
		return nil, 0
	}

	transformer := transforming.NewTransformer(
		callbackFinder,
	)
	transformer.Transform(input.ParseTree)

	//numOfHorizontalLines := len(horizontalLines)
	//fmt.Println("Number of horizontal lines:", numOfHorizontalLines)

	gHelper := matrix.NewMatrixHelper(horizontalLines)

	numOfHorizontalMatches := gHelper.FindHorizontalConsecutiveTokensNumber([]task_rules.LexingTokenType{
		task_rules.XCharToken,
		task_rules.MCharToken,
		task_rules.ACharToken,
		task_rules.SCharToken,
	}, true)

	numOfVerticalMatches := gHelper.FindVerticalConsecutiveTokensNumber([]task_rules.LexingTokenType{
		task_rules.XCharToken,
		task_rules.MCharToken,
		task_rules.ACharToken,
		task_rules.SCharToken,
	}, true)

	numOfDiagonalMatches := gHelper.FindDiagonallyConsecutiveTokensNumber([]task_rules.LexingTokenType{
		task_rules.XCharToken,
		task_rules.MCharToken,
		task_rules.ACharToken,
		task_rules.SCharToken,
	}, true)

	numOfDiagonalMatchesMasX := getNumOfXMases(horizontalLines)

	fmt.Println("Number of horizontal matches:", numOfHorizontalMatches)
	fmt.Println("Number of vertical matches:", numOfVerticalMatches)
	fmt.Println("Number of diagonal matches:", numOfDiagonalMatches)
	fmt.Println("Number of matches (MAS X):", numOfDiagonalMatchesMasX)

	input.Result = numOfHorizontalMatches + numOfVerticalMatches + numOfDiagonalMatches
	input.XMasResult = numOfDiagonalMatchesMasX

	//expectedHorizontal := 5
	//expectedVertical := 3
	//expectedDiagonal := 10
	//expectedTotal := 18
	//
	//fmt.Println(fmt.Sprintf("Difference Expected & Gotten Total: %v & %v - '%d'", expectedTotal, input.Result, expectedTotal-input.Result))
	//fmt.Println(fmt.Sprintf("Difference Expected & Gotten Horizontal: %v & %v - '%d'", expectedHorizontal, numOfHorizontalMatches, expectedHorizontal-numOfHorizontalMatches))
	//fmt.Println(fmt.Sprintf("Difference Expected & Gotten Vertical: %v & %v - '%d'", expectedVertical, numOfVerticalMatches, expectedVertical-numOfVerticalMatches))
	//fmt.Println(fmt.Sprintf("Difference Expected & Gotten Diagonal: %v & %v - '%d'", expectedDiagonal, numOfDiagonalMatches, expectedDiagonal-numOfDiagonalMatches))

	return input
}

func getNumOfXMases(lines [][]shared3.Token[task_rules.LexingTokenType]) int {
	matches := 0

	diagonalCorrect := func(corner1, corner2 shared3.Token[task_rules.LexingTokenType]) bool {
		return (corner1.Type == task_rules.MCharToken && corner2.Type == task_rules.SCharToken) ||
			(corner1.Type == task_rules.SCharToken && corner2.Type == task_rules.MCharToken)
	}

	isValid := func(topLeft, topRight, bottomLeft, bottomRight shared3.Token[task_rules.LexingTokenType]) bool {
		firstDiagonalCorrect := diagonalCorrect(topLeft, bottomRight)
		secondDiagonalCorrect := diagonalCorrect(topRight, bottomLeft)

		return firstDiagonalCorrect && secondDiagonalCorrect
	}

	for i := 1; i < len(lines)-1; i++ {
		for j := 1; j < len(lines[i])-1; j++ {
			if lines[i][j].Type != task_rules.ACharToken {
				continue
			}

			if isValid(lines[i-1][j-1], lines[i-1][j+1], lines[i+1][j-1], lines[i+1][j+1]) {
				matches++

				//fmt.Println(fmt.Sprintf("Found XMAS match at (%d, %d)", i+1, j+1))
			}
		}
	}

	return matches
}
