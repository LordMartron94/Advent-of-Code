package shared

import (
	"fmt"
	"strings"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
)

type ParseTree[T comparable] struct {
	Symbol   string
	Token    *shared.Token[T]
	Children []*ParseTree[T]
}

// Print prints the parse tree with indentation
func (pt *ParseTree[T]) Print(indent int, ignoreTokens []T) {
	if pt == nil {
		return
	}

	for _, ignoreToken := range ignoreTokens {
		if pt.Token == nil {
			continue
		}

		if pt.Token.Type == ignoreToken {
			return
		}
	}

	fmt.Println(strings.Repeat("  ", indent) + pt.Symbol)

	if pt.Token != nil {
		fmt.Println(strings.Repeat("  ", indent+1) + "Token: " + fmt.Sprintf("%s (%v)", pt.Token.Value, pt.Token.Type))
	}

	for _, child := range pt.Children {
		child.Print(indent+1, ignoreTokens)
	}
}

func (pt *ParseTree[T]) GetNumberOfTokens() int {
	var count int

	if pt.Token != nil {
		count++
	}

	for _, child := range pt.Children {
		if child == nil {
			continue
		}

		count += child.GetNumberOfTokens()
	}

	return count
}
