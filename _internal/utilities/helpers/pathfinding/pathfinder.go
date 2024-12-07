package pathfinding

import (
	"context"
	"errors"
	"fmt"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/matrix"
)

type OutOfBoundsError struct{}

func (e *OutOfBoundsError) Error() string {
	return "out of bounds"
}

type PathfindingRuleInterface[T any] interface {
	MatchFunc(currentPosition matrix.Position, currentDirection Direction, finder PathFinder[T]) (bool, error)
	GetNewPosition(currentPosition matrix.Position, currentDirection Direction, finder PathFinder[T]) matrix.Position
	GetNewDirection(currentPosition matrix.Position, currentDirection Direction) Direction
}

type PathfindingRuleset[T any] struct {
	IsBasic bool // Set to true if the ruleset is simple (follow current direction until obstacle, if obstacle, do x)... Two rules.
	Rules   []PathfindingRuleInterface[T]
}

type Direction struct {
	deltaR int
	deltaC int
}

func (d *Direction) TurnRight() Direction {
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
	ruleSet         *PathfindingRuleset[T]
	equalityChecker func(a, b T) bool
	debugMode       bool
}

func NewPathFinder[T any](matrixToUse [][]T, equalityChecker func(a, b T) bool, ruleset PathfindingRuleset[T], debug bool) *PathFinder[T] {
	return &PathFinder[T]{
		matrixHelper:    matrix.NewMatrixHelper(matrixToUse, equalityChecker),
		ruleSet:         &ruleset,
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

func (pf *PathFinder[T]) getRule(currentPosition matrix.Position, currentDirection Direction) (PathfindingRuleInterface[T], error) {
	for _, rule := range pf.ruleSet.Rules {
		match, err := rule.MatchFunc(currentPosition, currentDirection, *pf)

		if err != nil {
			return nil, err
		}

		if match {
			return rule, nil
		}
	}

	return nil, fmt.Errorf("no matching rule found")
}

type FollowPathContext struct {
	Position                   matrix.Position
	Direction                  Direction
	Path                       *[]matrix.Position
	Directions                 *[][]Direction
	currentPathIndex           int
	estimatedDirectionCapacity int
	numOfDirectionBatches      int
}

func (pf *PathFinder[T]) appendToPath(fpCtx *FollowPathContext) {
	// We only need to initialize the directions slice if we are at the cap and need to insert a new one.
	if fpCtx.currentPathIndex == len(*fpCtx.Directions) {
		for range fpCtx.numOfDirectionBatches {
			*fpCtx.Directions = append(*fpCtx.Directions, make([]Direction, 0, fpCtx.estimatedDirectionCapacity))
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

		// getRule has out of bounds check built in.
		rule, err := pf.getRule(fpCtx.Position, fpCtx.Direction)

		if err != nil {
			var err *OutOfBoundsError
			if errors.As(err, &err) {
				return nil
			}

			return err
		}

		newDirection := rule.GetNewDirection(fpCtx.Position, fpCtx.Direction)
		newPosition := rule.GetNewPosition(fpCtx.Position, fpCtx.Direction, *pf)

		if pf.debugMode {
			fmt.Println(fmt.Sprintf("Old position: (%d, %d), old direction: (%d, %d)", fpCtx.Position.RowIndex+1, fpCtx.Position.ColIndex+1, fpCtx.Direction.deltaR, fpCtx.Direction.deltaC))
			fmt.Println(fmt.Sprintf("New position: (%d, %d), new direction: (%d, %d)", newPosition.RowIndex+1, newPosition.ColIndex+1, newDirection.deltaR, newDirection.deltaC))
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
	directions := make([][]Direction, 0, 300)
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
	directions := make([][]Direction, 0, 300)
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

func (pf *PathFinder[T]) GetPositionInDirection(position matrix.Position, direction Direction, n int) matrix.Position {
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

func (pf *PathFinder[T]) EqualityCheck(a, b T) bool {
	return pf.equalityChecker(a, b)
}
