package shared

import (
	"fmt"
	"strings"

	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/lexing/shared"
)

type ParseTree struct {
	Symbol   string
	Token    *shared.Token
	Children []*ParseTree
}

// Print prints the parse tree with indentation
func (pt *ParseTree) Print(indent int) {
	fmt.Println(strings.Repeat("  ", indent) + pt.Symbol)
	if pt.Token != nil {
		fmt.Println(strings.Repeat("  ", indent+1) + "Token: " + fmt.Sprintf("%s (%v)", pt.Token.Value, pt.Token.Type))
	}
	for _, child := range pt.Children {
		child.Print(indent + 1)
	}
}
