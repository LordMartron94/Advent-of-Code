package pathfinding

import (
	"context"
	"fmt"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/matrix"
)

type Rule[T any] struct {
	MatchFunc       func(currentPosition matrix.Position, currentDirection Direction, finder PathFinder[T]) bool
	GetNewPosition  func(currentPosition matrix.Position, currentDirection Direction, finder PathFinder[T]) matrix.Position
	GetNewDirection func(currentPosition matrix.Position, currentDirection Direction) Direction
}

type Ruleset[T any] struct {
	IsBasic bool // Set to true if the ruleset is simple (follow current direction until obstacle, if obstacle, do x)... Two rules.
	Rules   []Rule[T]
}

type Direction struct {
	deltaR int
	deltaC int
}

func (d *Direction) Turn() Direction {
	return Direction{deltaR: d.deltaC, deltaC: -d.deltaR}
}

type DirectionExternal int

const (
	Down DirectionExternal = iota
	Left
	Right
	Up
)

func (d DirectionExternal) ToDirection() Direction {
	switch d {
	case Down:
		return Direction{deltaR: 1, deltaC: 0}
	case Left:
		return Direction{deltaR: 0, deltaC: -1}
	case Right:
		return Direction{deltaR: 0, deltaC: 1}
	case Up:
		return Direction{deltaR: -1, deltaC: 0}
	default:
		return Direction{}
	}
}

type PathFinder[T any] struct {
	matrixHelper    *matrix.MatrixHelper[T]
	ruleSet         *Ruleset[T]
	equalityChecker func(a, b T) bool
}

func NewPathFinder[T any](matrixToUse [][]T, equalityChecker func(a, b T) bool, ruleset Ruleset[T]) *PathFinder[T] {
	return &PathFinder[T]{
		matrixHelper:    matrix.NewMatrixHelper(matrixToUse, equalityChecker),
		ruleSet:         &ruleset,
		equalityChecker: equalityChecker,
	}
}

// getStartingPosition finds the starting position of the given item in the matrix.
func (pf *PathFinder[T]) getStartingPosition(startItem T) (matrix.Position, error) {
	foundPos := pf.matrixHelper.GetPositionOfTarget(startItem)

	if foundPos == nil {
		return matrix.Position{}, fmt.Errorf("start item not found in the matrix")
	}

	return *foundPos, nil
}

func (pf *PathFinder[T]) getRule(currentPosition matrix.Position, currentDirection Direction) (*Rule[T], error) {
	for _, rule := range pf.ruleSet.Rules {
		if rule.MatchFunc(currentPosition, currentDirection, *pf) {
			return &rule, nil
		}
	}

	return nil, fmt.Errorf("no matching rule found")
}

func (pf *PathFinder[T]) followPath(
	position matrix.Position,
	direction Direction,
	path map[matrix.Position][]Direction,
	beforeNewPositionCallbacks []func(position matrix.Position, direction Direction, path map[matrix.Position][]Direction),
	afterNewPositionCallbacks []func(position matrix.Position, direction Direction, path map[matrix.Position][]Direction),
	ctx context.Context,
) error {
	select {
	case <-ctx.Done():
		return nil
	default:
		for _, beforeNewPositionCallback := range beforeNewPositionCallbacks {
			beforeNewPositionCallback(position, direction, path)
		}

		path[position] = append(path[position], direction)

		rule, err := pf.getRule(position, direction)

		if err != nil {
			return err
		}

		newDirection := rule.GetNewDirection(position, direction)
		newPosition := rule.GetNewPosition(position, direction, *pf)

		if pf.OutOfBounds(newPosition) {
			return nil
		}

		if err != nil {
			return err
		}

		for _, afterNewPositionCallback := range afterNewPositionCallbacks {
			afterNewPositionCallback(newPosition, newDirection, path)
		}

		return pf.followPath(newPosition, newDirection, path, beforeNewPositionCallbacks, afterNewPositionCallbacks, ctx)
	}
}

