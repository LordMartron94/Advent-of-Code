package factory

import (
	"slices"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/rules"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/scanning"
)

// RuleFactory provides an API for creating lexing rules.
type RuleFactory[T comparable] struct {
}

// NewLexingRule creates a new lexing rule with the given symbol, match function, and associated token.
func (r *RuleFactory[T]) NewLexingRule(symbol string, isMatchFunc func(scanning.PeekInterface) bool, associatedToken T, getContentFunc func(peekInterface scanning.PeekInterface) []rune) rules.LexingRuleInterface[T] {
	rule := &baseLexingRule[T]{
		SymbolString:    symbol,
		MatchFunc:       isMatchFunc,
		AssociatedToken: associatedToken,
		GetContentFunc:  getContentFunc,
	}

	return rule
}

// NewMatchAnyTokenRule creates a new default lexing rule with the given associated token.
func (r *RuleFactory[T]) NewMatchAnyTokenRule(unknownToken T) rules.LexingRuleInterface[T] {
	return &baseLexingRule[T]{
		SymbolString:    "UnknownTokenRuleLexer",
		MatchFunc:       func(scanning.PeekInterface) bool { return true },
		AssociatedToken: unknownToken,
		GetContentFunc: func(scanner scanning.PeekInterface) []rune {
			return []rune{scanner.Current()}
		},
	}
}

func (r *RuleFactory[T]) NewIgnoreTokenLexingRule(symbol string, associatedToken T, matchFunc func(scanning.PeekInterface) bool) rules.LexingRuleInterface[T] {
	return &baseLexingRule[T]{
		SymbolString:    symbol,
		MatchFunc:       matchFunc,
		AssociatedToken: associatedToken,
		GetContentFunc: func(scanner scanning.PeekInterface) []rune {
			return []rune("IGNORE")
		},
	}
}

// NewKeywordLexingRule creates a new lexing rule with the given keyword, associated token, and symbol.
func (r *RuleFactory[T]) NewKeywordLexingRule(keyword string, associatedToken T, symbol string) rules.LexingRuleInterface[T] {
	return &baseLexingRule[T]{
		SymbolString: symbol,
		MatchFunc: func(scanner scanning.PeekInterface) bool {
			runesInKeyword := []rune(keyword)
			currentRune := scanner.Current()

			if currentRune != runesInKeyword[0] {
				return false
			}

			peekedRunes, err := scanner.Peek(len(runesInKeyword) - 2)

			if err != nil {
				return false
			}

			for i, r := range peekedRunes {
				if r != runesInKeyword[i+1] {
					return false
				}
			}

			return true
		},
		AssociatedToken: associatedToken,
		GetContentFunc: func(scanner scanning.PeekInterface) []rune {
			return []rune(keyword)
		},
	}
}

func (r *RuleFactory[T]) NewCharacterLexingRule(character rune, associatedToken T, symbol string) rules.LexingRuleInterface[T] {
	return &baseLexingRule[T]{
		SymbolString: symbol,
		MatchFunc: func(scanner scanning.PeekInterface) bool {
			currentRune := scanner.Current()

			if currentRune != character {
				return false
			}

			return true
		},
		AssociatedToken: associatedToken,
		GetContentFunc: func(scanner scanning.PeekInterface) []rune {
			return []rune{character}
		},
	}
}

func (r *RuleFactory[T]) NewCharacterOptionLexingRule(characters []rune, associatedToken T, symbol string) rules.LexingRuleInterface[T] {
	return &baseLexingRule[T]{
		SymbolString: symbol,
		MatchFunc: func(scanner scanning.PeekInterface) bool {
			currentRune := scanner.Current()

			if !slices.Contains(characters, currentRune) {
				return false
			}

			return true
		},
		AssociatedToken: associatedToken,
		GetContentFunc: func(scanner scanning.PeekInterface) []rune {
			return []rune{scanner.Current()}
		},
	}
}

func (r *RuleFactory[T]) NewNumberLexingRule(associatedToken T, symbol string) rules.LexingRuleInterface[T] {
	vFunc := func(r rune) bool {
		return '0' <= r && r <= '9'
	}
	mFunc := func(scanner scanning.PeekInterface) bool {
		currentRune := scanner.Current()

		if !vFunc(currentRune) {
			return false
		}

		return true
	}

	return &baseLexingRule[T]{
		SymbolString:    symbol,
		MatchFunc:       mFunc,
		AssociatedToken: associatedToken,
		GetContentFunc: func(scanner scanning.PeekInterface) []rune {
			runes := make([]rune, 0)
			runes = append(runes, scanner.Current())

			peekIndex := 1
			for {
				pRunes, err := scanner.Peek(peekIndex)

				if err != nil {
					break
				}

				peekedRune := pRunes[len(pRunes)-1]

				if vFunc(peekedRune) {
					runes = append(runes, peekedRune)
					peekIndex++
				} else {
					break
				}
			}

			return runes
		},
	}
}

func (r *RuleFactory[T]) NewWhitespaceLexingRule(tokenType T, symbol string) rules.LexingRuleInterface[T] {
	vFunc := func(r rune) bool {
		return r == '\t' || r == '\n' || r == '\r' || r == ' ' || r == '\f' || r == '\v'
	}
	mFunc := func(scanner scanning.PeekInterface) bool {
		currentRune := scanner.Current()

		if !vFunc(currentRune) {
			return false
		}

		return true
	}

	return &baseLexingRule[T]{
		SymbolString:    symbol,
		MatchFunc:       mFunc,
		AssociatedToken: tokenType,
		GetContentFunc: func(scanner scanning.PeekInterface) []rune {
			runes := make([]rune, 0)
			runes = append(runes, scanner.Current())

			peekIndex := 1
			for {
				pRunes, err := scanner.Peek(peekIndex)

				if err != nil {
					break
				}

				peekedRune := pRunes[len(pRunes)-1]

				if vFunc(peekedRune) {
					runes = append(runes, peekedRune)
					peekIndex++
				} else {
					break
				}
			}

			return runes
		},
	}
}
