package task_rules

import (
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/rules/factory"
)

type LexingTokenType int

const (
	IgnoreToken LexingTokenType = iota
	NumberToken
	NewLineToken
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

func (r *Ruleset) GetCharTokenRuleLex() rules.LexingRuleInterface[LexingTokenType] {
	return r.factory.NewAlphanumericCharacterLexingRuleSingle(NumberToken, "alphanumeric")
}

func (r *Ruleset) GetNewLineTokenRuleLex() rules.LexingRuleInterface[LexingTokenType] {
	return r.factory.NewCharacterOptionLexingRule([]rune{'\r', '\n'}, NewLineToken, "newline")
}
