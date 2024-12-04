package rules

import (
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/scanning"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
)

type LexingRuleInterface[T any] interface {
	// Symbol returns the lexical symbol this rule represents (e.g., "expression", "statement", "term").
	Symbol() string

	// IsMatch checks if the given sequence of runes matches this rule's pattern.
	IsMatch(scanner scanning.PeekInterface) bool

	// ExtractToken extracts a token from the given sequence of runes that matches this rule's pattern.
	// If no match is found, it will return an error.
	// It will also return the amount of runes consumed by the extraction.
	ExtractToken(scanner scanning.PeekInterface) (*shared.Token[T], error, int)
}
