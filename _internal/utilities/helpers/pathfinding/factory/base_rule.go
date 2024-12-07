package factory

import (
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/matrix"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/pathfinding"
)

type BasePathfindingRule[T any] struct {
	getMatch        func(currentPosition matrix.Position, currentDirection pathfinding.Direction, finder pathfinding.PathFinder[T], getNextItemFunc func(matrix.Position) (*T, error)) (bool, error)
	getNewDirection func(currentPosition matrix.Position, currentDirection pathfinding.Direction) pathfinding.Direction
	getNewPosition  func(currentPosition matrix.Position, newDirection pathfinding.Direction, finder pathfinding.PathFinder[T]) matrix.Position
}

func (b *BasePathfindingRule[T]) getNextItem(finder pathfinding.PathFinder[T], pos matrix.Position) (*T, error) {
	if finder.OutOfBounds(pos) {
		return nil, &pathfinding.OutOfBoundsError{}
	}

	item := finder.GetItemAtPosition(pos)
	return &item, nil
}

func (b *BasePathfindingRule[T]) MatchFunc(currentPosition matrix.Position, currentDirection pathfinding.Direction, finder pathfinding.PathFinder[T]) (bool, error) {
	return b.getMatch(currentPosition, currentDirection, finder, func(position matrix.Position) (*T, error) {
		return b.getNextItem(finder, position)
	})
}

func (b *BasePathfindingRule[T]) GetNewPosition(currentPosition matrix.Position, currentDirection pathfinding.Direction, finder pathfinding.PathFinder[T]) matrix.Position {
	newDirection := b.GetNewDirection(currentPosition, currentDirection)
	return b.getNewPosition(currentPosition, newDirection, finder)
}

func (b *BasePathfindingRule[T]) GetNewDirection(currentPosition matrix.Position, currentDirection pathfinding.Direction) pathfinding.Direction {
	return b.getNewDirection(currentPosition, currentDirection)
}
