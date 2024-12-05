package transforming

import (
	"fmt"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
)

// GridHelper is a helper struct for working with a 2D grid of tokens.
type GridHelper[T comparable] struct {
	tokensInGrid [][]shared.Token[T]
}

func NewGridHelper[T comparable](tokensInGrid [][]shared.Token[T]) *GridHelper[T] {
	return &GridHelper[T]{
		tokensInGrid: tokensInGrid,
	}
}

// reverseSlice reverses a slice of tokens.
func (g *GridHelper[T]) reverseSlice(tokens []T) []T {
	newSlice := make([]T, len(tokens))

	for i := 0; i < len(tokens); i++ {
		newSlice[i] = tokens[len(tokens)-1-i]
	}

	return newSlice
}

// rotateMatrixNormal rotates the matrix by 45 degrees from top-right to bottom-left.
func (g *GridHelper[T]) rotateMatrixNormal() [][]shared.Token[T] {
	// Inspired by: https://www.geeksforgeeks.org/rotate-matrix-by-45-degrees/

	height := len(g.tokensInGrid)   // m
	width := len(g.tokensInGrid[0]) // n

	diagonals := make([][]shared.Token[T], 0)

	counter := 0
	for counter < width+height-1 {
		diagonal := make([]shared.Token[T], 0)

		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if i+j == counter {
					diagonal = append(diagonal, g.tokensInGrid[i][j])
				}
			}
		}

		diagonals = append(diagonals, diagonal)
		counter++
	}

	return diagonals
}

// rotateMatrixReverse rotates the matrix by 45 degrees from top-left to bottom-right.
func (g *GridHelper[T]) rotateMatrixReverse() [][]shared.Token[T] {
	height := len(g.tokensInGrid)   // m
	width := len(g.tokensInGrid[0]) // n

	diagonals := make([][]shared.Token[T], 0)

	for counter := 1 - height; counter < width; counter++ {
		diagonal := make([]shared.Token[T], 0)

		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if i-j == counter {
					diagonal = append(diagonal, g.tokensInGrid[i][j])
				}
			}
		}

		diagonals = append(diagonals, diagonal)
	}

	return diagonals
}

func (g *GridHelper[T]) getColumns() [][]shared.Token[T] {
	verticalLines := make([][]shared.Token[T], len(g.tokensInGrid[0]))

	for i := range g.tokensInGrid[0] {
		for j := range g.tokensInGrid {
			verticalLines[i] = append(verticalLines[i], g.tokensInGrid[j][i])
		}
	}

	return verticalLines
}

func (g *GridHelper[T]) findMatchesInSlice(slice []shared.Token[T], tokenTypes []T, reverseAllowed bool) int {
	matches := 0

	for i := 0; i < len(slice)-len(tokenTypes)+1; i++ {
		match := true
		for j := 0; j < len(tokenTypes); j++ {
			if slice[i+j].Type != tokenTypes[j] {
				match = false
				break
			}
		}
		if match {
			matches++
		}
	}

	if reverseAllowed {
		newTokenTypes := g.reverseSlice(tokenTypes)
		matches += g.findMatchesInSlice(slice, newTokenTypes, false)
	}

	return matches
}

func (g *GridHelper[T]) FindHorizontalConsecutiveTokensNumber(tokenTypes []T, reverseAllowed bool) int {
	matches := 0

	for _, row := range g.tokensInGrid {
		matches += g.findMatchesInSlice(row, tokenTypes, reverseAllowed)
	}

	return matches
}

func (g *GridHelper[T]) FindVerticalConsecutiveTokensNumber(tokenTypes []T, reverseAllowed bool) int {
	matches := 0
	columns := g.getColumns()

	for _, column := range columns {
		matches += g.findMatchesInSlice(column, tokenTypes, reverseAllowed)
	}

	return matches
}

func (g *GridHelper[T]) getRowString(row []shared.Token[T]) string {
	rowString := ""
	for _, token := range row {
		rowString += fmt.Sprintf("%s ", token.Value)
	}
	return rowString[:len(rowString)-1]
}

func (g *GridHelper[T]) printMatrix() {
	for _, row := range g.tokensInGrid {
		fmt.Println(g.getRowString(row))
	}
}

func (g *GridHelper[T]) FindDiagonallyConsecutiveTokensNumber(tokenTypes []T, reverseAllowed bool) int {
	diagonals := g.rotateMatrixNormal()
	diagonals = append(diagonals, g.rotateMatrixReverse()...)

	matches := 0

	for _, diagonal := range diagonals {
		//for _, token := range diagonal {
		//	fmt.Println(fmt.Sprintf("Token: %s", token.Value))
		//}
		//
		//fmt.Println("------------------------")

		matches += g.findMatchesInSlice(diagonal, tokenTypes, reverseAllowed)
	}

	numOfDiagonals := len(diagonals)
	fmt.Println("Number of diagonals:", numOfDiagonals)

	//g.printMatrix()

	return matches
}
