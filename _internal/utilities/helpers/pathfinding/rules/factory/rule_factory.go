package factory

import (
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/matrix"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/matrix/shared"
)

type PathfindingRuleFactory[T any] struct {
}

func NewPathfindingRuleFactory[T any]() *PathfindingRuleFactory[T] {
	return &PathfindingRuleFactory[T]{}
}

// GetRule returns a new pathfinding rule.
func (p *PathfindingRuleFactory[T]) GetRule(
	matchFunc func(finder FinderInterface[T], nextTiles []T) int,
	getNewDirectionFunc func(currentPosition matrix.Position, currentDirection shared.Direction) shared.Direction,
	getNewPositionFunc func(currentPosition matrix.Position, newDirection shared.Direction, finder FinderInterface[T]) matrix.Position,
	directionNeedsPosition bool,
) PathfindingRuleInterface[T] {
	return &BasePathfindingRule[T]{getMatch: matchFunc, getNewDirection: getNewDirectionFunc, getNewPosition: getNewPositionFunc, directionNeedsPosition: directionNeedsPosition}
}

// GetBasicRule returns a new basic pathfinding rule.
func (p *PathfindingRuleFactory[T]) GetBasicRule(nextTileCondition func(finder FinderInterface[T], nextTile T) bool, nextDirectionFunc func(currentDirection shared.Direction) shared.Direction, amountOfMoves int) PathfindingRuleInterface[T] {
	return p.GetRule(func(finder FinderInterface[T], nextTiles []T) int {
		if nextTileCondition(finder, nextTiles[0]) {
			return 1
		}

		return 0
	}, func(_ matrix.Position, currentDirection shared.Direction) shared.Direction {
		return nextDirectionFunc(currentDirection)
	}, func(pos matrix.Position, direction shared.Direction, finder FinderInterface[T]) matrix.Position {
		return finder.GetPositionInDirection(pos, direction, amountOfMoves)
	}, false)
}
