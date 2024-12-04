package task_rules

import (
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/rules/factory"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/scanning"
)

type LexingTokenType int

const (
	InvalidToken LexingTokenType = iota
	IgnoreToken
	NewLineToken
	XCharToken
	MCharToken
	ACharToken
	SCharToken
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
	return r.factory.NewInvalidTokenLexingRule(IgnoreToken)
}

func (r *Ruleset) GetXCharRuleLex() rules.LexingRuleInterface[LexingTokenType] {
	return r.factory.NewCharacterLexingRule('X', XCharToken, "XCharacterRuleLexer")
}

func (r *Ruleset) GetMCharRuleLex() rules.LexingRuleInterface[LexingTokenType] {
	return r.factory.NewCharacterLexingRule('M', MCharToken, "MCharacterRuleLexer")
}

func (r *Ruleset) GetACharRuleLex() rules.LexingRuleInterface[LexingTokenType] {
	return r.factory.NewCharacterLexingRule('A', ACharToken, "ACharacterRuleLexer")
}

func (r *Ruleset) GetSCharRuleLex() rules.LexingRuleInterface[LexingTokenType] {
	return r.factory.NewCharacterLexingRule('S', SCharToken, "SCharacterRuleLexer")
}

func (r *Ruleset) GetNewLineRuleLexer() rules.LexingRuleInterface[LexingTokenType] {
	return r.factory.NewLexingRule("NewLineRuleLexer", func(peekInterface scanning.PeekInterface) bool {
		currentRune := peekInterface.Current()

		return currentRune == '\n' || currentRune == '\r' || currentRune == ' '
	}, NewLineToken, func(scanner scanning.PeekInterface) []rune {
		runes := make([]rune, 0)
		runes = append(runes, scanner.Current())

		peekIndex := 1
		for {
			pRunes, err := scanner.Peek(peekIndex)

			if err != nil {
				break
			}

			peekedRune := pRunes[len(pRunes)-1]

			if peekedRune == '\n' || peekedRune == '\r' || peekedRune == ' ' {
				runes = append(runes, peekedRune)
				peekIndex++
			} else {
				break
			}
		}

		ignoreRune := '-'

		// Replace each rune in the runes slice with the ignoreRune for consistency
		for i := range runes {
			runes[i] = ignoreRune
		}

		return runes
	})
}
