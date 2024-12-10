package matrix

import (
	"fmt"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/extensions"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/pathfinding/shared"
)

type DiagonalDirection int

const (
	DiagonalTopRight DiagonalDirection = iota
	DiagonalTopLeft
)

type Position struct {
	RowIndex int
	ColIndex int
}

func (p Position) String() string {
	return fmt.Sprintf("(%d, %d)", p.RowIndex+1, p.ColIndex+1)
}

func (p Position) Add(pos Position) Position {
	return Position{
		RowIndex: p.RowIndex + pos.RowIndex,
		ColIndex: p.ColIndex + pos.ColIndex,
	}
}

func (p Position) Subtract(pos Position) Position {
	return Position{
		RowIndex: p.RowIndex - pos.RowIndex,
		ColIndex: p.ColIndex - pos.ColIndex,
	}
}

func (p Position) Scale(scaleFactor float64) Position {
	return Position{
		RowIndex: int(float64(p.RowIndex) * scaleFactor),
		ColIndex: int(float64(p.ColIndex) * scaleFactor),
	}
}

func (p Position) AddDirection(direction shared.Direction, amount int) Position {
	return Position{
		RowIndex: p.RowIndex + (direction.DeltaR * amount),
		ColIndex: p.ColIndex + (direction.DeltaC * amount),
	}
}

// MatrixHelper is a helper struct for working with a 2D matrix.
type MatrixHelper[T any] struct {
	itemsInMatrixNormalRows       [][]T
	itemsInMatrixNormalCols       *[][]T // Internal pointers so we can have a cacheing mechanism for optimal performance.
	itemsInMatrixDiagonalTopRight *[][]T
	itemsInMatrixDiagonalTopLeft  *[][]T
	equalityComparer              func(a, b T) bool

	rowCount    int
	columnCount int
}

func NewMatrixHelper[T any](rows [][]T, equalityComparer func(a, b T) bool) *MatrixHelper[T] {
	mH := &MatrixHelper[T]{
		itemsInMatrixNormalRows:       rows,
		rowCount:                      len(rows),
		columnCount:                   len(rows[0]),
		itemsInMatrixNormalCols:       nil,
		itemsInMatrixDiagonalTopRight: nil,
		itemsInMatrixDiagonalTopLeft:  nil,
		equalityComparer:              equalityComparer,
	}
	return mH
}

// initDiagonalItems initializes the diagonal items of the matrix.
func (mH *MatrixHelper[T]) initDiagonalItems() {
	normalRotation := mH.rotateMatrixNormal()
	mH.itemsInMatrixDiagonalTopRight = &normalRotation

	reverseRotation := mH.rotateMatrixReverse()
	mH.itemsInMatrixDiagonalTopLeft = &reverseRotation
}

func (mH *MatrixHelper[T]) initColumns() {
	rows := mH.itemsInMatrixNormalRows
	numRows := len(rows)
	numColumns := len(rows[0])
	verticalLines := make([][]T, numColumns)

	for r := 0; r < numRows; r++ {
		for c := 0; c < numColumns; c++ {
			verticalLines[c] = append(verticalLines[c], rows[r][c])
		}
	}

	mH.itemsInMatrixNormalCols = &verticalLines
}

// GetDiagonals returns the diagonal items of the matrix.
func (mH *MatrixHelper[T]) GetDiagonals(direction DiagonalDirection) [][]T {
	if mH.itemsInMatrixDiagonalTopRight == nil || mH.itemsInMatrixDiagonalTopLeft == nil {
		mH.initDiagonalItems()
	}

	switch direction {
	case DiagonalTopRight:
		return *mH.itemsInMatrixDiagonalTopRight
	case DiagonalTopLeft:
		return *mH.itemsInMatrixDiagonalTopLeft
	default:
		panic("Invalid diagonal direction") // This will never happen.
	}
}

// GetColumns returns the columns of the matrix.
func (mH *MatrixHelper[T]) GetColumns() [][]T {
	if mH.itemsInMatrixNormalCols == nil {
		mH.initColumns()
	}

	return *mH.itemsInMatrixNormalCols
}

