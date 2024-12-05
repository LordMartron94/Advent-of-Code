package factory

import (
	"fmt"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/rules"
	shared2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
)

type ParsingRuleFactory[T comparable] struct {
}

// NewParsingRule returns a new ParsingRule instance.
func (p *ParsingRuleFactory[T]) NewParsingRule(symbol string, matchFunc func([]*shared.Token[T], int) (bool, string), getContentFunc func([]*shared.Token[T], int) *shared2.ParseTree[T]) rules.ParsingRuleInterface[T] {
	return &BaseParsingRule[T]{
		SymbolString:   symbol,
		matchFunc:      matchFunc,
		getContentFunc: getContentFunc,
	}
}

// NewSingleTokenParsingRule returns a new ParsingRule instance that matches a single token of the specified type.
func (p *ParsingRuleFactory[T]) NewSingleTokenParsingRule(symbol string, associatedTokenType T) rules.ParsingRuleInterface[T] {
	return p.NewParsingRule(symbol, func(tokens []*shared.Token[T], index int) (bool, string) {
		if index < len(tokens) && tokens[index].Type == associatedTokenType {
			return true, ""
		}
		return false, fmt.Sprintf("expected %s token", associatedTokenType)
	}, func(tokens []*shared.Token[T], index int) *shared2.ParseTree[T] {
		return &shared2.ParseTree[T]{
			Symbol: symbol,
			Token:  tokens[index],
			Children: []*shared2.ParseTree[T]{
				{Token: nil, Children: nil},
			},
		}
	})
}

// NewSequentialTokenParsingRule returns a new ParsingRule instance that matches a sequence of tokens for the specified types.
func (p *ParsingRuleFactory[T]) NewSequentialTokenParsingRule(symbol string, targetTokenTypeSequence []T, childSymbols []string) rules.ParsingRuleInterface[T] {
	if len(targetTokenTypeSequence) != len(childSymbols) {
		panic("targetTokenTypeSequence and childSymbols must have the same length")
	}

	return p.NewParsingRule(symbol, func(tokens []*shared.Token[T], index int) (bool, string) {
		if len(tokens) < len(targetTokenTypeSequence) {
			return false, fmt.Sprintf("expected at least %d tokens", len(targetTokenTypeSequence))
		}

		for i := 0; i < len(targetTokenTypeSequence); i++ {
			if tokens[index+i].Type != targetTokenTypeSequence[i] {
				return false, fmt.Sprintf("expected %s token at position %d", targetTokenTypeSequence[i], i+1)
			}
		}

		return true, ""
	}, func(tokens []*shared.Token[T], index int) *shared2.ParseTree[T] {
		children := make([]*shared2.ParseTree[T], len(childSymbols))

		for i := 0; i < len(childSymbols); i++ {
			children[i] = &shared2.ParseTree[T]{
				Symbol: childSymbols[i],
				Token:  tokens[index+i],
				Children: []*shared2.ParseTree[T]{
					{Token: nil, Children: nil},
				},
			}
		}

		return &shared2.ParseTree[T]{
			Symbol:   symbol,
			Token:    tokens[index],
			Children: children,
		}
	})
}
