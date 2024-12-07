package factory

import (
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/matrix"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/pathfinding/shared"
)

type BasePathfindingRule[T any] struct {
	getMatch               func(currentPosition matrix.Position, currentDirection shared.Direction, finder FinderInterface[T], getNextItemFunc func(matrix.Position) (*T, error)) (bool, error)
	getNewDirection        func(currentPosition matrix.Position, currentDirection shared.Direction) shared.Direction
	getNewPosition         func(currentPosition matrix.Position, newDirection shared.Direction, finder FinderInterface[T]) matrix.Position
	DirectionNeedsPosition bool
}

func (b *BasePathfindingRule[T]) getNextItem(finder FinderInterface[T], pos matrix.Position) (*T, error) {
	if finder.OutOfBounds(pos) {
		return nil, &shared.OutOfBoundsError{}
	}

	item := finder.GetItemAtPosition(pos)
	return &item, nil
}

func (b *BasePathfindingRule[T]) MatchFunc(currentPosition matrix.Position, currentDirection shared.Direction, finder FinderInterface[T]) (bool, error) {
	return b.getMatch(currentPosition, currentDirection, finder, func(position matrix.Position) (*T, error) {
		return b.getNextItem(finder, position)
	})
}

func (b *BasePathfindingRule[T]) GetNewPosition(currentPosition matrix.Position, currentDirection shared.Direction, finder FinderInterface[T]) matrix.Position {
	newDirection := b.GetNewDirection(currentPosition, currentDirection)
	return b.getNewPosition(currentPosition, newDirection, finder)
}

func (b *BasePathfindingRule[T]) GetNewDirection(currentPosition matrix.Position, currentDirection shared.Direction) shared.Direction {
	return b.getNewDirection(currentPosition, currentDirection)
}
