package matrix

import "github.com/LordMartron94/Advent-of-Code/_internal/utilities/extensions"

// FindConsecutiveMatchesNumberInRows finds the number of consecutive matches in the matrix's rows.
func (mH *MatrixHelper[T]) FindConsecutiveMatchesNumberInRows(targets []T, reverseAllowed bool) int {
	matches := 0
	rows := mH.itemsInMatrixNormalRows

	for _, row := range rows {
		matches += extensions.FindNumberOfMatchesInSliceV2(row, targets, reverseAllowed, mH.equalityComparer)
	}

	return 0
}

// FindConsecutiveMatchesNumberInColumns finds the number of consecutive matches in the matrix's columns.
func (mH *MatrixHelper[T]) FindConsecutiveMatchesNumberInColumns(targets []T, reverseAllowed bool) int {
	matches := 0

	if mH.itemsInMatrixNormalCols == nil {
		mH.initColumns()
	}

	columns := *mH.itemsInMatrixNormalCols

	for _, column := range columns {
		matches += extensions.FindNumberOfMatchesInSliceV2(column, targets, reverseAllowed, mH.equalityComparer)
	}

	return matches
}

// FindConsecutiveMatchesNumberInDiagonal finds the number of consecutive matches in the matrix's chosen diagonal.
func (mH *MatrixHelper[T]) FindConsecutiveMatchesNumberInDiagonal(targets []T, reverseAllowed bool, direction DiagonalDirection) int {
	matches := 0

	if mH.itemsInMatrixDiagonalTopRight == nil || mH.itemsInMatrixDiagonalTopLeft == nil {
		mH.initDiagonalItems()
	}

	var diagonals [][]T

	switch direction {
	case DiagonalTopLeft:
		diagonals = *mH.itemsInMatrixDiagonalTopLeft
		break
	case DiagonalTopRight:
		diagonals = *mH.itemsInMatrixDiagonalTopRight
		break
	}

	for _, diagonal := range diagonals {
		matches += extensions.FindNumberOfMatchesInSliceV2(diagonal, targets, reverseAllowed, mH.equalityComparer)
	}

	return matches
}

// FindConsecutiveMatchesNumberInDiagonals finds the number of consecutive matches in the matrix's diagonals.
func (mH *MatrixHelper[T]) FindConsecutiveMatchesNumberInDiagonals(targets []T, reverseAllowed bool) int {
	return mH.FindConsecutiveMatchesNumberInDiagonal(targets, reverseAllowed, DiagonalTopLeft) +
		mH.FindConsecutiveMatchesNumberInDiagonal(targets, reverseAllowed, DiagonalTopRight)
}

func (mH *MatrixHelper[T]) GetPositionOfTarget(target T) *Position {
	for r, row := range mH.itemsInMatrixNormalRows {
		for c, item := range row {
			if mH.equalityComparer(item, target) {
				return &Position{RowIndex: r, ColIndex: c}
			}
		}
	}

	return nil
}
