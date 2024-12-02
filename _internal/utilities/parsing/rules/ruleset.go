package rules

import (
	"fmt"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
)

type Ruleset struct {
	Rules []ParsingRuleInterface
}

func NewRuleset(rules []ParsingRuleInterface) *Ruleset {
	return &Ruleset{Rules: rules}
}

func (rs *Ruleset) GetMatchingRule(input []*shared.Token, currentIndex int) (ParsingRuleInterface, error) {
	for _, rule := range rs.Rules {
		_, err := rule.Match(input, currentIndex)

		if err == nil {
			//fmt.Println(fmt.Sprintf("Matched rule (ruleSet Matcher): %s for input '%s'", rule.GetName(), string(input)))
			return rule, nil
		}
	}

	return nil, fmt.Errorf("no matching rule found for input '%c'\n", input)
}
