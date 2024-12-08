package matrix

import (
	"fmt"

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

// GetMatrixVariation returns a new matrix with the specified position replaced with the replacement value.
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
