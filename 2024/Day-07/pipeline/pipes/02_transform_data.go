package pipes

import (
	"fmt"
	"strconv"

	"github.com/LordMartron94/Advent-of-Code/2024/Day-07/pipeline/common"
	"github.com/LordMartron94/Advent-of-Code/2024/Day-07/task_rules"
	shared3 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/parsing/shared"
	"github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming"
	shared2 "github.com/LordMartron94/Advent-of-Code/_internal/utilities/transforming/shared"
)

type TransformDataPipe struct{}

func (t *TransformDataPipe) Process(input common.PipelineContext[task_rules.LexingTokenType]) common.PipelineContext[task_rules.LexingTokenType] {
	equations := make([]common.Equation, 0)

	callbackFinder := func(node *shared3.ParseTree[task_rules.LexingTokenType]) (shared2.TransformCallback[task_rules.LexingTokenType], int) {
		switch node.Symbol {
		case "equation":
			return func(node *shared3.ParseTree[task_rules.LexingTokenType]) {
				equation := common.Equation{}
				testNumber, err := strconv.Atoi(string(node.Children[0].Children[0].Token.Value))
				if err != nil {
					fmt.Println("Error converting this node")
					node.Print(2, []task_rules.LexingTokenType{})
					panic(err)
				}

				equation.TestNumber = testNumber
				testParts := node.Children[1].Children

				for _, part := range testParts {
					if part.Token.Type != task_rules.NumberToken {
						continue
					}

					element, err := strconv.Atoi(string(part.Token.Value))
					if err != nil {
						panic(err)
					}
					equation.Elements = append(equation.Elements, element)
				}

				equations = append(equations, equation)
			}, 0
		}
		return nil, 0
	}

	transformer := transforming.NewTransformer(
		callbackFinder,
	)
	transformer.Transform(input.ParseTree)

	input.Equations = equations

	return input
}
