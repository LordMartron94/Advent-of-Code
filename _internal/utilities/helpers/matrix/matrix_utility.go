package matrix

import (
	"fmt"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/extensions"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/pathfinding/shared"
)

type Distance struct {
	PosA             Position
	PosB             Position
	cDistRaw         int
	rDistRaw         int
	TotalDistanceRaw int
	cDistAbs         int
	rDistAbs         int
	TotalDistanceAbs int
}

func (d Distance) String() string {
	return fmt.Sprintf("Distance from %s to %s: rDistAbs=%d, cDistAbs=%d, TotalDistanceAbs=%d",
		d.PosA, d.PosB, d.rDistAbs, d.cDistAbs, d.TotalDistanceAbs)
}

func (mH *MatrixHelper[T]) PrintMatrix() {
	for _, row := range mH.itemsInMatrixNormalRows {
		fmt.Println(extensions.GetFormattedString(row))
	}
}

func (mH *MatrixHelper[T]) OutOfBounds(row, col int) bool {
	return row < 0 || row >= mH.rowCount || col < 0 || col >= mH.columnCount
}

func (mH *MatrixHelper[T]) GetCoordinatesOfTypes(types []T, customEqualityComparer *func(a, b T) bool) ([]T, [][]Position) {
	var comparer = mH.equalityComparer
	if customEqualityComparer != nil {
		comparer = *customEqualityComparer
	}

	returnTypes := make([]T, 0)
	positions := make([][]Position, 0)

	for row := 0; row < mH.rowCount; row++ {
		for col := 0; col < mH.columnCount; col++ {
			currentType := mH.itemsInMatrixNormalRows[row][col]

			for _, t := range types {
				if comparer(currentType, t) {
					typeIndex := extensions.SliceGetIndexOfEqualityComparer(returnTypes, currentType, comparer)

					if typeIndex == -1 {
						returnTypes = append(returnTypes, currentType)
						positions = append(positions, []Position{{RowIndex: row, ColIndex: col}})
						break
					} else {
						positions[typeIndex] = append(positions[typeIndex], Position{RowIndex: row, ColIndex: col})
						break
					}
				}
			}
		}
	}

	return returnTypes, positions
}

func (mH *MatrixHelper[T]) GetCoordinatesOfTypesFiltered(types []T, customEqualityComparer *func(a, b T) bool, filterComparer func(a, b T) bool) ([]T, [][]Position) {
	var comparer = mH.equalityComparer
	if customEqualityComparer != nil {
		comparer = *customEqualityComparer
	}

	returnTypes := make([]T, 0)
	positions := make([][]Position, 0)

	for row := 0; row < mH.rowCount; row++ {
		for col := 0; col < mH.columnCount; col++ {
			currentType := mH.itemsInMatrixNormalRows[row][col]

			for _, t := range types {
				if comparer(currentType, t) {
					typeIndex := extensions.SliceGetIndexOfEqualityComparer(returnTypes, currentType, filterComparer)

					if typeIndex == -1 {
						returnTypes = append(returnTypes, currentType)
						positions = append(positions, []Position{{RowIndex: row, ColIndex: col}})
						break
					} else {
						positions[typeIndex] = append(positions[typeIndex], Position{RowIndex: row, ColIndex: col})
						break
					}
				}
			}
		}
	}

	return returnTypes, positions
}

func abs(x int) int {
	if x < 0 {
		return -x
	}
	return x
}

func (mH *MatrixHelper[T]) FindDistance(posA, posB Position) Distance {
	cDistRaw := posA.ColIndex - posB.ColIndex
	rDistRaw := posA.RowIndex - posB.RowIndex
	cDistAbs := abs(posA.ColIndex - posB.ColIndex)
	rDistAbs := abs(posA.RowIndex - posB.RowIndex)

	return Distance{
		PosA:             posA,
		PosB:             posB,
		cDistRaw:         cDistRaw,
		rDistRaw:         rDistRaw,
		TotalDistanceRaw: cDistRaw + rDistRaw,
		cDistAbs:         cDistAbs,
		rDistAbs:         rDistAbs,
		TotalDistanceAbs: cDistAbs + rDistAbs,
	}
}

func (mH *MatrixHelper[T]) PositionsOnSameLine(positions []Position) bool {
	if len(positions) < 2 {
		return true
	}

	// Use the first two positions to define the line
	posA := positions[0]
	posB := positions[1]

	// Calculate the delta values (differences)
	deltaRow := posA.RowIndex - posB.RowIndex
	deltaCol := posA.ColIndex - posB.ColIndex

	// Handle vertical lines (deltaCol == 0)
	if deltaCol == 0 {
		for _, pos := range positions {
			if pos.ColIndex != posA.ColIndex {
				return false
			}
		}
		return true
	}

	// Check if all other positions lie on the line using a cross-product-like technique
	for _, pos := range positions[2:] {
		deltaRow2 := posA.RowIndex - pos.RowIndex
		deltaCol2 := posA.ColIndex - pos.ColIndex

		// If the points are collinear, the following condition will hold
		if deltaRow*deltaCol2 != deltaRow2*deltaCol {
			return false
		}
	}

	return true
}

func (mH *MatrixHelper[T]) AggregateUniqueDistancesBetweenPositions(positions []Position) []Distance {
	distances := make([]Distance, 0)

	for i := 0; i < len(positions); i++ {
		for j := i + 1; j < len(positions); j++ {
			distance := mH.FindDistance(positions[i], positions[j])
			distances = append(distances, distance)
		}
	}

	return distances
}

func (mH *MatrixHelper[T]) GetExtendedLinePositions(distance Distance) (before *Position, after *Position) {
	// Direction
	direction := shared.Direction{
		DeltaR: distance.rDistRaw,
		DeltaC: distance.cDistRaw,
	}

	// Get the position 'before' posA
	bf := distance.PosA.AddDirection(direction, 1)
	before = &bf

	// Get the position 'after' posB
	af := distance.PosB.AddDirection(direction, -1)
	after = &af

	if mH.OutOfBounds(bf.RowIndex, bf.ColIndex) {
		before = nil
	}

	if mH.OutOfBounds(af.RowIndex, af.ColIndex) {
		after = nil
	}

	return before, after
}

func (mH *MatrixHelper[T]) GetLinePositions(distance Distance) (positions []Position) {
	// Direction
	direction := shared.Direction{
		DeltaR: distance.rDistRaw,
		DeltaC: distance.cDistRaw,
	}

	currentPos := distance.PosA
	positions = append(positions, currentPos)

	// Move in the 'before' direction
	for {
		nextPos := currentPos.AddDirection(direction, 1)
		if mH.OutOfBounds(nextPos.RowIndex, nextPos.ColIndex) {
			break
		}
		positions = append(positions, nextPos)
		currentPos = nextPos
	}

	currentPos = distance.PosA
	for {
		nextPos := currentPos.AddDirection(direction, -1)
		if mH.OutOfBounds(nextPos.RowIndex, nextPos.ColIndex) {
			break
		}
		positions = append(positions, nextPos)
		currentPos = nextPos
	}

	return positions
}
