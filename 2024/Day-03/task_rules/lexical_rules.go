package task_rules

import (
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/rules/factory"
)

type LexingTokenType int

const (
	InvalidToken LexingTokenType = iota
	MulKeywordToken
	OpenParenthesisToken
	CloseParenthesisToken
	CommaToken
	NumberToken
	DoKeywordToken
	DontKeywordToken
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
	return r.factory.NewInvalidTokenLexingRule(InvalidToken)
}

func (r *Ruleset) GetDigitRuleLex() rules.LexingRuleInterface[LexingTokenType] {
	return r.factory.NewNumberLexingRule(NumberToken, "DigitRuleLexer")
}

func (r *Ruleset) GetCommaRuleLex() rules.LexingRuleInterface[LexingTokenType] {
	return r.factory.NewCharacterLexingRule(',', CommaToken, "CommaRuleLex")
}

func (r *Ruleset) GetOpenParenthesisRuleLex() rules.LexingRuleInterface[LexingTokenType] {
	return r.factory.NewCharacterLexingRule('(', OpenParenthesisToken, "OpenParenthesisRuleLex")
}

func (r *Ruleset) GetCloseParenthesisRuleLex() rules.LexingRuleInterface[LexingTokenType] {
	return r.factory.NewCharacterLexingRule(')', CloseParenthesisToken, "CloseParenthesisRuleLex")
}

func (r *Ruleset) GetDoKeywordRuleLex() rules.LexingRuleInterface[LexingTokenType] {
	return r.factory.NewKeywordLexingRule("do", DoKeywordToken, "DoKeywordRuleLex")
}

func (r *Ruleset) GetDontKeywordRuleLex() rules.LexingRuleInterface[LexingTokenType] {
	return r.factory.NewKeywordLexingRule("don't", DontKeywordToken, "DontKeywordRuleLex")
}

func (r *Ruleset) GetMulKeywordRuleLex() rules.LexingRuleInterface[LexingTokenType] {
	return r.factory.NewKeywordLexingRule("mul", MulKeywordToken, "MulKeywordRuleLex")
}
