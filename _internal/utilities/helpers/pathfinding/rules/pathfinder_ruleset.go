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

func (prs *PathfindingRuleset[T]) GetRule(currentPosition matrix.Position, currentDirection shared.Direction, finder factory.FinderInterface[T]) (factory.PathfindingRuleInterface[T], error) {
	if rule, ok := prs.ruleCache[currentDirection]; ok {
		return rule, nil
	}

	for _, rule := range prs.Rules {
		match, err := rule.MatchFunc(currentPosition, currentDirection, finder)

		if err != nil {
			return nil, err
		}

		if match {
			return rule, nil
		}
	}

	return nil, fmt.Errorf("no matching rule found")
}
