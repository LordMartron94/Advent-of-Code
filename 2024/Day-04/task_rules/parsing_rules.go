package task_rules

import (
	"fmt"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
	shared2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
)

type HorizontalLineParserRule struct{}

func (X *HorizontalLineParserRule) Symbol() string {
	return "HorizontalLineParserRule"
}

func sliceContains(s []LexingTokenType, e LexingTokenType) bool {
	for _, a := range s {
		if a == e {
			return true
		}
	}
	return false
}

func (X *HorizontalLineParserRule) Match(tokens []*shared.Token[LexingTokenType], currentIndex int) (*shared2.ParseTree[LexingTokenType], error, int) {
	// Adds children to the node until a newlinetoken is encountered
	line := &shared2.ParseTree[LexingTokenType]{
		Symbol:   "horizontal_line",
		Token:    nil,
		Children: make([]*shared2.ParseTree[LexingTokenType], 0),
	}

	possibleChildren := []LexingTokenType{XCharToken, MCharToken, ACharToken, SCharToken}

	for i := currentIndex; i < len(tokens); i++ {
		if tokens[i].Type == IgnoreToken {
			continue
		}
		if !sliceContains(possibleChildren, tokens[i].Type) {
			break
		}

		line.Children = append(line.Children, &shared2.ParseTree[LexingTokenType]{
			Symbol:   string(tokens[i].Value),
			Token:    tokens[i],
			Children: nil,
		})
	}

	if len(line.Children) == 0 {
		return nil, fmt.Errorf("expected at least 1 character to form a horizontal line"), 0
	}

	return line, nil, len(line.Children)
}

type InvalidTokenParserRule struct{}

func (I *InvalidTokenParserRule) Symbol() string {
	return "InvalidToken"
}

func (I *InvalidTokenParserRule) Match(tokens []*shared.Token[LexingTokenType], currentIndex int) (*shared2.ParseTree[LexingTokenType], error, int) {
	if tokens[currentIndex].Type == IgnoreToken {
		return nil, fmt.Errorf("expected non-ignore token"), 0
	}

	invalidToken := &shared2.ParseTree[LexingTokenType]{
		Symbol:   "invalid_token",
		Token:    tokens[currentIndex],
		Children: nil,
	}

	return invalidToken, nil, 1
}

type IgnoreTokenParserRule struct{}

func (I *IgnoreTokenParserRule) Symbol() string {
	return "IgnoreToken"
}

func (I *IgnoreTokenParserRule) Match(tokens []*shared.Token[LexingTokenType], currentIndex int) (*shared2.ParseTree[LexingTokenType], error, int) {
	if tokens[currentIndex].Type != IgnoreToken {
		return nil, fmt.Errorf("expected ignore token"), 0
	}

	return &shared2.ParseTree[LexingTokenType]{
		Symbol:   "ignore",
		Token:    tokens[currentIndex],
		Children: nil,
	}, nil, 1
}
