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

func (p *ParsingRuleFactory) GetEmptySpotRule() rules.ParsingRuleInterface[LexingTokenType] {
	return p.factory.NewSingleTokenParsingRule("empty_spot", DotToken)
}

func (p *ParsingRuleFactory) GetAntennaRule() rules.ParsingRuleInterface[LexingTokenType] {
	return p.factory.NewSingleTokenParsingRule("antenna", AlphanumericToken)
}

func (p *ParsingRuleFactory) GetRowParsingRule() rules.ParsingRuleInterface[LexingTokenType] {
	return p.factory.NewOptionalNestedParsingRule("row", []rules.ParsingRuleInterface[LexingTokenType]{
		p.GetEmptySpotRule(),
		p.GetAntennaRule(),
	})
}

func (p *ParsingRuleFactory) GetInvalidTokenRule() rules.ParsingRuleInterface[LexingTokenType] {
	return p.factory.NewMatchAnyTokenParsingRule("invalid")
}