// GetRows returns the rows of the matrix.
func (mH *MatrixHelper[T]) GetRows() [][]T {
	return mH.itemsInMatrixNormalRows
}

// SetMatrix sets the matrix items.
func (mH *MatrixHelper[T]) SetMatrix(matrix [][]T) {
	mH.itemsInMatrixNormalRows = matrix
	mH.rowCount = len(matrix)
	mH.columnCount = len(matrix[0])
	mH.itemsInMatrixNormalCols = nil
	mH.itemsInMatrixDiagonalTopRight = nil
	mH.itemsInMatrixDiagonalTopLeft = nil
}

func (mH *MatrixHelper[T]) GetAtPosition(rowIndex int, colIndex int) T {
	if mH.OutOfBounds(rowIndex, colIndex) {
		panic("Index out of bounds")
	}

	return mH.itemsInMatrixNormalRows[rowIndex][colIndex]
}

// GetMatrixVariation returns a new matrix with the specified Position replaced with the replacement value.
func (mH *MatrixHelper[T]) GetMatrixVariation(r int, c int, value T) [][]T {
	newMatrix := make([][]T, mH.rowCount)
	for rI := 0; rI < mH.rowCount; rI++ {
		newRow := make([]T, mH.columnCount)
		for cI := 0; cI < mH.columnCount; cI++ {
			if rI == r && cI == c {
				newRow[cI] = value
			} else {
				newRow[cI] = mH.itemsInMatrixNormalRows[rI][cI]
			}
		}
		newMatrix[rI] = newRow
	}
	return newMatrix
}

func (mH *MatrixHelper[T]) ReplaceValueInPlace(r int, c int, value T) {
	if mH.OutOfBounds(r, c) {
		panic("Index out of bounds")
	}

	mH.itemsInMatrixNormalRows[r][c] = value
}

type Neighbor[T any] struct {
	Position Position
	Value    T
}

func (n Neighbor[T]) String() string {
	return fmt.Sprintf("[%s, %v]", n.Position, n.Value)
}

// GetAdjacencyListHorizontalVertical creates an adjacency list from a matrix,
// considering only horizontal and vertical neighbors, and preserving neighbor positions.
func (mH *MatrixHelper[T]) GetAdjacencyListHorizontalVertical() ([][]Neighbor[T], int) {
	adjacencyList := make([][]Neighbor[T], mH.rowCount*mH.columnCount)
	nodeCount := 0

	for r := 0; r < mH.rowCount; r++ {
		for c := 0; c < mH.columnCount; c++ {
			neighbors := make([]Neighbor[T], 0)

			if !mH.OutOfBounds(r-1, c) {
				neighbors = append(neighbors, Neighbor[T]{Position{RowIndex: r - 1, ColIndex: c}, mH.GetAtPosition(r-1, c)})
			}
			if !mH.OutOfBounds(r+1, c) {
				neighbors = append(neighbors, Neighbor[T]{Position{RowIndex: r + 1, ColIndex: c}, mH.GetAtPosition(r+1, c)})
			}
			if !mH.OutOfBounds(r, c-1) {
				neighbors = append(neighbors, Neighbor[T]{Position{RowIndex: r, ColIndex: c - 1}, mH.GetAtPosition(r, c-1)})
			}
			if !mH.OutOfBounds(r, c+1) {
				neighbors = append(neighbors, Neighbor[T]{Position{RowIndex: r, ColIndex: c + 1}, mH.GetAtPosition(r, c+1)})
			}

			index := r*mH.columnCount + c
			adjacencyList[index] = neighbors
			nodeCount++
		}
	}

	return adjacencyList, nodeCount
}

func (mH *MatrixHelper[T]) PrintAdjacencyListHorizontalVertical() {
	adjacencyList, _ := mH.GetAdjacencyListHorizontalVertical()

	for i, neighbors := range adjacencyList {
		row := i / mH.columnCount
		col := i % mH.columnCount
		fmt.Printf("(%d, %d) - %s\n", row, col, extensions.GetFormattedString(neighbors))
	}
}

func (mH *MatrixHelper[T]) GetColumnCount() int {
	return mH.columnCount
}

func (mH *MatrixHelper[T]) GetRowCount() int {
	return mH.rowCount
}
