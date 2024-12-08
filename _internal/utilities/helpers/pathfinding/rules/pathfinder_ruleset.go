package rules

import (
	"fmt"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/matrix"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/pathfinding/rules/factory"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/helpers/pathfinding/shared"
)

type PathfindingRuleset[T any] struct {
	IsBasic   bool // Set to true if the ruleset is simple (follow current direction until obstacle, if obstacle, do x)... Two rules.
	Rules     []factory.PathfindingRuleInterface[T]
	ruleCache map[shared.Direction]factory.PathfindingRuleInterface[T]
}

func NewPathfindingRuleset[T any](rules []factory.PathfindingRuleInterface[T], isBasic bool) *PathfindingRuleset[T] {
	return &PathfindingRuleset[T]{
		IsBasic:   isBasic,
		Rules:     rules,
		ruleCache: make(map[shared.Direction]factory.PathfindingRuleInterface[T]),
	}
}

func (prs *PathfindingRuleset[T]) GetRule(currentPosition matrix.Position, currentDirection shared.Direction, finder factory.FinderInterface[T], lastTile T) (factory.PathfindingRuleInterface[T], int, error) {
	nextTiles := finder.GetTilesInDirection(currentPosition, currentDirection)

	if len(nextTiles) == 0 {
		return nil, 0, &shared.OutOfBoundsError{}
	}

	// TODO - Fix Cacheing
	// TODO - Implement fast-forward based on match amount output
	//if rule, ok := prs.ruleCache[currentDirection]; ok {
	//	return rule, 0, nil
	//}

	for _, rule := range prs.Rules {
		amount := rule.MatchFunc(finder, nextTiles)

		if amount > 0 {
			if !rule.GetDirectionNeedsPosition() {
				prs.ruleCache[currentDirection] = rule
			}

			return rule, amount, nil
		}
	}

	return nil, 0, fmt.Errorf("no matching rule found")
}
