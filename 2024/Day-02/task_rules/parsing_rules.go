package task_rules

import (
	"fmt"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
	shared2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
)

type SpaceRuleParser struct{}

func (w *SpaceRuleParser) Symbol() string {
	return "space"
}

func (w *SpaceRuleParser) Match(tokens []*shared.Token, currentIndex int) (*shared2.ParseTree, error, int) {
	currentToken := tokens[currentIndex]

	if currentToken.Type != shared.SpaceToken {
		return nil, fmt.Errorf("expected whitespace token, got %v", currentToken.Type), 0
	}

	return nil, nil, 1 // Spaces are not used.
}

type NewLineRuleParser struct{}

func (w *NewLineRuleParser) Symbol() string {
	return "newline"
}

func (w *NewLineRuleParser) Match(tokens []*shared.Token, currentIndex int) (*shared2.ParseTree, error, int) {
	currentToken := tokens[currentIndex]

	if currentToken.Type != shared.NewLineToken {
		return nil, fmt.Errorf("expected whitespace token, got %v", currentToken.Type), 0
	}

	return nil, nil, 1 // Newlines aren't used
}

type ReportRule struct{}

func (p *ReportRule) Symbol() string {
	return "report"
}

func (p *ReportRule) Match(tokens []*shared.Token, currentIndex int) (*shared2.ParseTree, error, int) {
	consumed := 0

	if currentIndex >= len(tokens) {
		return nil, fmt.Errorf("expected at least 1 token to form a report"), consumed
	}

	children := make([]*shared2.ParseTree, 0)

	// Loop through each token and add them to the children until a newline is reached
	for i := currentIndex; i < len(tokens); i++ {
		if tokens[i].Type == shared.NewLineToken {
			consumed++
			break
		}
		if tokens[i].Type == shared.SpaceToken {
			consumed++
			continue
		}
		children = append(children, &shared2.ParseTree{
			Symbol:   "report_item",
			Token:    tokens[i],
			Children: make([]*shared2.ParseTree, 0),
		})
	}

	if len(children) == 0 {
		return nil, fmt.Errorf("expected at least 1 token to form a report"), consumed
	}

	consumed += len(children)

	// Create a new parse tree for the report rule with the children
	return &shared2.ParseTree{
		Symbol:   p.Symbol(),
		Token:    nil,
		Children: children,
	}, nil, consumed
}
