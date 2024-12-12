package factory

import (
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/matrix"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/matrix/shared"
)

type BasePathfindingRule[T any] struct {
	getMatch               func(finder FinderInterface[T], nextTiles []T) int
	getNewDirection        func(currentPosition matrix.Position, currentDirection shared.Direction) shared.Direction
	getNewPosition         func(currentPosition matrix.Position, newDirection shared.Direction, finder FinderInterface[T]) matrix.Position
	directionNeedsPosition bool
}

func (b *BasePathfindingRule[T]) MatchFunc(finder FinderInterface[T], nextTiles []T) int {
	return b.getMatch(finder, nextTiles)
}

func (b *BasePathfindingRule[T]) GetNewPosition(currentPosition matrix.Position, currentDirection shared.Direction, finder FinderInterface[T]) matrix.Position {
	newDirection := b.GetNewDirection(currentPosition, currentDirection)
	return b.getNewPosition(currentPosition, newDirection, finder)
}

func (b *BasePathfindingRule[T]) GetNewDirection(currentPosition matrix.Position, currentDirection shared.Direction) shared.Direction {
	return b.getNewDirection(currentPosition, currentDirection)
}

func (b *BasePathfindingRule[T]) GetDirectionNeedsPosition() bool {
	return b.directionNeedsPosition
}
