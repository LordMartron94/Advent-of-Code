package factory

import (
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/matrix"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/pathfinding/shared"
)

type PathfindingRuleFactory[T any] struct {
}

func NewPathfindingRuleFactory[T any]() *PathfindingRuleFactory[T] {
	return &PathfindingRuleFactory[T]{}
}

// GetRule returns a new pathfinding rule.
func (p *PathfindingRuleFactory[T]) GetRule(
	matchFunc func(currentPosition matrix.Position, currentDirection shared.Direction, finder FinderInterface[T], getNextItemFunc func(matrix.Position) (*T, error)) (bool, error),
	getNewDirectionFunc func(currentPosition matrix.Position, currentDirection shared.Direction) shared.Direction,
	getNewPositionFunc func(currentPosition matrix.Position, newDirection shared.Direction, finder FinderInterface[T]) matrix.Position,
	directionNeedsPosition bool,
) PathfindingRuleInterface[T] {
	return &BasePathfindingRule[T]{getMatch: matchFunc, getNewDirection: getNewDirectionFunc, getNewPosition: getNewPositionFunc, DirectionNeedsPosition: directionNeedsPosition}
}

// GetBasicRule returns a new basic pathfinding rule.
func (p *PathfindingRuleFactory[T]) GetBasicRule(nextTileCondition func(finder FinderInterface[T], nextTile T) bool, nextDirectionFunc func(currentDirection shared.Direction) shared.Direction, amountOfMoves int) PathfindingRuleInterface[T] {
	return p.GetRule(func(currentPosition matrix.Position, currentDirection shared.Direction, finder FinderInterface[T], getNextItemFunc func(matrix.Position) (*T, error)) (bool, error) {
		pos := finder.GetPositionInDirection(currentPosition, currentDirection, amountOfMoves)

		item, err := getNextItemFunc(pos)

		if err != nil {
			return false, err
		}

		if nextTileCondition(finder, *item) {
			return true, nil
		}

		return false, nil
	}, func(_ matrix.Position, currentDirection shared.Direction) shared.Direction {
		return nextDirectionFunc(currentDirection)
	}, func(pos matrix.Position, direction shared.Direction, finder FinderInterface[T]) matrix.Position {
		return finder.GetPositionInDirection(pos, direction, amountOfMoves)
	}, false)
}
