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

func (p *ParsingRuleFactory) GetNumberTokenRule() rules.ParsingRuleInterface[LexingTokenType] {
	return p.factory.NewSingleTokenParsingRule("number", NumberToken)
}

func (p *ParsingRuleFactory) GetRowParsingRule() rules.ParsingRuleInterface[LexingTokenType] {
	return p.factory.NewOptionalNestedParsingRule("row", []rules.ParsingRuleInterface[LexingTokenType]{
		p.GetNumberTokenRule(),
	})
}

func (p *ParsingRuleFactory) GetInvalidTokenRule() rules.ParsingRuleInterface[LexingTokenType] {
	return p.factory.NewMatchAnyTokenParsingRule("invalid")
}
