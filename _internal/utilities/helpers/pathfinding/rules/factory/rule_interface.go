package factory

import (
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/matrix"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/pathfinding/shared"
)

type FinderInterface[T any] interface {
	OutOfBounds(pos matrix.Position) bool
	GetItemAtPosition(pos matrix.Position) T
	GetPositionInDirection(position matrix.Position, direction shared.Direction, moves int) matrix.Position
	EqualityCheck(a, b T) bool
}

type PathfindingRuleInterface[T any] interface {
	MatchFunc(currentPosition matrix.Position, currentDirection shared.Direction, finder FinderInterface[T]) (bool, error)
	GetNewPosition(currentPosition matrix.Position, currentDirection shared.Direction, finder FinderInterface[T]) matrix.Position
	GetNewDirection(currentPosition matrix.Position, currentDirection shared.Direction) shared.Direction
}
