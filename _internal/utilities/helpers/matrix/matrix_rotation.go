package matrix

// rotateMatrixNormal rotates the matrix by 45 degrees from top-right to bottom-left.
func (mH *MatrixHelper[T]) rotateMatrixNormal() [][]T {
	// Inspired by: https://www.geeksforgeeks.org/rotate-matrix-by-45-degrees/

	height := len(mH.itemsInMatrixNormalRows)   // mH
	width := len(mH.itemsInMatrixNormalRows[0]) // n

	diagonals := make([][]T, 0)

	counter := 0
	for counter < width+height-1 {
		diagonal := make([]T, 0)

		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if i+j == counter {
					diagonal = append(diagonal, mH.itemsInMatrixNormalRows[i][j])
				}
			}
		}

		diagonals = append(diagonals, diagonal)
		counter++
	}

	return diagonals
}

// rotateMatrixReverse rotates the matrix by 45 degrees from top-left to bottom-right.
func (mH *MatrixHelper[T]) rotateMatrixReverse() [][]T {
	height := len(mH.itemsInMatrixNormalRows)   // mH
	width := len(mH.itemsInMatrixNormalRows[0]) // n

	diagonals := make([][]T, 0)

	for counter := 1 - height; counter < width; counter++ {
		diagonal := make([]T, 0)

		for i := 0; i < height; i++ {
			for j := 0; j < width; j++ {
				if i-j == counter {
					diagonal = append(diagonal, mH.itemsInMatrixNormalRows[i][j])
				}
			}
		}

		diagonals = append(diagonals, diagonal)
	}

	return diagonals
}
