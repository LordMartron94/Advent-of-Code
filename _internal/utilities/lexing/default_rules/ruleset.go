package default_rules

import (
	"fmt"
)

type Ruleset struct {
	Rules []LexingRule
}

func NewRuleset(rules []LexingRule) *Ruleset {
	return &Ruleset{Rules: rules}
}

func (rs *Ruleset) GetMatchingRule(input rune) (LexingRule, error) {
	for _, rule := range rs.Rules {
		if rule.Match(input) {
			//fmt.Println(fmt.Sprintf("Matched rule (ruleSet Matcher): %s for input '%s'", rule.GetName(), string(input)))
			return rule, nil
		}
	}

	return nil, fmt.Errorf("no matching rule found for input '%c'\n", input)
}
