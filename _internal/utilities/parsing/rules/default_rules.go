package rules

import (
	"fmt"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
	shared2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
)

type WhitespaceRule struct{}

func (w *WhitespaceRule) Symbol() string {
	return "whitespace"
}

func (w *WhitespaceRule) Match(tokens []*shared.Token, currentIndex int) (*shared2.ParseTree, error) {
	currentToken := tokens[currentIndex]

	if currentToken.Type != shared.WhitespaceToken {
		return nil, fmt.Errorf("expected whitespace token, got %v", currentToken.Type)
	}

	return nil, nil // Whitespace is not used.
}

type NumberRule struct{}

func (n *NumberRule) Symbol() string {
	return "number"
}

func (n *NumberRule) Match(tokens []*shared.Token, currentIndex int) (*shared2.ParseTree, error) {
	currentToken := tokens[currentIndex]

	if currentToken.Type != shared.NumberToken {
		return nil, fmt.Errorf("expected number token, got %v", currentToken.Type)
	}

	return &shared2.ParseTree{
		Symbol:   n.Symbol(),
		Token:    currentToken,
		Children: make([]*shared2.ParseTree, 0),
	}, nil
}

type PairRule struct{}

func (p *PairRule) Symbol() string {
	return "pair"
}

func (p *PairRule) Match(tokens []*shared.Token, currentIndex int) (*shared2.ParseTree, error) {
	if currentIndex+2 >= len(tokens) {
		return nil, fmt.Errorf("expected at least three tokens to form a pair")
	}

	firstChild, err := (&NumberRule{}).Match(tokens, currentIndex)
	if err != nil {
		return nil, err
	}
	firstChild.Symbol = "first_number"

	secondChild, err := (&NumberRule{}).Match(tokens, currentIndex+2)
	if err != nil {
		return nil, err
	}
	secondChild.Symbol = "second_number"

	children := []*shared2.ParseTree{
		firstChild,
		secondChild,
	}

	return &shared2.ParseTree{
		Symbol:   p.Symbol(),
		Token:    nil,
		Children: children,
	}, nil
}
