package default_rules

import (
	"fmt"
)

type Ruleset[T any] struct {
	Rules []LexingRuleInterface[T]
}

func NewRuleset[T any](rules []LexingRuleInterface[T]) *Ruleset[T] {
	return &Ruleset[T]{Rules: rules}
}

func (rs *Ruleset[T]) GetMatchingRule(input rune, peeker LexerInterface) (LexingRuleInterface[T], error) {
	for _, rule := range rs.Rules {
		if rule.Match(input, peeker) {
			//fmt.Println(fmt.Sprintf("Matched rule (ruleSet Matcher): %s for input '%s'", rule.GetName(), string(input)))
			return rule, nil
		}
	}

	return nil, fmt.Errorf("no matching rule found for input '%c'\n", input)
}
