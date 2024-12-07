package factory

import (
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/matrix"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/pathfinding"
)

type PathfindingRuleFactory[T any] struct {
}

func NewPathfindingRuleFactory[T any]() *PathfindingRuleFactory[T] {
	return &PathfindingRuleFactory[T]{}
}

// GetRule returns a new pathfinding rule.
func (p *PathfindingRuleFactory[T]) GetRule(
	matchFunc func(currentPosition matrix.Position, currentDirection pathfinding.Direction, finder pathfinding.PathFinder[T], getNextItemFunc func(matrix.Position) (*T, error)) (bool, error),
	getNewDirectionFunc func(currentPosition matrix.Position, currentDirection pathfinding.Direction) pathfinding.Direction,
	getNewPositionFunc func(currentPosition matrix.Position, newDirection pathfinding.Direction, finder pathfinding.PathFinder[T]) matrix.Position) pathfinding.PathfindingRuleInterface[T] {
	return &BasePathfindingRule[T]{getMatch: matchFunc, getNewDirection: getNewDirectionFunc, getNewPosition: getNewPositionFunc}
}

// GetBasicRule returns a new basic pathfinding rule.
func (p *PathfindingRuleFactory[T]) GetBasicRule(nextTileCondition func(finder pathfinding.PathFinder[T], nextTile T) bool, nextDirectionFunc func(currentDirection pathfinding.Direction) pathfinding.Direction, amountOfMoves int) pathfinding.PathfindingRuleInterface[T] {
	return p.GetRule(func(currentPosition matrix.Position, currentDirection pathfinding.Direction, finder pathfinding.PathFinder[T], getNextItemFunc func(matrix.Position) (*T, error)) (bool, error) {
		pos := finder.GetPositionInDirection(currentPosition, currentDirection, amountOfMoves)

		item, err := getNextItemFunc(pos)

		if err != nil {
			return false, err
		}

		if nextTileCondition(finder, *item) {
			return true, nil
		}

		return false, nil
	}, func(currentPosition matrix.Position, currentDirection pathfinding.Direction) pathfinding.Direction {
		return nextDirectionFunc(currentDirection)
	}, func(pos matrix.Position, direction pathfinding.Direction, finder pathfinding.PathFinder[T]) matrix.Position {
		return finder.GetPositionInDirection(pos, direction, amountOfMoves)
	})
}
