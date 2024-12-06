package matrix

import (
	"fmt"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/extensions"
)

func (mH *MatrixHelper[T]) PrintMatrix() {
	for _, row := range mH.itemsInMatrixNormalRows {
		fmt.Println(extensions.GetFormattedString(row))
	}
}

func (mH *MatrixHelper[T]) OutOfBounds(row, col int) bool {
	return row < 0 || row >= mH.rowCount || col < 0 || col >= mH.columnCount
}
