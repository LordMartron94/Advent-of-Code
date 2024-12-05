package task_rules

import (
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/rules/factory"
	shared2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
)

type ParsingRuleFactory struct {
	factory *factory.ParsingRuleFactory[LexingTokenType]
}

func NewParsingRuleFactory() *ParsingRuleFactory {
	return &ParsingRuleFactory{
		factory: factory.NewParsingRuleFactory[LexingTokenType](),
	}
}

func (f *ParsingRuleFactory) GetPairParsingRule() rules.ParsingRuleInterface[LexingTokenType] {
	return f.factory.NewParsingRule("Pair", func(tokens []*shared.Token[LexingTokenType], tokenIndex int) (bool, string) {
		if tokenIndex+4 >= len(tokens) {
			return false, "expected at least three tokens to form a pair"
		}

		if tokens[tokenIndex].Type != DigitToken || tokens[tokenIndex+4].Type != DigitToken {
			return false, "expected two digits in a pair"
		}

		return true, ""
	}, func(tokens []*shared.Token[LexingTokenType], tokenIndex int) *shared2.ParseTree[LexingTokenType] {
		leftNumber := &shared2.ParseTree[LexingTokenType]{
			Symbol:   "left_number",
			Token:    tokens[tokenIndex],
			Children: nil,
		}

		rightNumber := &shared2.ParseTree[LexingTokenType]{
			Symbol:   "right_number",
			Token:    tokens[tokenIndex+4],
			Children: nil,
		}

		children := []*shared2.ParseTree[LexingTokenType]{
			leftNumber,
			rightNumber,
		}

		return &shared2.ParseTree[LexingTokenType]{
			Symbol:   "Pair",
			Token:    nil,
			Children: children,
		}
	}, true, 3)
}

func (f *ParsingRuleFactory) GetInvalidTokenParsingRule() rules.ParsingRuleInterface[LexingTokenType] {
	return f.factory.NewMatchAnyTokenParsingRule("invalid_token")
}