func (pf *PathFinder[T]) GetNumberOfStepsUntilOutOfBounds(startItem T, startDirection DirectionExternal) (int, error) {
	startPos, err := pf.getStartingPosition(startItem)

	if err != nil {
		return 0, err
	}

	path := make(map[matrix.Position][]Direction)
	err = pf.followPath(
		startPos,
		startDirection.ToDirection(),
		path,
		[]func(_ matrix.Position, _ Direction, _ map[matrix.Position][]Direction){},
		[]func(position matrix.Position, _ Direction, _ map[matrix.Position][]Direction){},
		context.Background(),
	)

	if err != nil {
		return len(path), err
	}

	return 0, fmt.Errorf("no path found")
}

func (pf *PathFinder[T]) GetNumberOfUniqueNodesVisitedUntilOutOfBounds(startItem T, startDirection DirectionExternal) (int, error) {
	startPos, err := pf.getStartingPosition(startItem)

	if err != nil {
		return 0, err
	}
	visitedNodeCount := 0
	path := make(map[matrix.Position][]Direction)
	err = pf.followPath(
		startPos,
		startDirection.ToDirection(),
		path,
		[]func(position matrix.Position, _ Direction, _ map[matrix.Position][]Direction){
			func(position matrix.Position, _ Direction, path map[matrix.Position][]Direction) {
				if !pf.Seen(position, path) {
					visitedNodeCount++
				}
			},
		},
		[]func(position matrix.Position, _ Direction, _ map[matrix.Position][]Direction){},
		context.Background(),
	)

	if err != nil {
		return visitedNodeCount, err
	}

	return visitedNodeCount, nil
}

// SetMatrix sets the matrix for the pathfinder.
func (pf *PathFinder[T]) SetMatrix(matrix [][]T) {
	pf.matrixHelper.SetMatrix(matrix)
}

// DoesMatrixLoop checks if the matrix contains a loop with the ruleset.
func (pf *PathFinder[T]) DoesMatrixLoop(startItem T, startDirection Direction) (bool, error) {
	return pf.doesMatrixLoop(startItem, startDirection)
}

func (pf *PathFinder[T]) DoesOtherMatrixLoop(otherMatrix [][]T, startItem T, startDirection DirectionExternal) (bool, error) {
	currentMatrix := pf.matrixHelper.GetRows()
	pf.SetMatrix(otherMatrix)
	looping, err := pf.DoesMatrixLoop(startItem, startDirection.ToDirection())
	pf.SetMatrix(currentMatrix)

	return looping, err
}

func (pf *PathFinder[T]) GetNumberOfLoopingMatrices(otherMatrices [][][]T, startItem T, startDirection DirectionExternal) (int, error) {
	loopingMatrixCount := 0
	for _, otherMatrix := range otherMatrices {
		looping, err := pf.DoesOtherMatrixLoop(otherMatrix, startItem, startDirection)
		if err != nil {
			return loopingMatrixCount, err
		}
		if looping {
			loopingMatrixCount++
		}
	}

	return loopingMatrixCount, nil
}

func (pf *PathFinder[T]) GetNumberOfLoopingMatricesForGeneratedVariations(startItem T, startDirection DirectionExternal, replaceX, replaceWith T) (int, error) {
	loopingMatrixCount := 0

	for r, row := range pf.matrixHelper.GetRows() {
		for c, item := range row {
			if pf.equalityChecker(item, replaceX) {
				pf.matrixHelper.ReplaceValueInPlace(r, c, replaceWith)
				looping, err := pf.DoesMatrixLoop(startItem, startDirection.ToDirection())

				if err != nil {
					return loopingMatrixCount, err
				}

				if looping {
					loopingMatrixCount++
				}

				pf.matrixHelper.ReplaceValueInPlace(r, c, item)
			}
		}
	}

	return loopingMatrixCount, nil
}

func (pf *PathFinder[T]) GetPositionInDirectory(position matrix.Position, direction Direction, n int) matrix.Position {
	newPos := position
	newPos.RowIndex += direction.deltaR * n
	newPos.ColIndex += direction.deltaC * n

	return newPos
}

func (pf *PathFinder[T]) OutOfBounds(pos matrix.Position) bool {
	if pf.matrixHelper.OutOfBounds(pos.RowIndex, pos.ColIndex) {
		return true
	}

	return false
}

func (pf *PathFinder[T]) GetItemAtPosition(pos matrix.Position) T {
	return pf.matrixHelper.GetAtPosition(pos.RowIndex, pos.ColIndex)
}
