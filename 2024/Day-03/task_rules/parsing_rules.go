package task_rules

import (
	"fmt"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
	shared2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
)

type MultiplyOperationRuleParser struct{}

func (m *MultiplyOperationRuleParser) Symbol() string {
	return "MultiplyOperation"
}

func numWithinRange(target, min, max int) bool { // WOrks
	return target >= min && target <= max
}

func (m *MultiplyOperationRuleParser) Match(tokens []*shared.Token[LexingTokenType], currentIndex int) (*shared2.ParseTree[LexingTokenType], error, int) {
	if tokens[currentIndex].Type != MulKeywordToken {
		return nil, fmt.Errorf("not starting with keyword token"), 0
	}

	if tokens[currentIndex+1].Type != OpenParenthesisToken {
		return nil, fmt.Errorf("no open parenthesis token"), 0
	}

	if tokens[currentIndex+2].Type != NumberToken {
		return nil, fmt.Errorf("no number token on the left side"), 0
	}

	if !numWithinRange(len(tokens[currentIndex+2].Value), 1, 3) {
		return nil, fmt.Errorf("invalid number on the left side"), 0
	}

	if tokens[currentIndex+3].Type != CommaToken {
		return nil, fmt.Errorf("no comma separator"), 0
	}

	if tokens[currentIndex+4].Type != NumberToken {
		return nil, fmt.Errorf("no number token on the right side"), 0
	}

	if !numWithinRange(len(tokens[currentIndex+4].Value), 1, 3) {
		return nil, fmt.Errorf("invalid number on the right side"), 0
	}

	if tokens[currentIndex+5].Type != CloseParenthesisToken {
		return nil, fmt.Errorf("no closing parenthesis token"), 0
	}

	leftChild := &shared2.ParseTree[LexingTokenType]{
		Symbol:   "left_number",
		Token:    tokens[currentIndex+2],
		Children: nil,
	}

	rightChild := &shared2.ParseTree[LexingTokenType]{
		Symbol:   "right_number",
		Token:    tokens[currentIndex+4],
		Children: nil,
	}

	return &shared2.ParseTree[LexingTokenType]{
		Symbol:   "multiply_operation",
		Token:    nil,
		Children: []*shared2.ParseTree[LexingTokenType]{leftChild, rightChild},
	}, nil, 6
}

type InvalidTokenRuleParser struct{}

func (i *InvalidTokenRuleParser) Symbol() string {
	return "InvalidToken"
}

func (i *InvalidTokenRuleParser) Match(tokens []*shared.Token[LexingTokenType], currentIndex int) (*shared2.ParseTree[LexingTokenType], error, int) {
	invalidToken := &shared2.ParseTree[LexingTokenType]{
		Symbol:   "invalid_token",
		Token:    tokens[currentIndex],
		Children: nil,
	}

	return invalidToken, nil, 1
}

type DoRuleParser struct{}

func (d *DoRuleParser) Symbol() string {
	return "Do"
}

func (d *DoRuleParser) Match(tokens []*shared.Token[LexingTokenType], currentIndex int) (*shared2.ParseTree[LexingTokenType], error, int) {
	if tokens[currentIndex].Type != DoKeywordToken {
		return nil, fmt.Errorf("not starting with keyword token"), 0
	}

	if tokens[currentIndex+1].Type != OpenParenthesisToken {
		return nil, fmt.Errorf("no open parenthesis token"), 0
	}

	if tokens[currentIndex+2].Type != CloseParenthesisToken {
		return nil, fmt.Errorf("no closing parenthesis token"), 0
	}

	return &shared2.ParseTree[LexingTokenType]{
		Symbol:   "do",
		Token:    tokens[currentIndex],
		Children: nil,
	}, nil, 1
}

type DontRuleParser struct{}

func (d *DontRuleParser) Symbol() string {
	return "Dont"
}

func (d *DontRuleParser) Match(tokens []*shared.Token[LexingTokenType], currentIndex int) (*shared2.ParseTree[LexingTokenType], error, int) {
	if tokens[currentIndex].Type != DontKeywordToken {
		return nil, fmt.Errorf("not starting with keyword token"), 0
	}

	if tokens[currentIndex+1].Type != OpenParenthesisToken {
		return nil, fmt.Errorf("no open parenthesis token"), 0
	}

	if tokens[currentIndex+2].Type != CloseParenthesisToken {
		return nil, fmt.Errorf("no closing parenthesis token"), 0
	}

	return &shared2.ParseTree[LexingTokenType]{
		Symbol:   "don't",
		Token:    tokens[currentIndex],
		Children: nil,
	}, nil, 1
}
