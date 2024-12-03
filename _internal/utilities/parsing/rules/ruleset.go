package rules

import (
	"fmt"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
)

type Ruleset[T comparable] struct {
	Rules []ParsingRuleInterface[T]
}

func NewRuleset[T comparable](rules []ParsingRuleInterface[T]) *Ruleset[T] {
	return &Ruleset[T]{Rules: rules}
}

func (rs *Ruleset[T]) GetMatchingRule(input []*shared.Token[T], currentIndex int) (ParsingRuleInterface[T], error) {
	for _, rule := range rs.Rules {
		_, err, _ := rule.Match(input, currentIndex)

		if err == nil {
			//fmt.Println(fmt.Sprintf("Matched rule (ruleSet Matcher): %s for input '%s'", rule.Symbol(), input[currentIndex].Value))
			return rule, nil
		}
	}

	return nil, fmt.Errorf("no matching rule found for input '%c'\n", input)
}
