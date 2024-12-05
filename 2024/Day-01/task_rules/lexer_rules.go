package task_rules

import (
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/rules/factory"
)

type LexingTokenType int

const (
	IgnoreToken LexingTokenType = iota
	WhitespaceToken
	DigitToken
)

type RuleFactory struct {
	factory *factory.RuleFactory[LexingTokenType]
}

func NewRuleFactory() *RuleFactory {
	return &RuleFactory{
		factory: &factory.RuleFactory[LexingTokenType]{},
	}
}

func (f *RuleFactory) GetInvalidTokenRuleLex() rules.LexingRuleInterface[LexingTokenType] {
	return f.factory.NewMatchAnyTokenRule(IgnoreToken)
}

func (f *RuleFactory) GetWhitespaceRuleLex() rules.LexingRuleInterface[LexingTokenType] {
	return f.factory.NewCharacterLexingRule(' ', WhitespaceToken, "WhitespaceRuleLexer")
}

func (f *RuleFactory) GetDigitRuleLex() rules.LexingRuleInterface[LexingTokenType] {
	return f.factory.NewNumberLexingRule(DigitToken, "DigitRuleLexer")
}
