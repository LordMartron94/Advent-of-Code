package task_rules

import (
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/rules/factory"
)

type LexingTokenType int

const (
	IgnoreToken LexingTokenType = iota
	WhitespaceToken
	DotToken
	HashToken
	CarrotToken
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

func (r *Ruleset) GetWhitespaceTokenRuleLex() rules.LexingRuleInterface[LexingTokenType] {
	return r.factory.NewWhitespaceLexingRule(WhitespaceToken, "whitespace")
}

func (r *Ruleset) GetDotTokenRuleLex() rules.LexingRuleInterface[LexingTokenType] {
	return r.factory.NewCharacterLexingRule('.', DotToken, "dot")
}

func (r *Ruleset) GetHashTokenRuleLex() rules.LexingRuleInterface[LexingTokenType] {
	return r.factory.NewCharacterLexingRule('#', HashToken, "hash")
}

func (r *Ruleset) GetCarrotTokenRuleLex() rules.LexingRuleInterface[LexingTokenType] {
	return r.factory.NewCharacterLexingRule('^', CarrotToken, "carrot")
}
