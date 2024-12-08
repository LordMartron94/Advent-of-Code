package pathfinding

import (
	"context"
	"errors"
	"fmt"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/matrix"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/pathfinding/rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/pathfinding/rules/factory"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/pathfinding/shared"
)

type DirectionExternal int

const (
	Down DirectionExternal = iota
	Left
	Right
	Up
)

func (d DirectionExternal) ToDirection() shared.Direction {
	switch d {
	case Down:
		return shared.Direction{DeltaR: 1, DeltaC: 0}
	case Left:
		return shared.Direction{DeltaR: 0, DeltaC: -1}
	case Right:
		return shared.Direction{DeltaR: 0, DeltaC: 1}
	case Up:
		return shared.Direction{DeltaR: -1, DeltaC: 0}
	default:
		return shared.Direction{}
	}
}

type PathFinder[T any] struct {
	matrixHelper    *matrix.MatrixHelper[T]
	ruleSet         *rules.PathfindingRuleset[T]
	equalityChecker func(a, b T) bool
	debugMode       bool
}

func NewPathFinder[T any](matrixToUse [][]T, equalityChecker func(a, b T) bool, rulesToUse []factory.PathfindingRuleInterface[T], basicRules, debug bool) *PathFinder[T] {
	ruleset := rules.NewPathfindingRuleset(rulesToUse, basicRules)

	return &PathFinder[T]{
		matrixHelper:    matrix.NewMatrixHelper(matrixToUse, equalityChecker),
		ruleSet:         ruleset,
		equalityChecker: equalityChecker,
		debugMode:       debug,
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

type FollowPathContext struct {
	Position                   matrix.Position
	Direction                  shared.Direction
	Path                       *[]matrix.Position
	Directions                 *[][]shared.Direction
	currentPathIndex           int
	estimatedDirectionCapacity int
	numOfDirectionBatches      int
}

func (pf *PathFinder[T]) appendToPath(fpCtx *FollowPathContext) {
	// We only need to initialize the directions slice if we are at the cap and need to insert a new one.
	if fpCtx.currentPathIndex == len(*fpCtx.Directions) {
		for range fpCtx.numOfDirectionBatches {
			*fpCtx.Directions = append(*fpCtx.Directions, make([]shared.Direction, 0, fpCtx.estimatedDirectionCapacity))
		}
	}

	*fpCtx.Path = append(*fpCtx.Path, fpCtx.Position)
	(*fpCtx.Directions)[fpCtx.currentPathIndex] = append((*fpCtx.Directions)[fpCtx.currentPathIndex], fpCtx.Direction)
	fpCtx.currentPathIndex++
}

// followPath follows the path defined by the given ruleset.
func (pf *PathFinder[T]) followPath(
	fpCtx *FollowPathContext,
	beforeNewPositionCallbacks []func(ctx FollowPathContext),
	afterNewPositionCallbacks []func(ctx FollowPathContext),
	ctx context.Context,
) error {
	if fpCtx.numOfDirectionBatches < 1 {
		panic("At least one direction batch is required")
	}

	select {
	case <-ctx.Done():
		return nil
	default:
		for _, beforeNewPositionCallback := range beforeNewPositionCallbacks {
			beforeNewPositionCallback(*fpCtx)
		}

		pf.appendToPath(fpCtx)

		rule, _, err := pf.ruleSet.GetRule(fpCtx.Position, fpCtx.Direction, pf, pf.GetItemAtPosition(fpCtx.Position))

		if err != nil {
			if errors.Is(err, &shared.OutOfBoundsError{}) {
				return nil
			}

			return err
		}

		newDirection := rule.GetNewDirection(fpCtx.Position, fpCtx.Direction)
		newPosition := rule.GetNewPosition(fpCtx.Position, fpCtx.Direction, pf)

		if pf.debugMode {
			fmt.Println(fmt.Sprintf("Old position: (%d, %d), old direction: (%d, %d)", fpCtx.Position.RowIndex+1, fpCtx.Position.ColIndex+1, fpCtx.Direction.DeltaR, fpCtx.Direction.DeltaC))
			fmt.Println(fmt.Sprintf("New position: (%d, %d), new direction: (%d, %d)", newPosition.RowIndex+1, newPosition.ColIndex+1, newDirection.DeltaR, newDirection.DeltaC))
			fmt.Println("---------------")
		}

		fpCtx.Position = newPosition
		fpCtx.Direction = newDirection

		for _, afterNewPositionCallback := range afterNewPositionCallbacks {
			afterNewPositionCallback(*fpCtx)
		}

		return pf.followPath(fpCtx, beforeNewPositionCallbacks, afterNewPositionCallbacks, ctx)
	}
}

func (pf *PathFinder[T]) GetNumberOfStepsUntilOutOfBounds(startItem T, startDirection DirectionExternal) (int, error) {
	startPos, err := pf.getStartingPosition(startItem)

	if err != nil {
		return 0, err
	}

	path := make([]matrix.Position, 0, 300)
	directions := make([][]shared.Direction, 0, 300)
	err = pf.followPath(
		&FollowPathContext{
			Position:                   startPos,
			Direction:                  startDirection.ToDirection(),
			Path:                       &path,
			Directions:                 &directions,
			numOfDirectionBatches:      10,
			currentPathIndex:           0,
			estimatedDirectionCapacity: 3,
		},
		[]func(_ FollowPathContext){},
		[]func(_ FollowPathContext){},
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

	path := make([]matrix.Position, 0, 300)
	directions := make([][]shared.Direction, 0, 300)
	err = pf.followPath(
		&FollowPathContext{
			Position:                   startPos,
			Direction:                  startDirection.ToDirection(),
			Path:                       &path,
			Directions:                 &directions,
			numOfDirectionBatches:      10,
			currentPathIndex:           0,
			estimatedDirectionCapacity: 3,
		},
		[]func(pathContext FollowPathContext){
			func(pathContext FollowPathContext) {
				if !pf.Seen(pathContext.Position, &path) {
					visitedNodeCount++
				}
			},
		},
		[]func(_ FollowPathContext){},
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
func (pf *PathFinder[T]) DoesMatrixLoop(startItem T, startDirection shared.Direction) (bool, error) {
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

func (pf *PathFinder[T]) GetPositionInDirection(position matrix.Position, direction shared.Direction, n int) matrix.Position {
	newPos := position
	newPos.RowIndex += direction.DeltaR * n
	newPos.ColIndex += direction.DeltaC * n

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

func (pf *PathFinder[T]) GetTilesInDirection(pos matrix.Position, direction shared.Direction) []T {
	output := make([]T, 0)

	posToCheck := pf.GetPositionInDirection(pos, direction, 1)

	for !pf.OutOfBounds(posToCheck) {
		output = append(output, pf.GetItemAtPosition(posToCheck))
		posToCheck = pf.GetPositionInDirection(posToCheck, direction, 1)
	}

	return output
}

func (pf *PathFinder[T]) EqualityCheck(a, b T) bool {
	return pf.equalityChecker(a, b)
}
