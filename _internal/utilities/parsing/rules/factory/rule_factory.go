package factory

import (
	"fmt"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/rules"
	shared2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
)

func sliceContains[T comparable](s []T, e T) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

type ParsingRuleFactory[T comparable] struct {
}

func NewParsingRuleFactory[T comparable]() *ParsingRuleFactory[T] {
	return &ParsingRuleFactory[T]{}
}

// NewParsingRule returns a new ParsingRule instance.
func (p *ParsingRuleFactory[T]) NewParsingRule(
	symbol string,
	matchFunc func([]*shared.Token[T], int) (bool, string),
	getContentFunc func([]*shared.Token[T], int) *shared2.ParseTree[T],
	isShell bool,
	consumeExtra ...int) rules.ParsingRuleInterface[T] {
	if len(consumeExtra) < 1 {
		consumeExtra = append(consumeExtra, 0)
	}
	if len(consumeExtra) > 1 {
		panic("consumeExtra must have a single element")
	}

	return &BaseParsingRule[T]{
		SymbolString:   symbol,
		matchFunc:      matchFunc,
		getContentFunc: getContentFunc,
		isShell:        isShell,
		consumeExtra:   consumeExtra[0],
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
			Symbol:   symbol,
			Token:    tokens[index],
			Children: nil,
		}
	}, false)
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
				Symbol:   childSymbols[i],
				Token:    tokens[index+i],
				Children: nil,
			}
		}

		return &shared2.ParseTree[T]{
			Symbol:   symbol,
			Token:    nil,
			Children: children,
		}
	}, true)
}

// NewMatchUntilTokenParsingRule returns a new ParsingRule instance that matches until a specific token is encountered.
func (p *ParsingRuleFactory[T]) NewMatchUntilTokenParsingRule(symbol string, targetTokenType T, childSymbol string) rules.ParsingRuleInterface[T] {
	return p.NewParsingRule(symbol, func(tokens []*shared.Token[T], index int) (bool, string) {
		childCount := 0

		for i := index; i < len(tokens); i++ {
			if tokens[i].Type == targetTokenType {
				break
			}

			childCount++
		}

		if childCount == 0 {
			return false, fmt.Sprintf("expected at least 1 token to form a %s", symbol)
		}

		return true, ""
	}, func(tokens []*shared.Token[T], index int) *shared2.ParseTree[T] {
		children := make([]*shared2.ParseTree[T], 0)

		for i := index; i < len(tokens); i++ {
			if tokens[i].Type == targetTokenType {
				break
			}

			children = append(children, &shared2.ParseTree[T]{
				Symbol:   childSymbol,
				Token:    tokens[i],
				Children: nil,
			})
		}

		return &shared2.ParseTree[T]{
			Symbol:   symbol,
			Token:    nil,
			Children: children,
		}
	}, true)
}

// NewMatchUntilTokenWithFilterParsingRule returns a new ParsingRule instance that matches as long as tokens of the specified types are encountered.
func (p *ParsingRuleFactory[T]) NewMatchUntilTokenWithFilterParsingRule(symbol string, possibleChildrenTypes []T, childSymbol string) rules.ParsingRuleInterface[T] {
	return p.NewParsingRule(symbol, func(tokens []*shared.Token[T], index int) (bool, string) {
		childCount := 0

		for i := index; i < len(tokens); i++ {
			if !sliceContains(possibleChildrenTypes, tokens[i].Type) {
				break
			}

			childCount++
		}

		if childCount == 0 {
			return false, fmt.Sprintf("expected at least 1 token to form a %s", symbol)
		}

		return true, ""
	}, func(tokens []*shared.Token[T], index int) *shared2.ParseTree[T] {
		children := make([]*shared2.ParseTree[T], 0)

		for i := index; i < len(tokens); i++ {
			if !sliceContains(possibleChildrenTypes, tokens[i].Type) {
				break
			}

			children = append(children, &shared2.ParseTree[T]{
				Symbol:   childSymbol,
				Token:    tokens[i],
				Children: nil,
			})
		}

		return &shared2.ParseTree[T]{
			Symbol:   symbol,
			Token:    nil,
			Children: children,
		}
	}, true)
}

// NewMatchExceptParsingRule returns a new ParsingRule instance that matches tokens except for the specified type.
// Very useful for default parsing rules such as invalid tokens.
func (p *ParsingRuleFactory[T]) NewMatchExceptParsingRule(symbol string, excludeTokenType T) rules.ParsingRuleInterface[T] {
	return p.NewParsingRule(symbol, func(tokens []*shared.Token[T], index int) (bool, string) {
		if tokens[index].Type == excludeTokenType {
			return false, fmt.Sprintf("unexpected %s token", excludeTokenType)
		}

		return true, ""
	}, func(tokens []*shared.Token[T], index int) *shared2.ParseTree[T] {
		return &shared2.ParseTree[T]{
			Symbol:   symbol,
			Token:    tokens[index],
			Children: nil,
		}
	}, false)
}

// NewMatchAnyTokenParsingRule returns a new ParsingRule instance that matches any token.
// Very useful for default parsing rules such as invalid tokens.
func (p *ParsingRuleFactory[T]) NewMatchAnyTokenParsingRule(symbol string) rules.ParsingRuleInterface[T] {
	return p.NewParsingRule(symbol, func(tokens []*shared.Token[T], index int) (bool, string) {
		return true, ""
	}, func(tokens []*shared.Token[T], index int) *shared2.ParseTree[T] {
		return &shared2.ParseTree[T]{
			Symbol:   symbol,
			Token:    tokens[index],
			Children: nil,
		}
	}, false)
}
