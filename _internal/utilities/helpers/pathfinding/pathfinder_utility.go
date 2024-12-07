package pathfinding

import (
	"context"
	"fmt"
	"slices"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/matrix"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/pathfinding/shared"
)

// IsLoopSimple checks if the current position is part of a loop formed by the given path and current direction.
func (pf *PathFinder[T]) IsLoopSimple(currentPosition matrix.Position, currentDirection shared.Direction, path *[]matrix.Position, directions *[][]shared.Direction) bool {
	indexOfCurrentPosition := slices.IndexFunc(*path, func(p matrix.Position) bool {
		return p == currentPosition
	})

	if indexOfCurrentPosition == -1 {
		return false
	}

	sliceToCheck := *directions

	if slices.Contains(sliceToCheck[indexOfCurrentPosition], currentDirection) {
		return true
	}

	return false
}

func (pf *PathFinder[T]) doesMatrixLoop(startItem T, startDirection shared.Direction) (bool, error) {
	if !pf.ruleSet.IsBasic {
		return false, fmt.Errorf("loop detection only supported for basic rule sets")
	}

	startPos, err := pf.getStartingPosition(startItem)

	if err != nil {
		return false, err
	}

	looping := false

	ctx, cancel := context.WithCancel(context.Background())
	path := make([]matrix.Position, 0, 300)
	directions := make([][]shared.Direction, 0, 300)
	err = pf.followPath(
		&FollowPathContext{
			Position:                   startPos,
			Direction:                  startDirection,
			Path:                       &path,
			Directions:                 &directions,
			numOfDirectionBatches:      10,
			currentPathIndex:           0,
			estimatedDirectionCapacity: 3,
		},
		[]func(_ FollowPathContext){},
		[]func(_ FollowPathContext){
			func(pathContext FollowPathContext) {
				if pf.IsLoopSimple(pathContext.Position, pathContext.Direction, pathContext.Path, pathContext.Directions) {
					looping = true
					cancel()
				}
			},
		},
		ctx)

	if err != nil {
		return false, err
	}

	return looping, nil
}

// Seen checks if the current position has been visited before.
func (pf *PathFinder[T]) Seen(currentPosition matrix.Position, path *[]matrix.Position) bool {
	return slices.Contains(*path, currentPosition)
}
