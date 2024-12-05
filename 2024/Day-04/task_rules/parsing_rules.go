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

func (p *ParsingRuleFactory) GetHorizontalLineParserRule() rules.ParsingRuleInterface[LexingTokenType] {
	return p.factory.NewMatchUntilTokenWithFilterParsingRule("horizontal_line", []LexingTokenType{XCharToken, MCharToken, ACharToken, SCharToken}, []string{"horizontal_line_element", "horizontal_line_element", "horizontal_line_element", "horizontal_line_element"})
}

func (p *ParsingRuleFactory) GetNewLineTokenParserRule() rules.ParsingRuleInterface[LexingTokenType] {
	return p.factory.NewSingleTokenParsingRule("new_line", NewLineToken)
}

func (p *ParsingRuleFactory) GetInvalidTokenRule() rules.ParsingRuleInterface[LexingTokenType] {
	return p.factory.NewMatchAnyTokenParsingRule("invalid")
}
