package pathfinding

import (
	"context"
	"fmt"
	"slices"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/matrix"
)

// IsLoopSimple checks if the current position is part of a loop formed by the given path and current direction.
func (pf *PathFinder[T]) IsLoopSimple(currentPosition matrix.Position, currentDirection Direction, path map[matrix.Position][]Direction) bool {
	if slices.Contains(path[currentPosition], currentDirection) {
		return true
	}

	return false
}

func (pf *PathFinder[T]) doesMatrixLoop(startItem T, startDirection Direction) (bool, error) {
	if !pf.ruleSet.IsBasic {
		return false, fmt.Errorf("loop detection only supported for basic rule sets")
	}

	startPos, err := pf.getStartingPosition(startItem)

	if err != nil {
		return false, err
	}

	looping := false

	ctx, cancel := context.WithCancel(context.Background())
	path := make(map[matrix.Position][]Direction)
	err = pf.followPath(
		&FollowPathContext{
			Position:                   startPos,
			Direction:                  startDirection,
			Path:                       path,
			estimatedDirectionCapacity: 3,
		},
		[]func(_ FollowPathContext){},
		[]func(_ FollowPathContext){
			func(pathContext FollowPathContext) {
				if pf.IsLoopSimple(pathContext.Position, pathContext.Direction, path) {
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
func (pf *PathFinder[T]) Seen(currentPosition matrix.Position, path map[matrix.Position][]Direction) bool {
	_, exists := path[currentPosition]
	return exists
}
