package task_rules

import (
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/rules/factory"
)

type LexingTokenType int

const (
	IgnoreToken LexingTokenType = iota
	NumberToken
	WhitespaceToken
)

type Ruleset struct {
	factory factory.RuleFactory[LexingTokenType]
}

func NewRuleset() *Ruleset {
	return &Ruleset{
		factory: factory.RuleFactory[LexingTokenType]{},
	}
}

func (r *Ruleset) GetInvalidTokenRuleLex() rules.LexingRuleInterface[LexingTokenType] {
	return r.factory.NewMatchAnyTokenRule(IgnoreToken)
}

func (r *Ruleset) GetDigitTokenRuleLex() rules.LexingRuleInterface[LexingTokenType] {
	return r.factory.NewNumberLexingRule(NumberToken, "digit")
}

func (r *Ruleset) GetWhitespaceTokenRuleLex() rules.LexingRuleInterface[LexingTokenType] {
	return r.factory.NewWhitespaceLexingRule(IgnoreToken, "whitespace")
}
