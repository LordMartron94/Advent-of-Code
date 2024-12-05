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

func (p *ParsingRuleFactory) GetPairOfNumbersParserRule() rules.ParsingRuleInterface[LexingTokenType] {
	return p.factory.NewSequentialTokenParsingRule("pair_of_numbers", []LexingTokenType{DigitToken, PipeToken, DigitToken}, []string{"first_number", "pipe", "second_number"})
}

func (p *ParsingRuleFactory) GetUpdateCommandParserRule() rules.ParsingRuleInterface[LexingTokenType] {
	return p.factory.NewMatchUntilTokenWithFilterParsingRule("update_command", []LexingTokenType{DigitToken, CommaToken}, []string{"number", "comma"})
}

func (p *ParsingRuleFactory) GetInvalidTokenRule() rules.ParsingRuleInterface[LexingTokenType] {
	return p.factory.NewMatchAnyTokenParsingRule("invalid")
}

func (p *ParsingRuleFactory) GetWhitespaceTokenParserRule() rules.ParsingRuleInterface[LexingTokenType] {
	return p.factory.NewSingleTokenParsingRule("whitespace", WhitespaceToken)
}
