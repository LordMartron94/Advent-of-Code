package task_rules

import (
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/rules/factory"
)

type ParsingRuleFactory struct {
	factory *factory.ParsingRuleFactory[LexingTokenType]
}

func NewParsingRuleFactory() *ParsingRuleFactory {
	return &ParsingRuleFactory{
		factory: factory.NewParsingRuleFactory[LexingTokenType](),
	}
}

func (p *ParsingRuleFactory) GetRowRule() rules.ParsingRuleInterface[LexingTokenType] {
	return p.factory.NewMatchUntilTokenWithFilterParsingRule("row", []LexingTokenType{DotToken, HashToken, CarrotToken}, []string{"dot", "hash", "carrot"})
}

func (p *ParsingRuleFactory) GetInvalidTokenRule() rules.ParsingRuleInterface[LexingTokenType] {
	return p.factory.NewMatchAnyTokenParsingRule("invalid")
}

func (p *ParsingRuleFactory) GetWhitespaceTokenParserRule() rules.ParsingRuleInterface[LexingTokenType] {
	return p.factory.NewSingleTokenParsingRule("whitespace", WhitespaceToken)
}
