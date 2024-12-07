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

func (p *ParsingRuleFactory) getTestNumberRule() rules.ParsingRuleInterface[LexingTokenType] {
	return p.factory.NewMatchUntilTokenWithFilterParsingRule("test_number", []LexingTokenType{NumberToken, ColonToken}, []string{"number", "colon"})
}

func (p *ParsingRuleFactory) getEquationPartsRule() rules.ParsingRuleInterface[LexingTokenType] {
	return p.factory.NewMatchUntilTokenWithFilterParsingRule("equation_parts", []LexingTokenType{NumberToken, WhitespaceToken}, []string{"number", "whitespace"})
}

func (p *ParsingRuleFactory) GetEquationRule() rules.ParsingRuleInterface[LexingTokenType] {
	testNumberRule := p.getTestNumberRule()
	equationPartsRule := p.getEquationPartsRule()

	return p.factory.NewNestedParsingRule("equation", []rules.ParsingRuleInterface[LexingTokenType]{
		testNumberRule,
		equationPartsRule,
	})
}

func (p *ParsingRuleFactory) GetInvalidTokenRule() rules.ParsingRuleInterface[LexingTokenType] {
	return p.factory.NewMatchAnyTokenParsingRule("invalid")
}
